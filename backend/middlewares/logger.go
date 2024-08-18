package middlewares

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/rootspyro/50BEERS/config"
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
			" %s | %s | %s | %s | %s - %s\n",
			start.Local(),
			r.RemoteAddr,
			parseCode(recorder.StatusCode),
			duration,
			parseMethod(r.Method),
			r.RequestURI,
		)
	}
}

func parseMethod(method string) string {
	var methodcolor string = config.Colors.Reset

	switch method {
	case "GET":
		methodcolor = config.Colors.Green
		break;

	case "POST":
		methodcolor = config.Colors.Blue
		break;

	case "PUT":
		methodcolor = config.Colors.Purple
		break

	case "PATCH":
		methodcolor = config.Colors.Yellow
		break

	case "DELETE":
		methodcolor = config.Colors.Red
		break

	default:
		methodcolor = config.Colors.Cyan
		break
	}


	return fmt.Sprintf("%s%s%s", methodcolor, method, config.Colors.Reset)
}

func parseCode(code int) string {
	var codeColor string = config.Colors.Reset

	stringCode := fmt.Sprintf("%d", code)
	firstNum := strings.Split(stringCode, "")[0]

	switch firstNum {
		case "2": // HTTP - 2xx 
			codeColor = config.Colors.Green
			break

		case "3": // HTTP - 3xx
			codeColor = config.Colors.Blue
			break

		case "4": // HTTP - 4xx
			codeColor = config.Colors.Yellow
			break

		case "5": // HTTP - 5xx 
			codeColor = config.Colors.Red
			break

		default:
			codeColor = config.Colors.Reset
			break
	}

	return fmt.Sprintf("%s%d%s", codeColor, code, config.Colors.Reset)
}

func(r *ResponseRecorder) WriteHeader(code int) {
	r.StatusCode = code
	r.ResponseWriter.WriteHeader(code)
}

type ResponseRecorder struct {
	http.ResponseWriter
	StatusCode int
}
