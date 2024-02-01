package delete

import (
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5"

	"github.com/flashlabs/mailer-micro-service/internal/repository/message"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	log.Println("HTTP delete resource handler started")

	c := r.Context()

	vars := mux.Vars(r)
	id := vars["id"]
	log.Printf("delete message %q requested", id)

	_, err := message.ByID(c, id)

	switch {
	case err == nil:
		if err = message.DeleteByID(c, id); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.WriteHeader(http.StatusNoContent)
	case errors.Is(errors.Unwrap(err), pgx.ErrNoRows):
		http.NotFound(w, r)
	default:
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}

	log.Println("HTTP delete resource handler ended")
}
