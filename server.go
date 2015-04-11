package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("c:/__HTML"))
	http.Handle("/", fs)

	log.Println("3000 Listening...")
	http.ListenAndServe(":3000", nil)
}