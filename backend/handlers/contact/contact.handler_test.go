package contact_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/rootspyro/50BEERS/SDKs/mailtrap"
	"github.com/rootspyro/50BEERS/config/parser"
	"github.com/rootspyro/50BEERS/handlers/contact"
	"github.com/rootspyro/50BEERS/middlewares"
	"github.com/rootspyro/50BEERS/services"
)

func TestEmailFromBlogSuccess(t *testing.T) {

	godotenv.Load("../../.env")

	sdk := mailtrap.New(
		os.Getenv("MAILTRAP_HOST"),
		os.Getenv("MAILTRAP_API_TOKEN"),
		os.Getenv("MAILTRAP_DOMAIN_EMAIL"),
	)

	handler := contact.NewContactHandler(services.NewContactSrv(
		os.Getenv("AUTHOR_EMAIL"),
		sdk,
	))

	// build the testing server
	server := httptest.NewServer(
		http.Handler(middlewares.PipeContactBody(handler.EmailFromBlog)),
	)

	// build the request
	body := services.ContactDTO{
		Name: "rootspyro",
		Email: "rootspyro@gmail.com",
		Message: "I love bock beers",
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

	// run the request
	client := http.Client{}

	resp, err := client.Do(request)
	if err != nil {
		t.Error(err)
	}

	// evaluate api response
	if resp.StatusCode != http.StatusOK {
		t.Errorf("status code expected %d but got %d", http.StatusOK, resp.StatusCode)
	}

	// get response body
	var respBody parser.SuccessResponse
	
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}

	err = json.Unmarshal(b, &respBody)
	if err != nil {
		t.Error(err)
	}

	if respBody.Status != parser.Status.Success {
		t.Errorf("status expected '%s' but got '%s'", parser.Status.Success, respBody.Status)
	}

	if resp.StatusCode != respBody.StatusCode {
		t.Errorf("status code from response %d doesn't match with body status code %d", resp.StatusCode, respBody.StatusCode)
	}
}

func TestEmailFromBlogBadEmail(t *testing.T) {

	godotenv.Load("../../.env")

	sdk := mailtrap.New(
		os.Getenv("MAILTRAP_HOST"),
		os.Getenv("MAILTRAP_API_TOKEN"),
		os.Getenv("MAILTRAP_DOMAIN_EMAIL"),
	)

	handler := contact.NewContactHandler(services.NewContactSrv(
		os.Getenv("AUTHOR_EMAIL"),
		sdk,
	))

	// build the testing server
	server := httptest.NewServer(
		http.Handler(middlewares.PipeContactBody(handler.EmailFromBlog)),
	)

	// build the request
	body := services.ContactDTO{
		Name: "rootspyro",
		Email: "rootspyro@gmail", // bad format for email
		Message: "I love bock beers",
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

	// run the request
	client := http.Client{}

	resp, err := client.Do(request)
	if err != nil {
		t.Error(err)
	}

	// evaluate api response
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("status code expected %d but got %d", http.StatusBadRequest, resp.StatusCode)
	}

	// get response body
	var respBody parser.ErrorResponse
	
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}

	err = json.Unmarshal(b, &respBody)
	if err != nil {
		t.Error(err)
	}

	if respBody.Status != parser.Status.Error {
		t.Errorf("status expected '%s' but got '%s'", parser.Status.Success, respBody.Status)
	}

	if resp.StatusCode != respBody.StatusCode {
		t.Errorf("status code from response %d doesn't match with body status code %d", resp.StatusCode, respBody.StatusCode)
	}

	if respBody.Error.Code != parser.Errors.BAD_REQUEST_BODY.Code {
		t.Errorf("error code expected '%s' but got '%s'", parser.Errors.BAD_REQUEST_BODY.Code, respBody.Error.Code)
	}

	if respBody.Error.Details != "invalid email format" {
		t.Errorf("error details expected 'invalid email format' but got '%s'", respBody.Error.Details)
	}
}

