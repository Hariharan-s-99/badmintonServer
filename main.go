package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	// Create a new ServeMux
	mux := http.NewServeMux()

	// Define routes
	mux.HandleFunc("/", homeHandler)

	// Wrap mux with logging middleware
	loggedMux := loggingMiddleware(mux)

	// Start server
	port := ":8080"
	fmt.Printf("Server starting on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, loggedMux))
}

// Logging middleware
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Log incoming request
		log.Printf("[%s] %s %s", r.Method, r.URL.Path, r.RemoteAddr)

		// Call the next handler
		next.ServeHTTP(w, r)

		// Log request duration
		duration := time.Since(start)
		log.Printf("[%s] %s - Completed in %v", r.Method, r.URL.Path, duration)
	})
}