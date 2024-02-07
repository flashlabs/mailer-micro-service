package send

import (
	"errors"
	"log"
	"net/http"

	"github.com/flashlabs/mailer-micro-service/internal/registry"
	"github.com/flashlabs/mailer-micro-service/internal/repository/message"
	"github.com/flashlabs/mailer-micro-service/pkg"
)

type data struct {
	MailingID int
}

func Handle(w http.ResponseWriter, r *http.Request) {
	log.Println("HTTP send message handler started")

	rd, err := requestData(r)
	if err != nil {
		if errors.Is(errors.Unwrap(err), pkg.ErrInvalidPayload) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		} else {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		}

		return
	}

	log.Printf("Sending mailing #%d requested", rd.MailingID)

	c := r.Context()

	messages, err := message.FindByMailingID(c, rd.MailingID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	if err = registry.Mailer.SendBatch(messages); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	if err = message.DeleteByMailingID(c, rd.MailingID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusAccepted)

	log.Println("HTTP send message handler ended")
}
