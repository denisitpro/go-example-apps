package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := "8082" // default value port
	if envPort, ok := os.LookupEnv("PORT"); ok {
		port = envPort // check env PORT
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet || r.Method == http.MethodHead {
			fmt.Fprintf(w, "Hello World") // example answer
			logRequest(r)
		} else {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed) // error exception
		}
	})

	fmt.Printf("Server is listening on port %s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}

// log address request
func logRequest(r *http.Request) {
	fmt.Printf("Request from %s: %s %s\n", r.RemoteAddr, r.Method, r.URL.Path)
}