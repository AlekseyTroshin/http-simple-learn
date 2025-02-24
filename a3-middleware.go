package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", homeHandlerA3)
	http.HandleFunc("/about", aboutHandlerA3)
	fmt.Println("Server run on http://localhost:1234")
	http.ListenAndServe(":1234", nil)
}

func aboutHandlerA3(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintln(w, "Home Page")
	if err != nil {
		return
	}
}

func homeHandlerA3(w http.ResponseWriter, t *http.Request) {

}
