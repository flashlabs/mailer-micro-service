package initializer

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	messageCleanup "github.com/flashlabs/mailer-micro-service/internal/handler/message/cleanup"
	messageCreate "github.com/flashlabs/mailer-micro-service/internal/handler/message/create"
	messageDelete "github.com/flashlabs/mailer-micro-service/internal/handler/message/delete"
	messageSend "github.com/flashlabs/mailer-micro-service/internal/handler/message/send"
)

const (
	createMessage  = "/api/messages"
	deleteMessage  = "/api/messages/{id:[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}}"
	sendMessage    = "/api/messages/send"
	cleanupMessage = "/api/messages/cleanup"
)

func Handler(r *mux.Router) error {
	log.Println("Initializing handlers")

	r.HandleFunc(createMessage, messageCreate.Handle).Methods(http.MethodPost)
	r.HandleFunc(deleteMessage, messageDelete.Handle).Methods(http.MethodDelete)
	r.HandleFunc(sendMessage, messageSend.Handle).Methods(http.MethodPost)
	r.HandleFunc(cleanupMessage, messageCleanup.Handle).Methods(http.MethodPost)

	return nil
}
