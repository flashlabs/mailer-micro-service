package registry

import (
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/flashlabs/mailer-micro-service/internal/service"
)

var DBPool *pgxpool.Pool

var Mailer *service.Mailer
