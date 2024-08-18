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
	Code      string    `json:"code"`
	Message   string    `json:"message"`
	Details   string    `json:"details"`
	Timestamp time.Time `json:"timestamp"`
	Path      string    `json:"path"`
}

type CommonError struct {
	Code    string
	Message string
}

type ErrorList struct {
	PATH_NOT_FOUND        CommonError
	INTERNAL_SERVER_ERROR CommonError
}

var ERRORS ErrorList = ErrorList{
	PATH_NOT_FOUND: CommonError{
		Code:    "PATH_NOT_FOUND",
		Message: "path was not found",
	},
	INTERNAL_SERVER_ERROR: CommonError{
		Code:    "INTERNAL_SERVER_ERROR",
		Message: "error from the data layer",
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
				Code:    ERRORS.INTERNAL_SERVER_ERROR.Code,
				Message: ERRORS.INTERNAL_SERVER_ERROR.Message,
				Details: "error trying to build json response",
				Timestamp:    time.Now(),
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
