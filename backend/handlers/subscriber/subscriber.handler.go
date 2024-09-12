package subscriber

import (
	"fmt"
	"net/http"

	"github.com/rootspyro/50BEERS/config/parser"
	"github.com/rootspyro/50BEERS/services"
)

type SubscriberHandler struct {
	srv *services.SubscriberSrv
}

func NewSubscriberHandler(srv *services.SubscriberSrv) *SubscriberHandler {
	return &SubscriberHandler{
		srv: srv,
	}
}

func(h *SubscriberHandler) NewSub(w http.ResponseWriter, r *http.Request) {

	body := r.Context().Value("body").(services.SubscriberDTO)
	fmt.Println(body)

	parser.JSON(w, parser.SuccessResponse {
		Status: parser.Status.Success,
		StatusCode: http.StatusCreated,
		Data: "Subscription confirmed. Thank you for your support",
	})
}
