package bloguser

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/rootspyro/50BEERS/db"
	"github.com/rootspyro/50BEERS/db/repositories"
	"github.com/rootspyro/50BEERS/middlewares"
	"github.com/rootspyro/50BEERS/services"
)

func connString() (string, string) {
	godotenv.Load("../../.env")

	return fmt.Sprintf(
		"mongodb://%s:%s@%s:%s/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	), os.Getenv("DB_NAME")
}

func TestSignUpFromSite(t *testing.T) {
	connStr, dbName := connString()

	dbClient, err := db.New(connStr)
	if err != nil {
		t.Error(err)
	}

	database := dbClient.Database(dbName)

	repo := repositories.NewBlogUserRepo(database.Collection("blogUser"))
	srv := services.NewBlogUserSrv(repo)
	handler := NewBlogUserHandler(srv)

	// build the testing server
	server := httptest.NewServer(http.HandlerFunc(middlewares.PipeNewDrinkBody(handler.SignUp)))

	// do the request
	body := services.BlogUserDTO {
		Username: "user",
		Email: "user@gmail.com",
		Password: "S3cureP4$word",
	}
	
	bodyJSON, err := json.Marshal(body)
	if err != nil {
		t.Error(err)
	}
	
	bodyData := bytes.NewBuffer([]byte(bodyJSON))
	
	request, err := http.NewRequest(http.MethodPost, server.URL, bodyData)
	if err != nil {
		t.Error(err)
	}

	client := http.Client{}

	resp, err := client.Do(request)
	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("status code expected %d but got %d", http.StatusCreated, resp.StatusCode)
	}

}
