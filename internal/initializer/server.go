package initializer

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/flashlabs/mailer-micro-service/internal/initializer/server"
)

func Server(r *mux.Router, port int) (*http.Server, error) {
	return &http.Server{
		Addr:              fmt.Sprintf(":%d", port),
		Handler:           r,
		ReadTimeout:       server.ReadTimeout,
		WriteTimeout:      server.WriteTimeout,
		IdleTimeout:       server.IdleTimeout,
		ReadHeaderTimeout: server.ReadHeaderTimeout,
	}, nil
}
