
Chapter1:
==============  webserver in Go =====================

package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Hellp World")
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Oops"))
			return
		}

		log.Printf("Data %s\n", d)
	})

	http.HandleFunc("/goodBye", func(http.ResponseWriter, *http.Request) {
		log.Println("Bye bye World")
	})

	http.ListenAndServe(":9090", nil)
}

=====================
