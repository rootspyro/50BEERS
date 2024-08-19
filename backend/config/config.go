package config

import (
	"fmt"
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
	Database struct {
		Name     string
		Host     string
		Port     string
		Username string
		Password string
		URL      string
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


	// MONGODB
	app.Database.Host = os.Getenv("DB_HOST")
	app.Database.Port = os.Getenv("DB_PORT")
	app.Database.Username = os.Getenv("DB_USERNAME")
	app.Database.Password = os.Getenv("DB_PASSWORD")
	app.Database.Name = os.Getenv("DB_NAME")
	app.Database.URL = fmt.Sprintf(
		"mongodb://%s:%s@%s:%s/%s",
		app.Database.Username,
		app.Database.Password,
		app.Database.Host,
		app.Database.Port,
		app.Database.Name,
	)

	return app
}
