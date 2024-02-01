package service

import (
	"fmt"
	"log"

	"golang.org/x/sync/errgroup"

	"github.com/flashlabs/mailer-micro-service/internal/entity"
)

type Mailer struct {
}

func (m *Mailer) SendBatch(messages []entity.Message) error {
	if len(messages) == 0 {
		return nil
	}

	g := new(errgroup.Group)

	for _, message := range messages {
		msg := message

		g.Go(func() error {
			if err := m.sendMessage(msg); err != nil {
				return fmt.Errorf("error while sending a message: %w", err)
			}

			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return fmt.Errorf("error while sending a messages: %w", err)
	}

	return nil
}

func (m *Mailer) Send(message entity.Message) error {
	if err := m.sendMessage(message); err != nil {
		return fmt.Errorf("error while sending a message: %w", err)
	}

	return nil
}

func (m *Mailer) sendMessage(message entity.Message) error {
	// here goes send email implementation
	log.Printf("sending mailing #%d to the %q customer", message.MailingID, message.Email)

	return nil
}
