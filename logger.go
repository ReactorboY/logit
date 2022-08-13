package logit

import (
	"log"
	"net/http"
)

// Start logging HTTP Request and Response for every request
// @required -> Logger, router, and address to listen
func StartLogger(address string, l *log.Logger, router *http.ServeMux) *http.Server {
	return &http.Server{
		Addr:     address,
		Handler:  middlewares{logging(l)}.apply(router),
		ErrorLog: l,
	}
}
