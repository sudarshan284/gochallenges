//in this challenge we will build a simple basic hello world server using golang

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", NewServer)
	fmt.Printf("Server starting at port ...8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func NewServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world program")
}
