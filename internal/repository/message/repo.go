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

const (
	table = "message"
)

func ByID(ctx context.Context, id string) (entity.Message, error) {
	query := fmt.Sprintf(`SELECT m.id, m.email, m.title, m.content, m.mailing_id, m.insert_time FROM %s m WHERE m.id = $1`, table)

	var m entity.Message

	if err := registry.DBPool.QueryRow(ctx, query, id).Scan(&m.ID, &m.Email, &m.Title, &m.Content, &m.MailingID, &m.InsertTime); err != nil {
		return entity.Message{}, fmt.Errorf("error while executing SELECT registry.DBPool.Exec: %w", err)
	}

	return entity.Message{}, nil
}

func DeleteByID(ctx context.Context, id string) error {
	query := fmt.Sprintf(`DELETE FROM %s m WHERE m.id = $1`, table)

	if _, err := registry.DBPool.Exec(ctx, query, id); err != nil {
		return fmt.Errorf("error while executing DELETE registry.DBPool.Exec: %w", err)
	}

	return nil
}

func DeleteByMailingID(ctx context.Context, id uint) error {
	query := fmt.Sprintf(`DELETE FROM %s m WHERE m.mailing_id = $1`, table)

	if _, err := registry.DBPool.Exec(ctx, query, id); err != nil {
		return fmt.Errorf("error while executing registry.DBPool.Exec: %w", err)
	}

	return nil
}

func FindByMailingID(ctx context.Context, id uint) ([]entity.Message, error) {
	var result []entity.Message

	query := fmt.Sprintf(`SELECT m.id, m.email, m.title, m.content, m.mailing_id, m.insert_time FROM %s m WHERE m.mailing_id = $1`, table)

	rows, err := registry.DBPool.Query(
		ctx,
		query,
		id,
	)
	if err != nil {
		return nil, fmt.Errorf("error while executing registry.DBPool.Query: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		var m entity.Message
		if e := rows.Scan(&m.ID, &m.Email, &m.Title, &m.Content, &m.MailingID, &m.InsertTime); e != nil {
			return nil, fmt.Errorf("error while executing rows.Scan: %w", e)
		}

		result = append(result, m)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("rows contained errors while reading: %w", err)
	}

	return result, nil
}

func Create(ctx context.Context, e entity.Message) error {
	e.ID = field.New(&pgtype.UUID{}, uuid.NewString())

	query := fmt.Sprintf(`INSERT INTO %s (id, email, title, content, mailing_id, insert_time) VALUES ($1, $2, $3, $4, $5, $6)`, table)

	if _, err := registry.DBPool.Exec(
		ctx,
		query,
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
