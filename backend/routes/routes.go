package routes

import (
	"fmt"
	"net/http"
	"time"

	"github.com/rootspyro/50BEERS/config/parser"
	"github.com/rootspyro/50BEERS/middlewares"
)

func Init() *http.ServeMux {

	var router http.ServeMux = http.ServeMux{}

	var childRouters []Router = []Router{
		HealthRouter,
	}

	BuildRoutes(childRouters, "/api/v1", &router)

	// 404 - PATH NOT FOUND
	router.HandleFunc("/", middlewares.Logger(func(w http.ResponseWriter, r *http.Request) {

		parser.JSON(w, parser.ErrorResponse{
			Status: "error",
			StatusCode: http.StatusNotFound,
			Error: parser.Error{
				Code: parser.ERRORS.PATH_NOT_FOUND.Code,
				Message: parser.ERRORS.PATH_NOT_FOUND.Message,
				Details: fmt.Sprintf("the resource %s was not found", r.RequestURI),
				Timestamp: time.Now().Local(),	
				Path: r.RequestURI,
			},
		})

	}) )

	return &router

}

func BuildRoutes(routers []Router, basepath string, api *http.ServeMux) {

	for _, router := range routers {
		for _, route := range router.Routes {
			api.HandleFunc(
				ParseRoutePatter(route.Method, basepath + router.Basepath, route.Path), 
				middlewares.Logger(BuildHandler(route.Handler, route.Middlewares)),
			)
		}
	}

}

func BuildHandler(handler http.HandlerFunc, middlewares []Middleware) http.HandlerFunc {
	
	newHandler := handler

	for i := len(middlewares); i > 0; i-- {
		newHandler = middlewares[i-1](newHandler)
	} 

	return newHandler
}

func ParseRoutePatter(method, basepath, path string) string {
	route := fmt.Sprintf(
		"%s %s%s",
		method,
		basepath,
		path,
	)

	return route
}

type Router struct {
	Basepath string
	Routes   []Route
}

type Route struct {
	Path    string
	Method  string
	Handler http.HandlerFunc
	Middlewares []Middleware
}

type Middleware func(next http.HandlerFunc) http.HandlerFunc
