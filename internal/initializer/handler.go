package initializer

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/flashlabs/mailer-micro-service/internal/handler/message"
	messageCreate "github.com/flashlabs/mailer-micro-service/internal/handler/message/create"
	messageDelete "github.com/flashlabs/mailer-micro-service/internal/handler/message/delete"
	messageSend "github.com/flashlabs/mailer-micro-service/internal/handler/message/send"
)

func Handler(r *mux.Router) error {
	log.Println("Initializing handlers")

	r.HandleFunc(message.PatternCreateMessage, messageCreate.Handle).Methods(http.MethodPost)
	r.HandleFunc(message.PatternDeleteMessage, messageDelete.Handle).Methods(http.MethodDelete)
	r.HandleFunc(message.PatternSendMessage, messageSend.Handle).Methods(http.MethodPost)

	return nil
}
