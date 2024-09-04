package bloguser

import (
	"net/http"
	"time"

	"github.com/rootspyro/50BEERS/config/log"
	"github.com/rootspyro/50BEERS/config/parser"
	"github.com/rootspyro/50BEERS/services"
)

type BlogUserHandler struct {
	srv *services.BlogUserSrv
}

func NewBlogUserHandler(srv *services.BlogUserSrv) *BlogUserHandler {
	return &BlogUserHandler{
		srv: srv,	
	}
}

func(h *BlogUserHandler) SignUp(w http.ResponseWriter, r *http.Request) {

	body := r.Context().Value("body").(services.BlogUserDTO)

	newUser, err := h.srv.NewUserFromSite(body)	

	if err != nil {
		log.Error(err.Error())
		parser.JSON(w, parser.ErrorResponse{
			Status: parser.Status.Error,
			StatusCode: http.StatusInternalServerError,
			Error: parser.Error{
				Code: parser.Errors.INTERNAL_SERVER_ERROR.Code,
				Message: parser.Errors.INTERNAL_SERVER_ERROR.Message,
				Details: "something went wrong creating the new user",
				Suggestion: "",
				Path: r.RequestURI,
				Timestamp: time.Now().Local(),
			},
		})

		return
	}
	
	parser.JSON(w, parser.SuccessResponse{
		Status: parser.Status.Success,	
		StatusCode: http.StatusCreated,
		Data: newUser,
	})
}
