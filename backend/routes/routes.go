package routes

import (
	"fmt"
	"net/http"
)

func Init() *http.ServeMux {

	var router http.ServeMux = http.ServeMux{}

	var childRouters []Router = []Router{
		HealthRouter,
	}

	apiV1 := BuildRoutes(childRouters, "/api/v1")
	router.Handle("/", apiV1)

	return &router

}

func BuildRoutes(routers []Router, basepath string) *http.ServeMux {
	api := http.ServeMux{}

	for _, router := range routers {
		for _, route := range router.Routes {
			api.HandleFunc(
				ParseRoutePatter(route.Method, basepath + router.Basepath, route.Path), 
				route.Handler,
			)
		}
	}

	return &api
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
}
