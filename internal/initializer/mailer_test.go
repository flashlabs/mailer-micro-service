package initializer_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/flashlabs/mailer-micro-service/internal/initializer"
	"github.com/flashlabs/mailer-micro-service/internal/registry"
	"github.com/flashlabs/mailer-micro-service/internal/service"
)

func TestMailer(t *testing.T) {
	initializer.Mailer()

	assert.NotNil(t, registry.Mailer)
	assert.IsType(t, &service.Mailer{}, registry.Mailer)
}
