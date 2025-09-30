package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})

	http.HandleFunc("/sleep", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Second)
		fmt.Fprintf(w, "Hello World")
	})

	http.HandleFunc("/404", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Not an actual not_found")
	})

	fmt.Println("Server starting on :8080")
	http.ListenAndServe(":8080", nil)
}
