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

	aiName := "Victorika" // default value for AI_NAME
	if envAiName, ok := os.LookupEnv("AI_NAME"); ok {
		aiName = envAiName // check env AI_NAME
	}

	aiTarget := "" // default value for AI_TARGET
	if envAiTarget, ok := os.LookupEnv("AI_TARGET"); ok {
		aiTarget = envAiTarget // check env AI_TARGET
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet || r.Method == http.MethodHead {
			response := fmt.Sprintf("Hello World -2, %s is your assistant %s", aiName, aiTarget)
			fmt.Fprintf(w, response) // example answer
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