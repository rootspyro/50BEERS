package contact

import (
	"net/http"

	"github.com/rootspyro/50BEERS/config/parser"
	"github.com/rootspyro/50BEERS/services"
)

type ContactHadler struct {
	srv *services.ContactSrv
}

func NewContactHandler(srv *services.ContactSrv) *ContactHadler {
	return &ContactHadler{
		srv: srv,
	}
}

func(h *ContactHadler) EmailFromBlog(w http.ResponseWriter, r *http.Request) {

	if err := h.srv.SendEmail(); err != nil {
		parser.SERVER_ERROR(w, "error trying to send email", r.RequestURI)
		return
	}

	parser.JSON(w, parser.SuccessResponse{
		Status: parser.Status.Success,
		StatusCode: http.StatusOK,
		Data: "email was successfully sended",
	})
}
