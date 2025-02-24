package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", hiHandler)
	fmt.Println("Server tan on http://localhost:1234")
	err := http.ListenAndServe(":1234", nil)
	if err != nil {
		return
	}
}

func hiHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintln(w, "Hi, World !!!")
	if err != nil {
		return
	}
}
