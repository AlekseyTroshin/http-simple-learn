package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", fHandlerA3)

	wrappedMux := loggingMiddlewareA3(mux)

	fmt.Println("Server run on http://localhost:1234")
	err := http.ListenAndServe(":1234", wrappedMux)
	if err != nil {
		return
	}
}

func fHandlerA3(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello ! ))")
}

func loggingMiddlewareA3(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
	})
}
