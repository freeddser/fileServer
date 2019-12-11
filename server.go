package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

const static_path = "upload"

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	t := time.Unix(time.Now().Unix(), 0).Format("2006-01-02-15-04-05")
	name := r.URL.Query().Get("name")
	if name != "" {
		fileName := t + "_" + name
		file, err := os.Create(static_path + "/" + fileName)
		if err != nil {
			panic(err)
		}
		_, err = io.Copy(file, r.Body)
		if err != nil {
			panic(err)
		}
		log.Println(fileName + " upload success")
		w.Write([]byte(fileName + " upload success"))
	} else {
		w.Write([]byte("Please input Filename"))
	}

}

func main() {
	http.HandleFunc("/upload", uploadHandler)
	http.Handle("/", http.FileServer(http.Dir("upload")))
	var port int
	flag.IntVar(&port, "port", 5050, "server port, default 5050")
	flag.Parse()
	log.Print("Server running on port:", port)

	http.ListenAndServe(":"+strconv.Itoa(port), nil)

}
