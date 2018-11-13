package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"net/http"
	"time"
	"io/ioutil"
	"encoding/json"
	"os/user"
	"log"
)

var cnfDir = ".xnaming"

var command = make(chan string)
var workerRsp = make(chan int)
//var rareFirst = []string{}
//var rareSecond = []string{}



func main() {

	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	cnfDir = usr.HomeDir + "/" + cnfDir

	for i := 0; i < 2; i++ {
		go startWorker()
	}
	reader := bufio.NewReader(os.Stdin)
	for {
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		if "quit" == text {
			break
		}
		command <- text
		<-workerRsp
	}
}

func startWorker() {
	for {
		switch command := <-command; command {
		case "reload":
			log.Println("Reload command received.")
			loadRareWordsDict()
		case "home":
			initTempDir()
		default:
			log.Println("unknown command", command)
		}
		workerRsp <- 1
	}
}



func initTempDir() {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	xnamingPath := usr.HomeDir + "/" + cnfDir

	// create directory if needed
	if _, err := os.Stat(xnamingPath); os.IsNotExist(err) {
		os.Mkdir(xnamingPath, os.ModePerm)
	}
	// create configuration file
	xnamingConf := xnamingPath + "/xnaming.cnf"
	if _, err := os.Stat(xnamingConf); os.IsNotExist(err) {
		os.Create(xnamingConf)
	}

	// open configuration file
	cnf, err := os.Create(xnamingConf)
	log.Println(cnf)

	defer func() {
		if err := cnf.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	writer := bufio.NewWriter(cnf)
	if _, err := writer.WriteString("Github repo:"); err != nil {
		log.Fatalln(err)
	}
	if err := writer.Flush(); err != nil {
		log.Fatalln(err)
	}
}

// load rare words dict from github
func loadRareWordsDict() {

	// http get branch/master tree meta data
	dat := getHttpData(
		"https://api.github.com/repos/0x0010/notebook/branches/master",
		5*time.Second,
		"branch/master commit meta")
	//get tree url
	repoTreeUrl := dat["commit"].(map[string]interface{})["commit"].(map[string]interface{})["tree"].(map[string]interface{})["url"]

	//http get tree data
	dat = getHttpData(repoTreeUrl.(string), 5*time.Second, "index tree")
	//get tree list
	treeSlice := dat["tree"].([]interface{})

	resourcesUrl := ""
	for _, treeNode := range treeSlice {
		if "resources" == treeNode.(map[string]interface{})["path"] {
			resourcesUrl = treeNode.(map[string]interface{})["url"].(string)
		}
	}

	dat = getHttpData(resourcesUrl, 5*time.Second, "resources tree")
	treeSlice = dat["tree"].([]interface{})
	for _, treeNode := range treeSlice {
		blobPath := treeNode.(map[string]interface{})["path"]
		if "rarely-words-first-part.txt" == blobPath ||
			"rarely-words-second-part.txt" == blobPath {
				ioutil.WriteFile(cnfDir + "/" + blobPath.(string), strings.TrimSpace(getBlobData(treeNode.(map[string]interface{})["url"].(string))), os.ModePerm)
			//fmt.Println(strings.TrimSpace(getBlobData(treeNode.(map[string]interface{})["url"].(string))))
		}
	}

}

func getBlobData(url string) string {
	dat := getHttpData(url, 5*time.Second, "Get blob data")
	return dat["content"].(string)
}


func getHttpData(url string, timeout time.Duration, desc string) map[string]interface{} {
	httpClient := http.Client{Timeout: timeout}
	rsp, err := httpClient.Get(url)
	if nil != err {
		panic(fmt.Sprintf("http get failed, url[%s], desc[%s], err[%v]", url, desc, err))
	}
	defer rsp.Body.Close()
	body, err := ioutil.ReadAll(rsp.Body)
	if nil != err {
		panic(fmt.Sprintf("reading response content failed, url[%s], desc[%s], err[%v]", url, desc, err))
	}
	var dat map[string]interface{}
	if err := json.Unmarshal(body, &dat); nil != err {
		panic(err)
	}
	return dat
}
