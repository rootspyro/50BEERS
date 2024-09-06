package location_test

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
	"github.com/rootspyro/50BEERS/handlers/location"
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

func TestListLocationsForBlog(t *testing.T) {
	connStr, dbName := connString()

	dbclient, err := db.New(connStr)
	if err != nil {
		t.Error(err)
	}

	database := dbclient.Database(dbName)

	repo := repositories.NewLocationRepo(database.Collection("location"))
	srv := services.NewLocationSrv(repo)

	var handler location.LocationHandler = *location.NewLocationHandler(srv)

	server := httptest.NewServer(http.HandlerFunc(handler.ListLocationsForBlog))

	resp, err := http.Get(server.URL)
	if err != nil {
		t.Error(err)
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}

	var response locationListSuccessResponse

	err = json.Unmarshal(b, &response)

	if err != nil {
		t.Error(err)
	}

	// validate response
	if response.Status != "success" {
		t.Errorf("status expected 'success' but got '%s'", response.Status)
	}

	if response.StatusCode != http.StatusOK {
		t.Errorf("status code expected 200 but got %d", response.StatusCode)
	}

	dbclient.Disconnect(context.TODO())
}

type locationListSuccessResponse struct {
	Status     string            `json:"status"`
	StatusCode int               `json:"statusCode"`
	Data       location.LocationsResponse `json:"data"`
}
