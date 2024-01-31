package initializer

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/flashlabs/mailer-micro-service/internal/registry"
)

const (
	databaseEnvKey = "DATABASE_URL"
)

func Database(ctx context.Context) error {
	pool, err := pgxpool.New(ctx, os.Getenv(databaseEnvKey))
	if err != nil {
		return fmt.Errorf("unable to create connection pool: %w", err)
	}

	registry.DBPool = pool

	return nil
}
