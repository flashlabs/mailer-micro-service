package initializer_test

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/flashlabs/mailer-micro-service/internal/initializer"
	"github.com/flashlabs/mailer-micro-service/internal/registry"
)

func TestDatabase(t *testing.T) {
	require.NoError(t, initializer.Database(context.Background()))

	assert.NotNil(t, registry.DBPool)
	assert.IsType(t, &pgxpool.Pool{}, registry.DBPool)
}
