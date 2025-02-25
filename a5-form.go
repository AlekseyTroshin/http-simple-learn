package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/form", formHandlerA5)
	fmt.Println("Server started at http://localhost:1234")
	http.ListenAndServe(":1234", nil)
}

func formHandlerA5(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		name := r.FormValue("name")
		message := r.FormValue("message")
		fmt.Fprintf(w, "Received massage from %s: %s", name, message)
	} else {
		http.ServeFile(w, r, "/formA5.html")
	}
}
