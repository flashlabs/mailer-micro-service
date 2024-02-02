package initializer

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/flashlabs/mailer-micro-service/internal/registry"
)

func Database(ctx context.Context) error {
	conn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
	)

	pool, err := pgxpool.New(ctx, conn)
	if err != nil {
		return fmt.Errorf("unable to create connection pool: %w", err)
	}

	registry.DBPool = pool

	return nil
}
