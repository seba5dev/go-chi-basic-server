package server

import (
	"go-chi-basic-server/internal/middleware"
	"net/http"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
)

// ConfigServerChi holds the configuration for the Chi server
type ConfigServerChi struct {
	// ServerAddress is the address where the server will be listening
	ServerAddress string
	// LoaderFilePath is the path to the file that contains the database (json or csv now) we won't use it for now
	LoaderFilePath string
}

// NewServerChi is a function that returns a new instance of ServerChi
func NewServerChi(cfg *ConfigServerChi) *ServerChi {
	// default values
	defaultConfig := &ConfigServerChi{
		ServerAddress: ":8080",
	}
	// here we will override the default values with the values from the config if they are set
	if cfg != nil {
		if cfg.ServerAddress != "" {
			defaultConfig.ServerAddress = cfg.ServerAddress
		}
		if cfg.LoaderFilePath != "" {
			defaultConfig.LoaderFilePath = cfg.LoaderFilePath
		}
	}

	return &ServerChi{
		serverAddress:  defaultConfig.ServerAddress,
		loaderFilePath: defaultConfig.LoaderFilePath,
	}
}

// ServerChi is a struct that implements the Application interface
type ServerChi struct {
	// serverAddress is the address where the server will be listening
	serverAddress string
	// loaderFilePath is the path to the file that contains the  database (json or csv now) we won't use it for now
	loaderFilePath string
}

// Run is a method that runs the server
func (a *ServerChi) Run() error {
	// dependencies
	// - loader
	// - repository
	// - service
	// - handler
	// router
	r := chi.NewRouter()
	// - middlewares
	r.Use(middleware.Logging)
	r.Use(chiMiddleware.Recoverer)
	// - endpoints
	r.Route("/songs", func(rt chi.Router) {
		// - GET /songs
		rt.Get("/", func(w http.ResponseWriter, r *http.Request) {
			// For demonstration, we will just write a response
			w.Write([]byte("List of songs"))
		})
	})

	// Create 404 handler for unmatched routes
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		// This will send a 404 response with a custom message
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - Not Found"))
	})

	// Print the server address to the console
	println("Server is running on", a.serverAddress)
	// run server
	err := http.ListenAndServe(a.serverAddress, r)
	if err != nil {
		return err
	}
	return nil
}
