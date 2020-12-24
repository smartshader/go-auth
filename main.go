package main

import (
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

}

func decode(w http.ResponseWriter, r *http.Request) {

}
