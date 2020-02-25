package main

import (
	"log"
	"net/http"
	"os"
)

// HoneyKey is the key protection for the http calls.
var HoneyKey = os.Getenv("HONEY_KEY")

func main() {
	if HoneyKey == "" {
		log.Fatalln("make sure to have a HONEY_KEY env var")
	}
	router := NewRouter()
	http.ListenAndServe(":8080", router)
}
