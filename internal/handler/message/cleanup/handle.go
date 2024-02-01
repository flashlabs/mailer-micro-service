package cleanup

import (
	"log"
	"net/http"

	"github.com/flashlabs/mailer-micro-service/internal/repository/message"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	log.Println("HTTP cleanup message handler started")

	if err := message.DeleteOutdated(r.Context()); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusAccepted)

	log.Println("HTTP cleanup message handler ended")
}
