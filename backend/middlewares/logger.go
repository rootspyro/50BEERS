package middlewares

import (
	"fmt"
	"net/http"
	"time"
)

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		start := time.Now()

		recorder := &ResponseRecorder{
			ResponseWriter: w,
			StatusCode: http.StatusOK,
		}

		next(recorder,r)

		duration := time.Since(start)

		fmt.Printf(
			"\n%s | %s | %d | %s | %s - %s",
			start.Local(),
			r.RemoteAddr,
			recorder.StatusCode,
			duration,
			r.Method,
			r.RequestURI,
		)
	}
}

func(r *ResponseRecorder) WriteHeader(code int) {
	r.StatusCode = code
	r.ResponseWriter.WriteHeader(code)
}

type ResponseRecorder struct {
	http.ResponseWriter
	StatusCode int
}
