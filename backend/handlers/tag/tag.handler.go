package tag

import (
	"net/http"
	"time"

	"github.com/rootspyro/50BEERS/config/parser"
	"github.com/rootspyro/50BEERS/services"
)

type TagHandler struct {
	srv *services.TagSrv
}

func NewTagHandler(srv *services.TagSrv) *TagHandler {
	return &TagHandler{
		srv: srv,
	}
}

func(h *TagHandler) ListCategoriesForBlog(w http.ResponseWriter, r *http.Request) {

	lang := r.Context().Value("lang").(string)

	tags, err := h.srv.GetAllTags(lang)

	if err != nil {
		parser.JSON(w, parser.ErrorResponse{
			Status: parser.Status.Error,
			StatusCode:  http.StatusInternalServerError,
			Error: parser.Error{
				Code: parser.Errors.INTERNAL_SERVER_ERROR.Code,
				Message: parser.Errors.INTERNAL_SERVER_ERROR.Message,	
				Details: "error getting the categories",
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
		Data: TagsResponse{
			ItemsFound: len(tags),
			Items: tags,
		},
	})
}
