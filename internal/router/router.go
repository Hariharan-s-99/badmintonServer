package router

import (
	"net/http"

	"myapp/internal/handler"
)

func NewRouter(homeHandler *handler.HomeHandler) *http.ServeMux {
	mux := http.NewServeMux()

	// Define routes
	mux.HandleFunc("/", homeHandler.Home)

	return mux
}