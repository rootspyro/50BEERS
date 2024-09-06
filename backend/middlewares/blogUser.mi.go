package middlewares

import (
	"context"
	"encoding/json"
	"net/http"
	"regexp"
	"time"

	"github.com/rootspyro/50BEERS/config/parser"
	bloguser "github.com/rootspyro/50BEERS/handlers/blogUser"
	"github.com/rootspyro/50BEERS/services"
)

func evalEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}


func isPasswordSecure(password string) bool {
	// Password should be at least 8 characters long
	if len(password) < 8 {
		return false
	}

	// Check if password contains at least one uppercase letter
	uppercasePattern := `[A-Z]`
	uppercaseRegex := regexp.MustCompile(uppercasePattern)
	if !uppercaseRegex.MatchString(password) {
		return false
	}

	// Check if password contains at least one lowercase letter
	lowercasePattern := `[a-z]`
	lowercaseRegex := regexp.MustCompile(lowercasePattern)
	if !lowercaseRegex.MatchString(password) {
		return false
	}

	// Check if password contains at least one digit
	digitPattern := `[0-9]`
	digitRegex := regexp.MustCompile(digitPattern)
	if !digitRegex.MatchString(password) {
		return false
	}

	// Check if password contains at least one special character
	specialCharPattern := `[!@#\$%\^&\*\(\)\-\+=_{}\[\]:;"'<>,\.\?\/\\|~\]]`
	specialCharRegex := regexp.MustCompile(specialCharPattern)
	if !specialCharRegex.MatchString(password) {
		return false
	}

	return true
}

func PipeLoginBody(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body bloguser.LoginDTO			
	
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			parser.JSON(w, parser.ErrorResponse{
				Status: parser.Status.Error,
				StatusCode: http.StatusBadRequest,
				Error: parser.Error{
					Code: parser.Errors.BAD_REQUEST_BODY.Code,
					Message: parser.Errors.BAD_REQUEST_BODY.Message,
					Details: "body of the request is missing",
					Suggestion: "add the body on json format",
					Path: r.RequestURI,
					Timestamp: parser.Timestamp(),
				},
			})

			return
		}

		if body.User == "" {
			parser.JSON(w, parser.ErrorResponse{
				Status: parser.Status.Error,
				StatusCode: http.StatusBadRequest,
				Error: parser.Error{
					Code: parser.Errors.BAD_REQUEST_BODY.Code,
					Message: parser.Errors.BAD_REQUEST_BODY.Message,
					Details: "user cannot be empty",
					Suggestion: "insert username or email",
					Path: r.RequestURI,
					Timestamp: time.Now().Local(),
				},
			})
			return
		}

		if body.Password == "" {

			parser.JSON(w, parser.ErrorResponse{
				Status: parser.Status.Error,
				StatusCode: http.StatusBadRequest,
				Error: parser.Error{
					Code: parser.Errors.BAD_REQUEST_BODY.Code,
					Message: parser.Errors.BAD_REQUEST_BODY.Message,
					Details: "password cannot be empty",
					Suggestion: "insert the login password",
					Path: r.RequestURI,
					Timestamp: time.Now().Local(),
				},
			})
			return
		}

		// pass the parsed body to the handler by the request context
		ctx := context.WithValue(r.Context(), "body", body)
		next.ServeHTTP(w,r.WithContext(ctx))
	}
}

func PipeNewBlogUserBody(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var body services.BlogUserDTO
		
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			parser.JSON(w, parser.ErrorResponse{
				Status: parser.Status.Error,
				StatusCode: http.StatusBadRequest,
				Error: parser.Error{
					Code: parser.Errors.BAD_REQUEST_BODY.Code,
					Message: parser.Errors.BAD_REQUEST_BODY.Message,
					Details: "body of the request is missing",
					Suggestion: "add the body on json format",
					Path: r.RequestURI,
					Timestamp: time.Now().Local(),
				},
			})
			return
		}

		if len(body.Username) < 4 || len(body.Username) > 30 {
			parser.JSON(w, parser.ErrorResponse{
				Status: parser.Status.Error,
				StatusCode: http.StatusBadRequest,
				Error: parser.Error{
					Code: parser.Errors.BAD_REQUEST_BODY.Code,
					Message: parser.Errors.BAD_REQUEST_BODY.Message,
					Details: "invalid username",
					Suggestion: "username length must be greater than 4 and lower than 30",
					Path: r.RequestURI,
					Timestamp: time.Now().Local(),
				},
			})
			return
		}

		if !evalEmail(body.Email) {
			parser.JSON(w, parser.ErrorResponse{
				Status: parser.Status.Error,
				StatusCode: http.StatusBadRequest,
				Error: parser.Error{
					Code: parser.Errors.BAD_REQUEST_BODY.Code,
					Message: parser.Errors.BAD_REQUEST_BODY.Message,
					Details: "invalid email format",
					Suggestion: "use a valid email",
					Path: r.RequestURI,
					Timestamp: time.Now().Local(),
				},
			})
			return
		}

		if !isPasswordSecure(body.Password) {
			parser.JSON(w, parser.ErrorResponse{
				Status: parser.Status.Error,
				StatusCode: http.StatusBadRequest,
				Error: parser.Error{
					Code: parser.Errors.BAD_REQUEST_BODY.Code,
					Message: parser.Errors.BAD_REQUEST_BODY.Message,
					Details: "invalid password format",
					Suggestion: "password must include at least 8 characters long, one uppercase letter, one lowercase letter, on digit, one special character",
					Path: r.RequestURI,
					Timestamp: time.Now().Local(),
				},
			})
			return
		}

		ctx := context.WithValue(r.Context(), "body", body)
		next.ServeHTTP(w,r.WithContext(ctx))
	}
}

