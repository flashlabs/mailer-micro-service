package create_test

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/pashagolub/pgxmock/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/flashlabs/mailer-micro-service/internal/initializer"
	"github.com/flashlabs/mailer-micro-service/internal/registry"
)

var (
	payload = []byte(`{
		"email": "jan.kowalski-12345@example.com",
		"title": "Interview",
		"content": "simple text",
		"mailing_id": 1,
		"insert_time": "2020-04-24T05:42:38.725412916Z"
	}`)
)

func TestCreateHandleCompletesSuccessfully(t *testing.T) {
	p, err := pgxmock.NewPool()
	if err != nil {
		t.Fatal(err)
	}
	defer p.Close()

	tm, err := time.Parse(time.RFC3339Nano, "2020-04-24T05:42:38.725412916Z")
	if err != nil {
		t.Fatal(err)
	}

	p.ExpectExec("INSERT INTO message").
		WithArgs(
			pgxmock.AnyArg(),
			"jan.kowalski-12345@example.com",
			"Interview",
			"simple text",
			1,
			tm,
		).
		WillReturnResult(pgxmock.NewResult("INSERT", 1))

	registry.DBPool = p

	r, err := initializer.Router()
	require.NoError(t, err)

	err = initializer.Handler(r)
	require.NoError(t, err)

	req, err := http.NewRequestWithContext(context.Background(), http.MethodPost, "/api/messages", bytes.NewBuffer(payload))
	require.NoError(t, err)

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)

	if err = p.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
