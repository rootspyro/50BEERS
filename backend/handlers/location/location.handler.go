package location

import (
	"net/http"
	"time"

	"github.com/rootspyro/50BEERS/config/parser"
	"github.com/rootspyro/50BEERS/services"
)

type LocationHandler struct {
	srv *services.LocationSrv
}

func NewLocationHandler(srv *services.LocationSrv) *LocationHandler {
	return &LocationHandler{
		srv: srv,
	}
}

func(h *LocationHandler) ListLocationsForBlog(w http.ResponseWriter, r *http.Request) {
	locations, err := h.srv.GetAllLocations()
	if err != nil {
		parser.JSON(w, parser.ErrorResponse{
			Status: parser.Status.Error,
			StatusCode: http.StatusInternalServerError,
			Error: parser.Error{
				Code: parser.Errors.INTERNAL_SERVER_ERROR.Code,
				Message: parser.Errors.INTERNAL_SERVER_ERROR.Message,
				Details: "error getting locations data",
				Suggestion: parser.Errors.INTERNAL_SERVER_ERROR.Suggestion,
				Path: r.RequestURI,
				Timestamp: time.Now().Local(),
			},
		})	
		return
	}

	parser.JSON(w, parser.SuccessResponse{
		Status: parser.Status.Success,
		StatusCode: http.StatusOK,
		Data: LocationsResponse{
			ItemsFound: len(locations),
			Items: locations,
		},
	})
}
