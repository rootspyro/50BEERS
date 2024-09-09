package bloguser

import (
	"net/http"

	"github.com/rootspyro/50BEERS/config/log"
	"github.com/rootspyro/50BEERS/config/parser"
	"github.com/rootspyro/50BEERS/services"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type BlogUserHandler struct {
	srv *services.BlogUserSrv
}

func NewBlogUserHandler(srv *services.BlogUserSrv) *BlogUserHandler {
	return &BlogUserHandler{
		srv: srv,	
	}
}

func(h *BlogUserHandler) Login(w http.ResponseWriter, r *http.Request) {

	// get body
	body := r.Context().Value("body").(LoginDTO)

	invalidCredentialsError := parser.ErrorResponse {
		Status: parser.Status.Error,
		StatusCode: http.StatusUnauthorized,
		Error: parser.Error{
			Code: parser.Errors.UNAUTHORIZED.Code,
			Message: parser.Errors.UNAUTHORIZED.Message,
			Details: "incorrect user or password",
			Suggestion: "check the credentials or try to sign up",
			Path: r.RequestURI,
			Timestamp: parser.Timestamp(),
		},
	}

	// get user data
	user, err := h.srv.GetUserForLogin(body.User)
	if err != nil {
		if err == mongo.ErrNoDocuments {

			parser.JSON(w, invalidCredentialsError)
			return
		}
		
		log.Error(err.Error())
		parser.SERVER_ERROR(w, "error trying to get user", r.RequestURI)
		return
	}

	// evaluate password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		parser.JSON(w, invalidCredentialsError)
		return
	}

	// build session cookie
	cookkie := http.Cookie {
		Name: "token",
		Value: "hello world from the beer paradise",
		HttpOnly: true,
		Secure: false,
	}

	http.SetCookie(w, &cookkie)

	parser.JSON(w, parser.SuccessResponse{
		Status: parser.Status.Success,
		StatusCode: http.StatusOK,
		Data: "successfull login",
	})
} 

func(h *BlogUserHandler) SignUp(w http.ResponseWriter, r *http.Request) {

	body := r.Context().Value("body").(services.BlogUserDTO)

	// validate if user already exists
	_, err := h.srv.GetUserByUsername(body.Username)
	if err == nil {
		parser.JSON(w, parser.ErrorResponse{
			Status: parser.Status.Error,
			StatusCode: http.StatusConflict,
			Error: parser.Error{
				Code: parser.Errors.CONFLICT.Code,
				Message: parser.Errors.CONFLICT.Message,
				Details: "username is already taken",
				Suggestion: "try another username",
				Path: r.RequestURI,
				Timestamp: parser.Timestamp(),
			},
		})
		return

	} else {
		if err != mongo.ErrNoDocuments {
			log.Error(err.Error())
			parser.SERVER_ERROR(w, "error trying to validate username", r.RequestURI)
			return
		}
	}

	_, err = h.srv.GetUserByEmail(body.Email)

	if err == nil {
		parser.JSON(w, parser.ErrorResponse{
			Status: parser.Status.Error,
			StatusCode: http.StatusConflict,
			Error: parser.Error{
				Code: parser.Errors.CONFLICT.Code,
				Message: parser.Errors.CONFLICT.Message,
				Details: "there is already an account with this email",
				Suggestion: "Try to login or use another email",
				Path: r.RequestURI,
				Timestamp: parser.Timestamp(),
			},
		})
		return

	} else {
		if err != mongo.ErrNoDocuments {
			log.Error(err.Error())
			parser.SERVER_ERROR(w, "error trying to validate email", r.RequestURI)
			return
		}
	}

	newUser, err := h.srv.NewUserFromSite(body)	

	if err != nil {
		log.Error(err.Error())
		parser.SERVER_ERROR(w, "error trying to create a new user", r.RequestURI)
		return
	}
	
	parser.JSON(w, parser.SuccessResponse{
		Status: parser.Status.Success,	
		StatusCode: http.StatusCreated,
		Data: newUser,
	})
}
