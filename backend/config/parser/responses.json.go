package parser

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"time"
)

type SuccessResponse struct {
	Status     string `json:"status"`
	StatusCode int    `json:"statusCode"`
	Data       any    `json:"data"`
}

type ErrorResponse struct {
	Status     string `json:"status"`
	StatusCode int    `json:"statusCode"`
	Error      Error  `json:"error"`
}

type Error struct {
	Code       string    `json:"code"`
	Message    string    `json:"message"`
	Details    string    `json:"details"`
	Suggestion string    `json:"suggestion"`
	Timestamp  time.Time `json:"timestamp"`
	Path       string    `json:"path"`
}

type CommonError struct {
	Code       string
	Message    string
	Suggestion string
}

type statusList struct {
	Success string
	Error   string
}

var Status statusList = statusList{
	Success: "success",
	Error:   "error",
}

type errorList struct {
	BAD_REQUEST_QUERY     CommonError
	PATH_NOT_FOUND        CommonError
	INTERNAL_SERVER_ERROR CommonError
}

var Errors errorList = errorList{
	BAD_REQUEST_QUERY: CommonError{
		Code:    "BAD_REQUEST",
		Message: "invalid format for request query",
	},
	PATH_NOT_FOUND: CommonError{
		Code:    "NOT_FOUND",
		Message: "path was not found",
	},
	INTERNAL_SERVER_ERROR: CommonError{
		Code:    "INTERNAL_SERVER_ERROR",
		Message: "error from the data layer",
		Suggestion: "check server status at GET:/api/v1/health",
	},
}

func JSON(w http.ResponseWriter, response interface{}) {
	// Define Content-Type for JSON
	w.Header().Set("Content-Type", "application/json")

	data := reflect.ValueOf(response)

	// If response is not an struct
	if data.Kind() != reflect.Struct {

		errorResp := ErrorResponse{
			Status:     "error",
			StatusCode: http.StatusInternalServerError,
			Error: Error{
				Code:      Errors.INTERNAL_SERVER_ERROR.Code,
				Message:   Errors.INTERNAL_SERVER_ERROR.Message,
				Details:   "error trying to build json response",
				Timestamp: time.Now(),
			},
		}

		jsonResp, _ := json.Marshal(errorResp)

		fmt.Fprintln(w, string(jsonResp))
	}

	var status int
	statusCodeValue := data.FieldByName("StatusCode")

	if !statusCodeValue.IsValid() || statusCodeValue.Kind() != reflect.Int {
		status = 200
	} else {
		status = int(statusCodeValue.Int())
	}

	// Define http status
	w.WriteHeader(status)

	// Format and Send the json response
	jsonBytes, _ := json.Marshal(response)
	fmt.Fprintln(w, string(jsonBytes))
}
