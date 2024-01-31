package create

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	log.Println("HTTP delete resource handler started")

	vars := mux.Vars(r)
	log.Printf("delete message %q requested", vars["id"])

	log.Println("HTTP delete resource handler ended")
}
