package bloguser

import (
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

	var subject string
	accessToken, atErr := r.Cookie("access_token")
	refreshToken, rtErr := r.Cookie("refresh_token")

	var token string

	if atErr != nil {
		if rtErr != nil {
			parser.JSON(w, parser.ErrorResponse{
				Status: parser.Status.Error,
				StatusCode: http.StatusUnauthorized,
				Error: parser.Error{
					Code: parser.Errors.UNAUTHORIZED.Code,
					Message: parser.Errors.UNAUTHORIZED.Message,
					Details: "access token is missing",
					Timestamp: parser.Timestamp(),
					Path: r.RequestURI,
				},
			})
			return
		}

		token = refreshToken.Value
	} else {
		token = accessToken.Value
	}

	subject, err := jwt.Decode(token)

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

	// get blog user data
	user, err := h.srv.GetUserByEmail(subject)
	if err != nil {
		if err == mongo.ErrNoDocuments {

			//remove the accessToken cookie
			cookie := http.Cookie {
				Name: "access_token",
				Value: "",
				HttpOnly: true,
				Secure: true,
				SameSite: http.SameSiteNoneMode,
				Path: "/",
				MaxAge: -1, // new cookie that instantly dies
			}

			http.SetCookie(w, &cookie)

			parser.JSON(w, parser.ErrorResponse{
				Status: parser.Status.Error,
				StatusCode: http.StatusUnauthorized,
				Error: parser.Error{
					Code: parser.Errors.UNAUTHORIZED.Code,
					Message: parser.Errors.UNAUTHORIZED.Message,
					Details: "invalid user",
					Path: r.RequestURI,
					Timestamp: parser.Timestamp(),
				},
			})
			return
		}

		parser.SERVER_ERROR(w, "error validating user", r.RequestURI)
		return
	}

	if atErr != nil {
		newAccessToken, err := jwt.SignToken(user.Email)
		if err != nil {
			parser.SERVER_ERROR(w, "error signing new token", r.RequestURI)
			return
		}

		cookie := http.Cookie{
			Name: "access_token",
			Value: newAccessToken,
			HttpOnly: true,
			Secure: true,
			SameSite: http.SameSiteNoneMode,
			Path: "/",
			MaxAge: int(time.Hour.Seconds()),
		}		

		http.SetCookie(w, &cookie)
	}

	parser.JSON(w, parser.SuccessResponse{
		Status: parser.Status.Success,
		StatusCode: http.StatusOK,
		Data: user,
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
		Secure: true,
		SameSite: http.SameSiteNoneMode,
		Path: "/",
		MaxAge: 3600,
	}

	cookkieRefreshToken := http.Cookie {
		Name: "refresh_token",
		Value: refreshToken,
		HttpOnly: true,
		Secure: true,
		SameSite: http.SameSiteNoneMode,
		Path: "/",
		MaxAge: 3600 * 12,
	}

	http.SetCookie(w, &cookkieToken)
	http.SetCookie(w, &cookkieRefreshToken)

	parser.JSON(w, parser.SuccessResponse{
		Status: parser.Status.Success,
		StatusCode: http.StatusOK,
		Data: "logged in",
	})
} 

func(h *BlogUserHandler) Logout(w http.ResponseWriter, r *http.Request) {
	// rewrite access and refresh token cookies
	accessCookie := http.Cookie {
		Name: "access_token",
		Value: "",
		HttpOnly: true,
		Secure: true,
		SameSite: http.SameSiteNoneMode,
		Path: "/",
		MaxAge: -1,
	}

	refreshCookie := http.Cookie {
		Name: "refresh_token",
		Value: "",
		HttpOnly: true,
		Secure: true,
		SameSite: http.SameSiteNoneMode,
		Path: "/",
		MaxAge: -1,
	}

	http.SetCookie(w, &accessCookie)
	http.SetCookie(w, &refreshCookie)

	parser.JSON(w, parser.SuccessResponse{
		Status: parser.Status.Success,
		StatusCode: http.StatusOK,
		Data: "successfull logout",
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
