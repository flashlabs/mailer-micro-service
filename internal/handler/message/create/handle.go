package create

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/flashlabs/mailer-micro-service/internal/entity"
	"github.com/flashlabs/mailer-micro-service/internal/repository/message"
	"github.com/flashlabs/mailer-micro-service/pkg"
)

type data struct {
	InsertTime time.Time
	Email      string
	Title      string
	Content    string
	MailingID  int
}

func Handle(w http.ResponseWriter, r *http.Request) {
	log.Println("HTTP create resource handler started")

	rd, err := requestData(r)
	if err != nil {
		if errors.Is(errors.Unwrap(err), pkg.ErrInvalidPayload) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		} else {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		}

		return
	}

	if err = message.Create(r.Context(), entity.Message{
		InsertTime: rd.InsertTime,
		Email:      rd.Email,
		Title:      rd.Title,
		Content:    rd.Content,
		MailingID:  rd.MailingID,
	}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusCreated)

	log.Println("HTTP create resource handler ended")
}
