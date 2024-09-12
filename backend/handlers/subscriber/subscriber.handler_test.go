package subscriber_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/rootspyro/50BEERS/config/parser"
	"github.com/rootspyro/50BEERS/db"
	"github.com/rootspyro/50BEERS/db/repositories"
	"github.com/rootspyro/50BEERS/handlers/subscriber"
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
	server := httptest.NewServer(http.HandlerFunc(middlewares.PipeSubscriberBody(handler.NewSub)))

	// build the request
	body := services.SubscriberDTO{
		Email: "subscriber@mail.com",
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

	var result parser.SuccessResponse

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}

	err = json.Unmarshal(b, &result)
	if err != nil {
		t.Error(err)
	}

	if result.Status != parser.Status.Success {
		t.Errorf("status expected '%s' but got '%s'", parser.Status.Success, result.Status)
	}

	if result.StatusCode != resp.StatusCode {
		t.Errorf("status code from response %d doesn't match with body status code %d", resp.StatusCode, result.StatusCode)
	}
}

func TestNewSubscriberBadEmail(t *testing.T) {
	connStr, dbName := connString()

	dbClient, err := db.New(connStr)
	if err != nil {
		t.Error(err)
	}

	database := dbClient.Database(dbName)

	handler := buildHandler(database)

	// build the testing server
	server := httptest.NewServer(http.HandlerFunc(middlewares.PipeSubscriberBody(handler.NewSub)))

	// build the request
	body := services.SubscriberDTO{
		Email: "bad.email.format",
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

	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("status code expected %d but got %d", http.StatusBadRequest, resp.StatusCode)
	}

	var result parser.ErrorResponse

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}

	err = json.Unmarshal(b, &result)
	if err != nil {
		t.Error(err)
	}

	if result.Status != parser.Status.Error {
		t.Errorf("status expected '%s' but got '%s'", parser.Status.Error, result.Status)
	}

	if result.StatusCode != resp.StatusCode {
		t.Errorf("status code from response %d doesn't match with body status code %d", resp.StatusCode, result.StatusCode)
	}

	if result.Error.Code != parser.Errors.BAD_REQUEST_BODY.Code {
		t.Errorf("error code expected '%s' but got '%s'", parser.Errors.BAD_REQUEST_BODY.Code, result.Error.Code)
	}

}

func TestNewSubscriberConflict(t *testing.T) {
	connStr, dbName := connString()

	dbClient, err := db.New(connStr)
	if err != nil {
		t.Error(err)
	}

	database := dbClient.Database(dbName)

	handler := buildHandler(database)

	// build the testing server
	server := httptest.NewServer(http.HandlerFunc(middlewares.PipeSubscriberBody(handler.NewSub)))

	// build the request
	body := services.SubscriberDTO{
		Email: "subscriber@mail.com",
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

	if resp.StatusCode != http.StatusConflict {
		t.Errorf("status code expected %d but got %d", http.StatusConflict, resp.StatusCode)
	}

	var result parser.ErrorResponse

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}

	err = json.Unmarshal(b, &result)
	if err != nil {
		t.Error(err)
	}

	if result.Status != parser.Status.Error {
		t.Errorf("status expected '%s' but got '%s'", parser.Status.Error, result.Status)
	}

	if result.StatusCode != resp.StatusCode {
		t.Errorf("status code from response %d doesn't match with body status code %d", resp.StatusCode, result.StatusCode)
	}

	if result.Error.Code != parser.Errors.CONFLICT.Code {
		t.Errorf("error code expected '%s' but got '%s'", parser.Errors.CONFLICT.Code, result.Error.Code)
	}

}

func TestRemoveSubscriberSuccess(t *testing.T) {
	connStr, dbName := connString()

	dbClient, err := db.New(connStr)
	if err != nil {
		t.Error(err)
	}

	database := dbClient.Database(dbName)

	handler := buildHandler(database)

	// build the testing server
	server := httptest.NewServer(http.HandlerFunc(middlewares.PipeSubscriberBody(handler.RemoveSubscriber)))

	// build the request
	body := services.SubscriberDTO{
		Email: "subscriber@mail.com",
	}

	bodyJSON, err := json.Marshal(body)
	if err != nil {
		t.Error(err)
	}

	bodyData := bytes.NewBuffer([]byte(bodyJSON))

	request, err := http.NewRequest(http.MethodDelete, server.URL, bodyData)
	if err != nil {
		t.Error(err)
	}

	client := http.Client{}

	resp, err := client.Do(request)
	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("status code expected %d but got %d", http.StatusOK, resp.StatusCode)
	}

	var result parser.SuccessResponse

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}

	err = json.Unmarshal(b, &result)
	if err != nil {
		t.Error(err)
	}

	if result.Status != parser.Status.Success {
		t.Errorf("status expected '%s' but got '%s'", parser.Status.Success, result.Status)
	}

	if result.StatusCode != resp.StatusCode {
		t.Errorf("status code from response %d doesn't match with body status code %d", resp.StatusCode, result.StatusCode)
	}
}

func TestRemoveSubscriberBadEmail(t *testing.T) {
	connStr, dbName := connString()

	dbClient, err := db.New(connStr)
	if err != nil {
		t.Error(err)
	}

	database := dbClient.Database(dbName)

	handler := buildHandler(database)

	// build the testing server
	server := httptest.NewServer(http.HandlerFunc(middlewares.PipeSubscriberBody(handler.RemoveSubscriber)))

	// build the request
	body := services.SubscriberDTO{
		Email: "bad.email.format",
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

	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("status code expected %d but got %d", http.StatusBadRequest, resp.StatusCode)
	}

	var result parser.ErrorResponse

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}

	err = json.Unmarshal(b, &result)
	if err != nil {
		t.Error(err)
	}

	if result.Status != parser.Status.Error {
		t.Errorf("status expected '%s' but got '%s'", parser.Status.Error, result.Status)
	}

	if result.StatusCode != resp.StatusCode {
		t.Errorf("status code from response %d doesn't match with body status code %d", resp.StatusCode, result.StatusCode)
	}

	if result.Error.Code != parser.Errors.BAD_REQUEST_BODY.Code {
		t.Errorf("error code expected '%s' but got '%s'", parser.Errors.BAD_REQUEST_BODY.Code, result.Error.Code)
	}

}

func TestRemoveSubscriberNotFound(t *testing.T) {
	connStr, dbName := connString()

	dbClient, err := db.New(connStr)
	if err != nil {
		t.Error(err)
	}

	database := dbClient.Database(dbName)

	handler := buildHandler(database)

	// build the testing server
	server := httptest.NewServer(http.HandlerFunc(middlewares.PipeSubscriberBody(handler.RemoveSubscriber)))

	// build the request
	body := services.SubscriberDTO{
		Email: "subscriber@mail.com",
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

	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("status code expected %d but got %d", http.StatusNotFound, resp.StatusCode)
	}

	var result parser.ErrorResponse

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}

	err = json.Unmarshal(b, &result)
	if err != nil {
		t.Error(err)
	}

	if result.Status != parser.Status.Error {
		t.Errorf("status expected '%s' but got '%s'", parser.Status.Error, result.Status)
	}

	if result.StatusCode != resp.StatusCode {
		t.Errorf("status code from response %d doesn't match with body status code %d", resp.StatusCode, result.StatusCode)
	}

	if result.Error.Code != parser.Errors.NOT_FOUND.Code {
		t.Errorf("error code expected '%s' but got '%s'", parser.Errors.NOT_FOUND.Code, result.Error.Code)
	}

}
