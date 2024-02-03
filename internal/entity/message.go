package entity

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type Message struct {
	InsertTime time.Time
	ID         *pgtype.UUID
	Email      string
	Title      string
	Content    string
	MailingID  int
}
