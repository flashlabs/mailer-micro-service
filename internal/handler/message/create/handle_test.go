package create_test

import (
	"bytes"
	"context"
	"errors"
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

	ErrDbError = errors.New("database error")
)

func TestCreateHandleCompletesWithError(t *testing.T) {
	p, err := pgxmock.NewPool()
	if err != nil {
		t.Fatal(err)
	}
	defer p.Close()

	p.ExpectExec("INSERT INTO message").
		WithArgs(pgxmock.AnyArg(), pgxmock.AnyArg(), pgxmock.AnyArg(), pgxmock.AnyArg(), pgxmock.AnyArg(), pgxmock.AnyArg()).
		WillReturnError(ErrDbError)

	registry.DBPool = p

	r, err := initializer.Router()
	require.NoError(t, err)

	err = initializer.Handler(r)
	require.NoError(t, err)

	req, err := http.NewRequestWithContext(context.Background(), http.MethodPost, "/api/messages", bytes.NewBuffer(payload))
	require.NoError(t, err)

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)

	if err = p.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

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

func TestCreateHandlePayloads(t *testing.T) {
	type args struct {
		message string
		payload []byte
		status  int
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "missing email field",
			args: args{
				payload: []byte(`{
					"invalid_email_field": "jan.kowalski-12345@example.com",
					"title": "Interview",
					"content": "simple text",
					"mailing_id": 1,
					"insert_time": "2020-04-24T05:42:38.725412916Z"
				}`),
				status:  http.StatusUnprocessableEntity,
				message: "email error: missing param\n",
			},
		},
		{
			name: "invalid mailing_id type",
			args: args{
				payload: []byte(`{
					"email": "jan.kowalski-12345@example.com",
					"title": "Interview",
					"content": "simple text",
					"mailing_id": "1",
					"insert_time": "2020-04-24T05:42:38.725412916Z"
				}`),
				status:  http.StatusUnprocessableEntity,
				message: "mailing_id error: invalid data type\n",
			},
		},
		{
			name: "malformed payload",
			args: args{
				payload: []byte(`{
					"email": "jan.kowalski-12345@example.com",
					"title": "Interview",
					"content": "simple text",
					"mailing_id": "1""",
					"insert_time": "2020-04-24T05:42:38.725412916Z"
				}`),
				status:  http.StatusBadRequest,
				message: "error while executing handler.Payload: invalid payload\n",
			},
		},
	}

	r, err := initializer.Router()
	require.NoError(t, err)

	err = initializer.Handler(r)
	require.NoError(t, err)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequestWithContext(context.Background(), http.MethodPost, "/api/messages", bytes.NewBuffer(tt.args.payload))
			require.NoError(t, err)

			rr := httptest.NewRecorder()
			r.ServeHTTP(rr, req)

			assert.Equal(t, tt.args.status, rr.Code)
			assert.Equal(t, tt.args.message, rr.Body.String())
		})
	}
}
