package drinks

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
	"github.com/rootspyro/50BEERS/middlewares"
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

func buildHandler(database *mongo.Database) DrinkHandler {		
	// required repositories
	countryRepo := repositories.NewCountriesRepo(database.Collection("country"))
	locationRepo := repositories.NewLocationRepo(database.Collection("location"))
	drinksRepo := repositories.NewDrinksRepo(database.Collection("drink"))

	// build service
	srv := services.NewDrinkSrv(
		countryRepo,
		locationRepo,
		drinksRepo,
	)

	return *NewDrinkHandler(srv)
}


func TestCountDrinksForBlog(t *testing.T) {
	connStr, dbName := connString()

	dbClient, err := db.New(connStr)
	if err != nil {
		t.Error(err)
	}

	defer dbClient.Disconnect(context.TODO())

	handler := buildHandler(dbClient.Database(dbName))

	// build testing server
	server := httptest.NewServer(http.HandlerFunc(handler.CountDrinks))

	// build http Get Request
	resp, err := http.Get(server.URL)
	if err != nil {
		t.Error(err)
	}

	// read response body
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}

	var result successResponse

	err = json.Unmarshal(b, &result)
	if err != nil {
		t.Error(err)
	}

	// eval response
	if resp.StatusCode != result.StatusCode {
		t.Error("the statusCode of the response and the statusCode of the body do not match")
	}

	if result.StatusCode != http.StatusOK {
		t.Errorf("status code expected 200 but got %d", result.StatusCode)
	}

	if result.Status != "success" {
		t.Errorf("response status expected 'success' but got '%s'", result.Status)
	}
}

func TestGetAllDrinksDefault(t *testing.T) {
	connStr, dbName := connString()

	dbClient, err := db.New(connStr)
	if err != nil {
		t.Error(err)
	}

	defer dbClient.Disconnect(context.TODO())

	handler := buildHandler(dbClient.Database(dbName))

	// build testing server
	server := httptest.NewServer(http.HandlerFunc(middlewares.ValidateDrinksBlogFilters(handler.ListDrinksForBlog)))

	// build http Get Request
	resp, err := http.Get(server.URL)
	if err != nil {
		t.Error(err)
	}

	// build response body
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}

	var result drinksSuccessResponse

	err = json.Unmarshal(b, &result)
	if err != nil {
		t.Error(err)
	}

	// eval status 
	if resp.StatusCode != result.StatusCode {
		t.Error("the statusCode of the response and the statusCode of the body do not match")
	}

	if result.StatusCode != http.StatusOK {
		t.Errorf("status code expected 200 but got %d", result.StatusCode)
	}

	if result.Status != "success" {
		t.Errorf("response status expected 'success' but got '%s'", result.Status)
	}

	// eval pagination
	pagination := result.Data.Pagination

	if pagination.Page > 1 {
		t.Errorf("page expected 1 but got %d", pagination.Page)
	} 

	if pagination.PageSize != 10 {
		t.Errorf("page size expected 10 but got %d", pagination.PageSize)
	}

	// eval default filters
	if result.Data.FiltersApplied.SortBy != "created_at" {
		t.Errorf("sortBy filter expected 'created_at' but got %s", result.Data.FiltersApplied.SortBy)
	}

	if result.Data.FiltersApplied.Direction != "desc" {
		t.Errorf("direction expected 'desc' but got '%s'", result.Data.FiltersApplied.Direction)
	}
}

type successResponse struct {
	Status     string `json:"status"`
	StatusCode int `json:"statusCode"`
}

type drinksSuccessResponse struct {
	Status     string `json:"status"`
	StatusCode int `json:"statusCode"`
	Data       DrinksResponse
}
