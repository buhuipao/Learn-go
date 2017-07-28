package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", HelloHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal("ListenServer: ", err.Error())
	}
}

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Inside HelloHandler...")
	// 把字符串写到w里
	fmt.Fprintf(w, "Hello, "+req.URL.Path[1:])
}
