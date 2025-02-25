package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Message struct {
	Name string `json:"name"`
	Text string `json:"text"`
}

func jsonHandlerA4(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var msg Message
		err := json.NewDecoder(r.Body).Decode(&msg)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		response := fmt.Sprintf("Received message from %s: %s", msg.Name, msg.Text)
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(map[string]string{"response": response})
		if err != nil {
			return
		}
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/json", jsonHandlerA4)
	fmt.Println("Server started at http://localhost:1234")
	http.ListenAndServe(":1234", nil)
}
