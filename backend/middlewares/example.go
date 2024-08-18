package middlewares

import (
	"net/http"

	"github.com/rootspyro/50BEERS/config/log"
)

func Example1(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Info("This is an example middleware")
		next(w,r)
	}
}

func Example2(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Info("This is another example middleware")
		next(w,r)
	}
}
