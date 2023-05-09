package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/time", timeHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}


func helloHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("serving request: %s", r.URL.Path)

	host, err := os.Hostname()
	
	if err != nil {
		http.Error(w, "Failed to get the hostname", http.StatusInternalServerError)
		log.Printf("Error getting hostname: %v", err)
		return
	}

	fmt.Fprintf(w, "Congratulations, You've successfully deployed a simple web server without TLS in a docker.\n")
	fmt.Fprintf(w, "Protocol: %s!\n", r.Proto)
	fmt.Fprintf(w, "Hostname: %s\n", host)

	if headerIP := r.Header.Get("X-Forwarded-For"); headerIP != "" {
		fmt.Fprintf(w, "Client IP (X-Forwarded-For): %s\n", headerIP)
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("serving request: %s", r.URL.Path)

	// Return a 200 status code and a message to indicate that the server is running
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Server is running OK. ❤️ ")
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("serving request: %s", r.URL.Path)

	// Return the current server time
	currentTime := time.Now().Format(time.RFC1123)
	fmt.Fprintf(w, "Current server time: %s", currentTime)
}