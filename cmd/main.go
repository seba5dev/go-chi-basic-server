package main

import (
	"fmt"
	"go-chi-basic-server/cmd/server"
)

func main() {
	// env
	// ...

	// app
	// config
	cfg := &server.ConfigServerChi{
		ServerAddress:  ":8080",
		LoaderFilePath: "docs/db/songs.json",
	}
	app := server.NewServerChi(cfg)
	// run
	if err := app.Run(); err != nil {
		fmt.Println("Error running the server:", err)
		return
	}
}
