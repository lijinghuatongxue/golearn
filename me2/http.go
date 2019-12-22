package main

import (
	"fmt"
	"net/http"
)

func SayHello(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello"))
}

func main() {
	http.HandleFunc("/hello", SayHello)
	fmt.Println("Starting http，port is 8001")
	http.ListenAndServe("0.0.0.0:8001", nil)
}

