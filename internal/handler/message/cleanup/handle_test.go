package cleanup_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/pashagolub/pgxmock/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/flashlabs/mailer-micro-service/internal/initializer"
	"github.com/flashlabs/mailer-micro-service/internal/registry"
)

func TestHandle(t *testing.T) {
	p, err := pgxmock.NewPool()
	if err != nil {
		t.Fatal(err)
	}
	defer p.Close()

	p.ExpectExec("DELETE FROM message").WillReturnResult(pgxmock.NewResult("DELETE", 1))

	registry.DBPool = p

	r, err := initializer.Router()
	require.NoError(t, err)

	err = initializer.Handler(r)
	require.NoError(t, err)

	req, err := http.NewRequestWithContext(context.Background(), http.MethodPost, "/api/messages/cleanup", nil)
	require.NoError(t, err)

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusAccepted, rr.Code)

	if err = p.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
