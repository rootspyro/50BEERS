package tag_test

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/rootspyro/50BEERS/db"
	"github.com/rootspyro/50BEERS/db/repositories"
	"github.com/rootspyro/50BEERS/handlers/tag"
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

func TestListCategoriesForBlog(t *testing.T) {
	// build handler
	connStr, dbName := connString()

	dbClient, err := db.New(connStr)
	if err != nil {
		t.Error(err)
	}

	defer dbClient.Disconnect(context.TODO())

	database := dbClient.Database(dbName)

	repo := repositories.NewTagRepo(database.Collection("tag"))
	srv := services.NewTagSrv(repo)

	handler := tag.NewTagHandler(srv)

	// build testing server
	handlerFunc := middlewares.LangHeader(handler.ListCategoriesForBlog)
	server := httptest.NewServer(http.HandlerFunc(handlerFunc))

	// make http GET request
	resp, err := http.Get(server.URL)
	if err != nil {
		t.Error(err)
	}

	// read body

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}

	var result tagSuccessResponse

	err = json.Unmarshal(b, &result)
	if err != nil {
		t.Error(err)
	}

	// eval response
	if result.Status != "success" {
		t.Errorf("expected 'success' but got '%s'", result.Status)
	}

	if result.StatusCode != http.StatusOK {
		t.Errorf("status code expected 200 but got %d", result.StatusCode)
	}
}

type tagSuccessResponse struct {
	Status     string `json:"status"`
	StatusCode int    `json:"statusCode"`
	Data       any    `json:"data"`
}
