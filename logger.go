package logit

import (
	"log"
	"net/http"
)

// Start logging HTTP Request and Response for every request
func StartLogger(address string, l *log.Logger, router *http.ServeMux) *http.Server {
	return &http.Server{
		Addr:     address,
		Handler:  middlewares{logging(l)}.apply(router),
		ErrorLog: l,
	}
}
