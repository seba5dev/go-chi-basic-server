package middleware

import (
	"fmt"
	"net/http"
	"time"
)

type Audit struct {
	Method        string    `json:"method"`
	Path          string    `json:"path"`
	ContentLength int64     `json:"content_lenght"`
	Ts            time.Time `json:"ts"`
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Create an audit log entry
		audit := Audit{
			Method:        r.Method,
			Path:          r.URL.Path,
			ContentLength: r.ContentLength,
			Ts:            time.Now(),
		}

		// Print the audit log entry to the console
		fmt.Println(audit)

		// Call the next handler in the chain
		next.ServeHTTP(w, r)
	})
}
