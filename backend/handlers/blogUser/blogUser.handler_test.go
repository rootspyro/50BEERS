package bloguser_test

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
	bloguser "github.com/rootspyro/50BEERS/handlers/blogUser"
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
	handler := bloguser.NewBlogUserHandler(srv)

	// build the testing server
	server := httptest.NewServer(http.HandlerFunc(middlewares.PipeNewBlogUserBody(handler.SignUp)))

	// do the request
	body := services.BlogUserDTO{
		Username: "user",
		Email:    "user@gmail.com",
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

	var result signUpSuccessResponse

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}

	err = json.Unmarshal(b, &result)

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("status code expected %d but got %d", http.StatusCreated, resp.StatusCode)
	}

	if result.StatusCode != resp.StatusCode {
		t.Errorf("status code from response %d doesn't match with body status code %d", resp.StatusCode, result.StatusCode)
	}

	if result.Status != parser.Status.Success {
		t.Errorf("status expected '%s' but got '%s'", parser.Status.Success, result.Status)
	}
}

func TestSignUpBadEmail(t *testing.T) {
	connStr, dbName := connString()

	dbClient, err := db.New(connStr)
	if err != nil {
		t.Error(err)
	}

	database := dbClient.Database(dbName)

	repo := repositories.NewBlogUserRepo(database.Collection("blogUser"))
	srv := services.NewBlogUserSrv(repo)
	handler := bloguser.NewBlogUserHandler(srv)

	// build the testing server
	server := httptest.NewServer(http.HandlerFunc(middlewares.PipeNewBlogUserBody(handler.SignUp)))

	// do the request
	body := services.BlogUserDTO{
		Username: "user",
		Email:    "user@gmail",
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

	var result parser.ErrorResponse 

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}

	err = json.Unmarshal(b, &result)

	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("status code expected %d but got %d", http.StatusBadRequest, resp.StatusCode)
	}

	if result.StatusCode != resp.StatusCode {
		t.Errorf("status code from response %d doesn't match with body status code %d", resp.StatusCode, result.StatusCode)
	}

	if result.Status != parser.Status.Error {
		t.Errorf("status expected '%s' but got '%s'", parser.Status.Error, result.Status)
	}

	if result.Error.Code != parser.Errors.BAD_REQUEST_BODY.Code {
		t.Errorf("error code expected '%s' but got '%s'", parser.Errors.BAD_REQUEST_BODY.Code, result.Error.Code)
	}
}

func TestSignUpInsecurePassword(t *testing.T) {
	connStr, dbName := connString()

	dbClient, err := db.New(connStr)
	if err != nil {
		t.Error(err)
	}

	database := dbClient.Database(dbName)

	repo := repositories.NewBlogUserRepo(database.Collection("blogUser"))
	srv := services.NewBlogUserSrv(repo)
	handler := bloguser.NewBlogUserHandler(srv)

	// build the testing server
	server := httptest.NewServer(http.HandlerFunc(middlewares.PipeNewBlogUserBody(handler.SignUp)))

	// do the request
	body := services.BlogUserDTO{
		Username: "user",
		Email:    "user@gmail",
		Password: "us3radm1n#", // for this case we are not using uppercase letters
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

	var result parser.ErrorResponse 

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}

	err = json.Unmarshal(b, &result)

	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("status code expected %d but got %d", http.StatusBadRequest, resp.StatusCode)
	}

	if result.StatusCode != resp.StatusCode {
		t.Errorf("status code from response %d doesn't match with body status code %d", resp.StatusCode, result.StatusCode)
	}

	if result.Status != parser.Status.Error {
		t.Errorf("status expected '%s' but got '%s'", parser.Status.Error, result.Status)
	}

	if result.Error.Code != parser.Errors.BAD_REQUEST_BODY.Code {
		t.Errorf("error code expected '%s' but got '%s'", parser.Errors.BAD_REQUEST_BODY.Code, result.Error.Code)
	}
}

func TestSignUpFromSiteConflictUserExists(t *testing.T) {
	connStr, dbName := connString()

	dbClient, err := db.New(connStr)
	if err != nil {
		t.Error(err)
	}

	database := dbClient.Database(dbName)

	repo := repositories.NewBlogUserRepo(database.Collection("blogUser"))
	srv := services.NewBlogUserSrv(repo)
	handler := bloguser.NewBlogUserHandler(srv)

	// build the testing server
	server := httptest.NewServer(http.HandlerFunc(middlewares.PipeNewBlogUserBody(handler.SignUp)))

	// do the request
	body := services.BlogUserDTO{
		Username: "user",
		Email:    "user@gmail.com",
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

	var result parser.ErrorResponse 

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}

	err = json.Unmarshal(b, &result)

	if resp.StatusCode != http.StatusConflict {
		t.Errorf("status code expected %d but got %d", http.StatusConflict, resp.StatusCode)
	}

	if result.StatusCode != resp.StatusCode {
		t.Errorf("status code from response %d doesn't match with body status code %d", resp.StatusCode, result.StatusCode)
	}

	if result.Status != parser.Status.Error {
		t.Errorf("status expected '%s' but got '%s'", parser.Status.Success, result.Status)
	}

	if result.Error.Code != parser.Errors.CONFLICT.Code {
		t.Errorf("error code expected '%s' but got '%s'", parser.Errors.CONFLICT.Code, result.Error.Code)
	}
}

type signUpSuccessResponse struct {
	Status     string `json:"status"`
	StatusCode int    `json:"statusCode"`
	Data       services.BlogUser
}
