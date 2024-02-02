package create

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/flashlabs/mailer-micro-service/internal/entity"
	"github.com/flashlabs/mailer-micro-service/internal/repository/message"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	log.Println("HTTP create resource handler started")

	var m entity.Message
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if err := message.Create(r.Context(), m); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusCreated)
	}

	log.Println("HTTP create resource handler ended")
}
