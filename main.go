package main

import (
	"fmt"
	"net/http"
	"log"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Horus!")
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":1123", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
