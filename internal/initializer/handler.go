package initializer

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	messageCreate "github.com/flashlabs/mailer-micro-service/internal/handler/message/create"
	messageDelete "github.com/flashlabs/mailer-micro-service/internal/handler/message/delete"
	messageSend "github.com/flashlabs/mailer-micro-service/internal/handler/message/send"
)

const (
	createMessage = "/api/messages"
	deleteMessage = "/api/messages/{id:[0-9]+}"
	sendMessage   = "/api/messages/send"
)

func Handler(r *mux.Router) error {
	log.Println("Initializing handlers")

	r.HandleFunc(createMessage, messageCreate.Handle).Methods(http.MethodPost)
	r.HandleFunc(deleteMessage, messageDelete.Handle).Methods(http.MethodDelete)
	r.HandleFunc(sendMessage, messageSend.Handle).Methods(http.MethodPost)

	return nil
}
