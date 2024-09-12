package subscriber_test

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
	"github.com/rootspyro/50BEERS/handlers/subscriber"
	"github.com/rootspyro/50BEERS/services"
	"go.mongodb.org/mongo-driver/mongo"
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

func buildHandler(database *mongo.Database) subscriber.SubscriberHandler {
	repo := repositories.NewSubscriberRepo(database.Collection("subscriber"))
	srv := services.NewSubscriberSrv(repo)
	handler := subscriber.NewSubscriberHandler(srv)

	return *handler
}

func TestNewSubscriberSuccess(t *testing.T) {
	connStr, dbName := connString()

	dbClient, err := db.New(connStr)
	if err != nil {
		t.Error(err)
	}

	database := dbClient.Database(dbName)

	handler := buildHandler(database)	
	
	// build the testing server
	server := httptest.NewServer(http.HandlerFunc(handler.NewSub))

	// build the request
	body := subscriber.NewSubscriberDTO {
		Email: "subscriberl@mail.com",
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
