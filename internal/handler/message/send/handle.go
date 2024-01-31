package create

import (
	"log"
	"net/http"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	log.Println("HTTP send message handler started")

	log.Println("HTTP send message handler ended")
}
