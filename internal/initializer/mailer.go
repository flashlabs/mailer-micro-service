package initializer

import (
	"github.com/flashlabs/mailer-micro-service/internal/registry"
	"github.com/flashlabs/mailer-micro-service/internal/service"
)

func Mailer() {
	registry.Mailer = &service.Mailer{}
}
