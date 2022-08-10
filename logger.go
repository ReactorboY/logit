package logit

import (
	"log"
	"net/http"
)

func StartLogger(address string, l *log.Logger, router *http.ServeMux) *http.Server {
	return &http.Server{
		Addr:     address,
		Handler:  middlewares{logging(l)}.apply(router),
		ErrorLog: l,
	}
}
