package logit

import (
	"log"
	"net/http"
)

type wrappedResponse struct {
	http.ResponseWriter
	status int
}

func (w *wrappedResponse) WriteHeader(code int) {
	w.status = code
	w.ResponseWriter.WriteHeader(code)
}

func newWriter(w http.ResponseWriter) *wrappedResponse {
	return &wrappedResponse{ResponseWriter: w}
}

// create alias of middleware
type middleware func(http.Handler) http.Handler


// list of middlewares
type middlewares []middleware

// router and middlewares
func (list middlewares) apply(handler http.Handler) http.Handler {
	if len(list) == 0 {
		return handler
	}
	return list[1:].apply(list[0](handler))
}

// logging function
func logging(logger *log.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			writer := newWriter(w)
			defer func() {
				logger.Println(r.Method, r.URL.Path, writer.status)
			}()
			next.ServeHTTP(writer, r)
		})
	}
}