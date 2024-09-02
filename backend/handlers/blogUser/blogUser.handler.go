package bloguser

import (
	"net/http"

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
	
	parser.JSON(w, parser.SuccessResponse{
		Status: parser.Status.Success,	
		StatusCode: http.StatusOK,
		Data: "Success",
	})
}
