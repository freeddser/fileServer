package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	var fileName string
	var server string
	flag.StringVar(&server, "server", "http://127.0.0.1:5050", "server address, default:http://127.0.0.1:5050")
	flag.StringVar(&fileName, "filename", "", "filename and path: /tmp/a.txt")
	flag.Parse()
	if fileName==""{
		fmt.Println("\n\nUse -h to get more information on command line options\n")
		fmt.Println("You must specify a filename")
		os.Exit(1)
	}
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	res, err := http.Post(server+"/upload?name="+fileName, "binary/octet-stream", file)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	message, _ := ioutil.ReadAll(res.Body)
	log.Printf(string(message))
}
