package initializer_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/flashlabs/mailer-micro-service/internal/initializer"
	"github.com/flashlabs/mailer-micro-service/internal/initializer/server"
)

func TestServer(t *testing.T) {
	r := mux.NewRouter()
	srv, err := initializer.Server(r, server.Port)

	require.NoError(t, err)
	assert.Equal(t, &http.Server{
		Addr:              fmt.Sprintf(":%d", server.Port),
		Handler:           r,
		ReadTimeout:       server.ReadTimeout,
		WriteTimeout:      server.WriteTimeout,
		IdleTimeout:       server.IdleTimeout,
		ReadHeaderTimeout: server.ReadHeaderTimeout,
	}, srv)
}
