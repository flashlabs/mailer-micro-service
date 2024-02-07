package initializer

import (
	"context"
	"fmt"
	"log"
	"time"

	"golang.org/x/sync/errgroup"

	"github.com/flashlabs/mailer-micro-service/internal/repository/message"
)

const (
	interval = time.Minute * 1
)

func Cleanup(ctx context.Context) {
	g, c := errgroup.WithContext(ctx)

	g.Go(func() error {
		if err := deleteOutdated(c); err != nil {
			return fmt.Errorf("error while executing deleteOutdated: %w", err)
		}

		return nil
	})

	go func() {
		if err := g.Wait(); err != nil {
			log.Println("error while executing g.Wait: ", err)
		}
	}()
}

func deleteOutdated(ctx context.Context) error {
	t := time.NewTicker(interval)
	defer t.Stop()

	for {
		select {
		case <-ctx.Done():
			return fmt.Errorf("context is done: %w", ctx.Err())
		case <-t.C:
			log.Println("delete outdated cycle")

			if err := message.DeleteOutdated(ctx); err != nil {
				return fmt.Errorf("error while executing message.DeleteOutdated: %w", err)
			}
		}
	}
}
