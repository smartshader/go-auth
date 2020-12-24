package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type person struct {
	First string
}

func main() {
	http.HandleFunc("/encode", encode)
	http.HandleFunc("/decode", decode)

	http.ListenAndServe(":8080", nil)
}

func encode(w http.ResponseWriter, r *http.Request) {
	p := person{First: "Jenny"}

	encoder := json.NewEncoder(w)
	err := encoder.Encode(p)
	if err != nil {
		http.Error(w, "error encoding json", http.StatusInternalServerError)
	}
}

func decode(w http.ResponseWriter, r *http.Request) {
	var p person

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, "error decoding json", http.StatusInternalServerError)
	}

	log.Println("Person: ", p)
}
