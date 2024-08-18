package config

import (
	"os"

	"github.com/joho/godotenv"
)

type App struct {
	Author struct {
		Name string
	}
	Server struct {
		Port   string
		Host   string
		Socket string
	}
}

func Init() App {
	godotenv.Load()

	var app App

	// AUTHOR
	app.Author.Name = os.Getenv("AUTHOR_NAME")
	if app.Author.Name == "" {
		app.Author.Name = "anonymous"
	}

	// SERVER
	app.Server.Host = "localhost"
	if os.Getenv("HOST") != "" {
		app.Server.Host = os.Getenv("HOST")
	}

	app.Server.Port = "3000"
	if os.Getenv("PORT") != "" {
		app.Server.Port = os.Getenv("PORT")
	}

	// SOCKET = localhost:3000
	app.Server.Socket = app.Server.Host + ":" + app.Server.Port

	return app
}
