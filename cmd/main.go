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

	// Create a GET route which receives path parameters to search an article by date and slug
	r.Get("/articles/{date}/{slug}", getArticle)

	// Start the HTTP server on port 8080
	http.ListenAndServe(":8080", r)
}

// getArticle handles the request to fetch an article by date and slug
func getArticle(w http.ResponseWriter, r *http.Request) {
	dateParam := chi.URLParam(r, "date")
	slugParam := chi.URLParam(r, "slug")
	// article, err := database.GetArticle(dateParam, slugParam)
	// For demonstration, we will just write the parameters back to the response
	w.Write([]byte("Article Date: " + dateParam + ", Slug: " + slugParam))
}
