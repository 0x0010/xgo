package main

import (
	"log"
	"bytes"
	"encoding/binary"
	"net"
	"time"
	"strings"
	"os"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-sql-driver/mysql"
)

const (
	MsgTypeC2S = uint16(689)
)

type DyProtocol struct {
	length   uint32
	msgType  uint16
	encrypt  uint8
	reserved uint8
	data     string
}

type MessageBody struct {
	msgType string
	uid     string
	level   string
	nn      string
	txt     string
	bnn     string
	bl      string
}

func newDyProtocol(data string, msgType uint16) *DyProtocol {
	return &DyProtocol{0, msgType, uint8(0), uint8(0), data}
}

func (p *DyProtocol) serialize() []byte {
	dataBytes := []byte(p.data)
	msgLenBytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(msgLenBytes, uint32(len(dataBytes)+9))

	buffer := bytes.NewBuffer([]byte{})
	buffer.Write(msgLenBytes)
	buffer.Write(msgLenBytes)

	binary.Write(buffer, binary.LittleEndian, p.msgType)
	binary.Write(buffer, binary.LittleEndian, p.encrypt)
	binary.Write(buffer, binary.LittleEndian, p.reserved)
	// writes message body
	buffer.Write(dataBytes)
	// message body end
	binary.Write(buffer, binary.LittleEndian, uint8(0))
	return buffer.Bytes()
}

func main() {
	if len(os.Args) <= 1 {
		log.Panic("Invalid arguments, roomId is required!")
	}
	roomId := os.Args[1:][0]

	config := mysql.NewConfig()
	config.User = "root"
	config.Passwd = "root123"
	config.Addr = "127.0.0.1:3306"
	config.DBName = "douyubarrage"

	//open mysql
	//db, err := sql.Open("mysql", "root:root123@127.0.0.1:3306/douyubarrage")
	db, err := sql.Open("mysql", config.FormatDSN())

	if err != nil {
		log.Panic(err)
	}
	defer db.Close()
	db.SetConnMaxLifetime(10 * time.Minute)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)


	conn := dialServer()
	defer logout(conn)

	loginRsp := login(conn, roomId)
	if len(loginRsp) <= 0 {
		log.Panic("Login Barrage Server Failed!!!!")
	}

	log.Print("Login Success!")
	joinGroup(conn, roomId)
	log.Print("Join Group Success!")

	go heartbeat(conn)

	for {
		msg, err := readMessage(conn, 5*time.Second)
		if nil != err {
			log.Fatal(err)
		}
		if len(msg) > 0 {
			message := decodeMessage(msg)
			if strings.Compare("chatmsg", message.msgType) == 0 {
				log.Printf("UserId: %10s, UserName:%s, UserLvl: %s, Bnn: %s, BnLvl:%s, Txt:%s",
					message.uid, message.nn, message.level, message.bnn, message.bl, message.txt)
				saveMessageToDB(db, message, roomId)
			}
		}
	}
}

func dialServer() net.Conn {
	conn, err := net.Dial("tcp", "openbarrage.douyutv.com:8601")
	if nil != err {
		log.Panic(err)
	}
	return conn
}

func login(conn net.Conn, roomId string) string {
	conn.Write(newDyProtocol("type@=loginreq/roomid@="+roomId+"/", MsgTypeC2S).serialize())
	msg, err := readMessage(conn, 5*time.Second)
	if nil != err {
		log.Panic(err)
	}
	return msg
}

func logout(conn net.Conn) {
	conn.Write(newDyProtocol("type@=logout/", MsgTypeC2S).serialize())
	conn.Close()
}

func joinGroup(conn net.Conn, roomId string) {
	conn.Write(newDyProtocol("type@=joingroup/gid@=-9999/rid@="+roomId+"/", MsgTypeC2S).serialize())
}

func readMessage(conn net.Conn, d time.Duration) (msg string, err error) {

	// read first 4 bytes
	first4bytes := make([]byte, 4)
	conn.SetReadDeadline(time.Now().Add(d))
	conn.Read(first4bytes)
	msgBytesCount := binary.LittleEndian.Uint32(first4bytes)
	if msgBytesCount == 0 {
		time.Sleep(1 * time.Second)
		return "", nil
	}
	//log.Print("message length ", msgBytesCount)

	msgBody := make([]byte, msgBytesCount)
	conn.SetReadDeadline(time.Now().Add(d))
	count, err := conn.Read(msgBody)
	//log.Print("message read length ", count)
	if uint32(count) != msgBytesCount {
		//log.Fatal(err)
		return "", nil
	}
	if nil != err {
		return "", err
	}
	return string(msgBody[8: count-1]), nil
}

func heartbeat(conn net.Conn) {
	for {
		time.Sleep(45 * time.Second)
		conn.Write(newDyProtocol("type@=mrkl/", MsgTypeC2S).serialize())
		log.Print("Heartbeat Sent")
	}
}

func decodeMessage(message string) MessageBody {
	kvs := strings.Split(message, "/")
	mb := MessageBody{}
	for _, kv := range kvs {
		entry := strings.Split(kv, "@=")
		if len(entry) != 2 {
			continue
		}
		entryValue := entry[1]
		switch entry[0] {
		case "type":
			mb.msgType = entryValue
			break
		case "uid":
			mb.uid = entryValue
			break
		case "nn":
			mb.nn = entryValue
			break
		case "level":
			mb.level = entryValue
			break
		case "txt":
			mb.txt = entryValue
			break
		case "bnn":
			mb.bnn = entryValue
			break
		case "bl":
			mb.bl = entryValue
			break
		default:
			continue
		}
	}
	return mb
}

func saveMessageToDB(db *sql.DB, msg MessageBody, roomId string) {
	stmtIns, err := db.Prepare("INSERT INTO chatmessage(RoomId, UID, UName, ULevel, BNN, BNNLevel, TXT) VALUES( ?, ?, ?, ?, ?, ?, ?)") // ? = placeholder
	if err != nil {
		log.Print("Prepare error! ", err)
	}
	_, err = stmtIns.Exec(roomId, msg.uid, msg.nn, msg.level, msg.bnn, msg.bl, msg.txt)
	if err != nil {
		log.Print("Prepare execute error! ", err)
	}
}
