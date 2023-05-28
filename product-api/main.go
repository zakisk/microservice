package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Oops error", http.StatusBadRequest)
			return
		}

		fmt.Fprintf(w, "Hello %s\n", data)
	})

	http.HandleFunc("/goodbye", func(w http.ResponseWriter, r *http.Request) {
		data, _ := ioutil.ReadAll(r.Body)
		fmt.Fprintf(w, "Goodbye %s\n", data)
	})

	log.Fatal(http.ListenAndServe(":9090", nil))
}