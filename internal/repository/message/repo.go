package message

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"

	"github.com/flashlabs/mailer-micro-service/internal/entity"
	"github.com/flashlabs/mailer-micro-service/internal/entity/field"
	"github.com/flashlabs/mailer-micro-service/internal/registry"
)

func Create(ctx context.Context, e entity.Message) error {
	e.ID = field.New(&pgtype.UUID{}, uuid.NewString())

	if _, err := registry.DBPool.Exec(
		ctx,
		`INSERT INTO message(id, email, title, content, mailing_id, insert_time) VALUES ($1, $2, $3, $4, $5, $6)`,
		e.ID,
		e.Email,
		e.Title,
		e.Content,
		e.MailingID,
		e.InsertTime,
	); err != nil {
		return fmt.Errorf("error while executing registry.DBPool.Exec: %w", err)
	}

	return nil
}
