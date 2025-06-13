package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// Create a new router
	r := chi.NewRouter()
	// Use middleware for logging (this is made from chi)
	r.Use(middleware.Logger)
	// Create a simple GET route with a string response "Hello World!"
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		// byte is used to convert string to bytes
		w.Write([]byte("Hello World!"))
	})
	// Start the HTTP server on port 8080
	http.ListenAndServe(":8080", r)
}
