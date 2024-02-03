package send

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/flashlabs/mailer-micro-service/internal/registry"
	"github.com/flashlabs/mailer-micro-service/internal/repository/message"
)

type payload struct {
	MailingID uint `json:"mailing_id"`
}

func Handle(w http.ResponseWriter, r *http.Request) {
	log.Println("HTTP send message handler started")

	c := r.Context()

	var p payload
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	log.Printf("sending mailing #%d requested", p.MailingID)

	messages, err := message.FindByMailingID(c, p.MailingID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	if err = registry.Mailer.SendBatch(messages); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	if err = message.DeleteByMailingID(c, p.MailingID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusAccepted)

	log.Println("HTTP send message handler ended")
}
