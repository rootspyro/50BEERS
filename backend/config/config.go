package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

var App app = app{}

type app struct {
	Author struct {
		Name string
		Email string
	}
	Server struct {
		Secret  string
		Port    string
		Host    string
		Socket  string
		Origins []string
	}
	Database struct {
		Name     string
		Host     string
		Port     string
		Username string
		Password string
		URL      string
	}
	SDKs struct {
		Mailtrap struct {
			Host     string
			APIToken string
		}
	}
}

func init() {
	godotenv.Load()

	// AUTHOR
	App.Author.Name = os.Getenv("AUTHOR_NAME")
	if App.Author.Name == "" {
		App.Author.Name = "anonymous"
	}

	App.Author.Email = os.Getenv("AUTHOR_EMAIL")

	// SERVER
	App.Server.Host = "localhost"
	if os.Getenv("SECRET") != "" {
		App.Server.Host = os.Getenv("SECRET")
	}

	App.Server.Host = "localhost"
	if os.Getenv("HOST") != "" {
		App.Server.Host = os.Getenv("HOST")
	}

	App.Server.Port = "3000"
	if os.Getenv("PORT") != "" {
		App.Server.Port = os.Getenv("PORT")
	}

	App.Server.Origins = parseAllowedOrigins(os.Getenv("ORIGINS"))

	// SOCKET = localhost:3000
	App.Server.Socket = App.Server.Host + ":" + App.Server.Port

	// MONGODB
	App.Database.Host = os.Getenv("DB_HOST")
	App.Database.Port = os.Getenv("DB_PORT")
	App.Database.Username = os.Getenv("DB_USER")
	App.Database.Password = os.Getenv("DB_PASSWORD")
	App.Database.Name = os.Getenv("DB_NAME")
	App.Database.URL = fmt.Sprintf(
		"mongodb://%s:%s@%s:%s/%s",
		App.Database.Username,
		App.Database.Password,
		App.Database.Host,
		App.Database.Port,
		App.Database.Name,
	)

	// SDKs
	App.SDKs.Mailtrap.Host = os.Getenv("MAILTRAP_HOST")
	App.SDKs.Mailtrap.APIToken = os.Getenv("MAILTRAP_API_TOKEN")
}

func parseAllowedOrigins(data string) []string {
	var origins []string

	parts := strings.Split(data, ",")

	for _, origin := range parts {
		origins = append(origins, origin)
	}

	return origins
}
