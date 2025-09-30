package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/sleep", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Second)
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, "<!DOCTYPE html><html><head><title>Sleep Complete</title></head><body><h1>Hello World</h1><p>Slept for 1 second</p></body></html>")
	})

	// Root handler with 404 for unmatched routes
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "<!DOCTYPE html><html><head><title>404 Not Found</title></head><body><h1>404 - Page Not Found</h1><p>The requested page could not be found.</p></body></html>")
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, "<!DOCTYPE html><html><head><title>Hello World</title></head><body><h1>Hello World</h1></body></html>")
	})

	fmt.Println("Server starting on :8080")
	http.ListenAndServe(":8080", nil)
}
