package entity

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type Message struct {
	InsertTime time.Time    `json:"insert_time"`
	ID         *pgtype.UUID `json:"-"`
	Email      string       `json:"email"`
	Title      string       `json:"title"`
	Content    string       `json:"content"`
	MailingID  int          `json:"mailing_id"`
}
