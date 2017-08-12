package main

import (
	"os"
	"log"
	"net/http"
	"time"
	"io/ioutil"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		log.Fatal("Wrong arguments: ", args)
	}

	httpClient := http.Client{
		Timeout: time.Duration( 5 * time.Second),
	}
	response, err := httpClient.Get(args[0])
	if nil != err {
		log.Fatalf("Fetch url [%s] failed, error message: %s", args[0], err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	log.Println(string(body))
}