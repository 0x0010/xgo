package main

import (
	"net"
	"log"
	"fmt"
	"os"
	"time"
	"strings"
	"strconv"
)

func main() {
	binding := ":8091"

	tcpAddr, err := net.ResolveTCPAddr("tcp4", binding)
	checkErr(err, "Unknown socket binding address")

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkErr(err, fmt.Sprintf("Failed to listen on %s ", tcpAddr))

	for {
		conn, err := listener.Accept()
		if nil != err {
			fmt.Fprintln(os.Stderr, "Accept exception occurred")
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	connInfo := conn.RemoteAddr().String()
	log.Println("Client connected:", connInfo)
	conn.SetReadDeadline(time.Now().Add(30 * time.Second))
	request := make([]byte, 128)
	defer conn.Close()

	for {
		readLen, err := conn.Read(request)
		if nil != err {
			fmt.Println(err)
			break
		}

		if readLen == 0 {
			break
		}

		// refresh read timeout deadline
		conn.SetReadDeadline(time.Now().Add(30 * time.Second))

		requestStr := strings.TrimSpace(string(request[:readLen]))
		log.Println(connInfo, "->", requestStr)

		if requestStr == "timestamp" {
			daytime := strconv.FormatInt(time.Now().Unix(), 10) + "\n"
			conn.Write([]byte(daytime))
		} else if requestStr == "quit" {
			// close connection by defer
			//conn.Close()
			break
		} else {
			daytime := time.Now().String() + "\n"
			conn.Write([]byte(daytime))
		}
		request = make([]byte, 128)

	}

}

func checkErr(err error, msg string) {
	if nil != err {
		log.Fatalln(msg)
	}
}
