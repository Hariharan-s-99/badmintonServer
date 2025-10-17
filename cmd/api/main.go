package main

import (
	"fmt"
	"log"
	"net/http"

	"myapp/internal/handler"
	"myapp/internal/middleware"
	"myapp/internal/router"
)

func main() {
	// Initialize handlers
	homeHandler := handler.NewHomeHandler()

	// Initialize router
	mux := router.NewRouter(homeHandler)

	// Wrap with middleware
	loggedMux := middleware.LoggingMiddleware(mux)

	// Start server
	port := ":8080"
	fmt.Printf("Server starting on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, loggedMux))
}