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

const static_path = "server/upload"

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	t := time.Now().Unix()
	name := r.URL.Query().Get("name")
	if name != "" {
		fileName := strconv.FormatInt(t, 10) + "_" + name
		file, err := os.Create(static_path + "/" + fileName)
		if err != nil {
			panic(err)
		}
		_, err = io.Copy(file, r.Body)
		if err != nil {
			panic(err)
		}
		w.Write([]byte(fileName + " upload success"))
	} else {
		w.Write([]byte("Please input Filename"))
	}

}

func main() {
	http.HandleFunc("/upload", uploadHandler)
	var port int
	flag.IntVar(&port, "port", 5050, "server port, default 5050")
	flag.Parse()
	log.Print("Server running on port:", port)
	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}
