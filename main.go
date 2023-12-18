package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(http.ResponseWriter, *http.Request) {
		log.Println("Hellp World")
	})

	http.HandleFunc("/goodBye", func(http.ResponseWriter, *http.Request) {
		log.Println("Bye bye World")
	})

	http.ListenAndServe(":9090", nil)
}
