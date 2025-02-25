package main

import (
	"fmt"
	"net/http"
)

//проверит запрос curl -X POST http://localhost:1234 -d "name=Testovushkin"

func main() {
	http.HandleFunc("/", fHandlerA2)
	http.HandleFunc("/form", formHandlerA2)
	fmt.Println("Server run on http://localhost:1234")
	http.ListenAndServe(":1234", nil)
}

func fHandlerA2(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			return
		}

		name := r.FormValue("name")
		_, err = fmt.Fprintf(w, "Hello, %s!\n", name)
		if err != nil {
			return
		}
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func formHandlerA2(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintln(w, "Hello Hi ! )")
	if err != nil {
		return
	}
}
