package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func getFiles() string {
	var file_name []string
	files, err := ioutil.ReadDir("./proto/")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.Name() != "._home.proto" {
			file_name = append(file_name, file.Name())
		}
		// fmt.Println(file.Name(), file.IsDir())
	}
	//fmt.Println("file name list : ", file_name)
	file := strings.Join(file_name, ",")
	return file
}

func getFileName(w http.ResponseWriter, r *http.Request) {
	file_name := getFiles()
	fmt.Println(file_name)
	//b := []byte(file_name)
	//w.Write(b)
	w.Header().Add("file_data", file_name)
	io.WriteString(w, file_name)
}

func main() {
	fs := http.FileServer(http.Dir("./"))
	http.Handle("/", fs)
	http.HandleFunc("/file/", getFileName)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
