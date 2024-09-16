package contact

import (
	"net/http"

	"github.com/rootspyro/50BEERS/config/log"
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

	body := r.Context().Value("body").(services.ContactDTO)

	if err := h.srv.SendContactEmail(body); err != nil {
		log.Error(err.Error())
		parser.SERVER_ERROR(w, "error trying to send email", r.RequestURI)
		return
	}

	parser.JSON(w, parser.SuccessResponse{
		Status: parser.Status.Success,
		StatusCode: http.StatusOK,
		Data: "email was successfully sended",
	})
}
