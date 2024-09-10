package middlewares

import (
	"context"
	"net/http"
)

func LangHeader(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// get language header
		header := r.Header

		language := header.Get("Accept-Language")
		
		// if header is empty or value is not valid set default language as "en" (ENGLISH)
		if language == "" || language != "es" && language != "en" {
			language = "en"
		}

		ctx := context.WithValue(r.Context(), "lang", language)
		next.ServeHTTP(w,r.WithContext(ctx))
	}
}
