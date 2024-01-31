package initializer_test

import (
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"

	"github.com/flashlabs/mailer-micro-service/internal/initializer"
)

func TestHandler(t *testing.T) {
	require.NoError(t, initializer.Handler(mux.NewRouter()))
}
