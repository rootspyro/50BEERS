package bloguser

import (
	"fmt"
	"net/http"
	"time"

	"github.com/rootspyro/50BEERS/config/jwt"
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

func(h *BlogUserHandler) ValidateToken(w http.ResponseWriter, r *http.Request) {
	accessToken, err := r.Cookie("access_token")

	if err != nil {

		parser.JSON(w, parser.ErrorResponse{
			Status: parser.Status.Error,
			StatusCode: http.StatusUnauthorized,
			Error: parser.Error{
				Code: parser.Errors.UNAUTHORIZED.Code,
				Message: parser.Errors.UNAUTHORIZED.Message,
				Details: err.Error(),
				Timestamp: parser.Timestamp(),
				Path: r.RequestURI,
			},
		})
		return
	}

	fmt.Println(accessToken)

	parser.JSON(w, parser.SuccessResponse{
		Status: parser.Status.Success,
		StatusCode: http.StatusOK,
		Data: "all ok",
	})
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

	// generate token
	token, err := jwt.SignToken(user.Email)
	if err != nil {
		log.Error(err.Error())
		parser.SERVER_ERROR(w, "token couldn't be generated", r.RequestURI)
		return
	}

	refreshToken, err := jwt.SignRefreshToken(user.Email)
	if err != nil {
		log.Error(err.Error())
		parser.SERVER_ERROR(w, "refresh token couldn't be generated", r.RequestURI)
		return
	}

	// build cookie
	cookkieToken := http.Cookie {
		Name: "access_token",
		Value: token,
		HttpOnly: true,
		Secure: false,
		Path: r.RequestURI,
		MaxAge: int(time.Now().Add(time.Hour * 1).Unix()),
	}

	cookkieRefreshToken := http.Cookie {
		Name: "refresh_token",
		Value: refreshToken,
		HttpOnly: true,
		Secure: false,
		Path: r.RequestURI,
		MaxAge: int(time.Now().Add(time.Hour * 1).Unix()),
	}

	http.SetCookie(w, &cookkieToken)
	http.SetCookie(w, &cookkieRefreshToken)

	parser.JSON(w, parser.SuccessResponse{
		Status: parser.Status.Success,
		StatusCode: http.StatusOK,
		Data: "logged in",
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