func TestEmailFromBlogMessageMinLength(t *testing.T) {

	godotenv.Load("../../.env")

	sdk := mailtrap.New(
		os.Getenv("MAILTRAP_HOST"),
		os.Getenv("MAILTRAP_API_TOKEN"),
		os.Getenv("MAILTRAP_DOMAIN_EMAIL"),
	)

	handler := contact.NewContactHandler(services.NewContactSrv(
		os.Getenv("AUTHOR_EMAIL"),
		sdk,
	))

	// build the testing server
	server := httptest.NewServer(
		http.Handler(middlewares.PipeContactBody(handler.EmailFromBlog)),
	)

	// build the request
	body := services.ContactDTO{
		Name: "rootspyro",
		Email: "rootspyro@gmail.com", 
		Message: "I lo", // message is too small
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

	// run the request
	client := http.Client{}

	resp, err := client.Do(request)
	if err != nil {
		t.Error(err)
	}

	// evaluate api response
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("status code expected %d but got %d", http.StatusBadRequest, resp.StatusCode)
	}

	// get response body
	var respBody parser.ErrorResponse
	
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}

	err = json.Unmarshal(b, &respBody)
	if err != nil {
		t.Error(err)
	}

	if respBody.Status != parser.Status.Error {
		t.Errorf("status expected '%s' but got '%s'", parser.Status.Success, respBody.Status)
	}

	if resp.StatusCode != respBody.StatusCode {
		t.Errorf("status code from response %d doesn't match with body status code %d", resp.StatusCode, respBody.StatusCode)
	}

	if respBody.Error.Code != parser.Errors.BAD_REQUEST_BODY.Code {
		t.Errorf("error code expected '%s' but got '%s'", parser.Errors.BAD_REQUEST_BODY.Code, respBody.Error.Code)
	}

	if respBody.Error.Details != "the message must be at least 5 characters long" {
		t.Errorf("unexpected error details message: %s", respBody.Error.Details)
	}
}

func TestEmailFromBlogMessageMaxLength(t *testing.T) {

	godotenv.Load("../../.env")

	sdk := mailtrap.New(
		os.Getenv("MAILTRAP_HOST"),
		os.Getenv("MAILTRAP_API_TOKEN"),
		os.Getenv("MAILTRAP_DOMAIN_EMAIL"),
	)

	handler := contact.NewContactHandler(services.NewContactSrv(
		os.Getenv("AUTHOR_EMAIL"),
		sdk,
	))

	// build the testing server
	server := httptest.NewServer(
		http.Handler(middlewares.PipeContactBody(handler.EmailFromBlog)),
	)

	// build the request
	body := services.ContactDTO{
		Name: "rootspyro",
		Email: "rootspyro@gmail.com", 
		// message is to large
		Message: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
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

	// run the request
	client := http.Client{}

	resp, err := client.Do(request)
	if err != nil {
		t.Error(err)
	}

	// evaluate api response
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("status code expected %d but got %d", http.StatusBadRequest, resp.StatusCode)
	}

	// get response body
	var respBody parser.ErrorResponse
	
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}

	err = json.Unmarshal(b, &respBody)
	if err != nil {
		t.Error(err)
	}

	if respBody.Status != parser.Status.Error {
		t.Errorf("status expected '%s' but got '%s'", parser.Status.Success, respBody.Status)
	}

	if resp.StatusCode != respBody.StatusCode {
		t.Errorf("status code from response %d doesn't match with body status code %d", resp.StatusCode, respBody.StatusCode)
	}

	if respBody.Error.Code != parser.Errors.BAD_REQUEST_BODY.Code {
		t.Errorf("error code expected '%s' but got '%s'", parser.Errors.BAD_REQUEST_BODY.Code, respBody.Error.Code)
	}

	if respBody.Error.Details != "the message cannot exceed 300 characters" {
		t.Errorf("unexpected error details message: %s", respBody.Error.Details)
	}
}
