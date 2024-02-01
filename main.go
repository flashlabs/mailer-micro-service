package main

import (
	"context"
	"fmt"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"

	"github.com/flashlabs/mailer-micro-service/internal/initializer"
	"github.com/flashlabs/mailer-micro-service/internal/initializer/server"
)

const (
	appWelcomeMessage           = "Mailer Micro Service app started" // #nosec G101
	appDoneMessage              = "App done"
	appInitializationFailedCode = 1
)

func main() {
	log.Println(appWelcomeMessage)

	c := context.Background()

	if err := initApp(c); err != nil {
		log.Printf("error while executing initApp: %q", err)

		os.Exit(appInitializationFailedCode)
	}

	log.Println(appDoneMessage)
}

func initApp(c context.Context) error {
	if err := initializer.Database(c); err != nil {
		return fmt.Errorf("error while executing initializer.Database: %w", err)
	}

	initializer.Mailer()

	r, err := initializer.Router()
	if err != nil {
		return fmt.Errorf("error while executing initializer.Router: %w", err)
	}

	if err = initializer.Handler(r); err != nil {
		return fmt.Errorf("error while executing initializer.Handler: %w", err)
	}

	srv, err := initializer.Server(r, server.Port)
	if err != nil {
		return fmt.Errorf("error while executing initializer.Server: %w", err)
	}

	if err = srv.ListenAndServe(); err != nil {
		return fmt.Errorf("error while executing srv.ListenAndServe: %w", err)
	}

	return nil
}
