package main

import (
	"log"
	"net/http"

	"github.com/abhinavdahiya/go-messenger-bot"
)

func main() {
	bot := mbotapi.NewBotAPI("ACCESS_TOKEN", "VERIFY_TOKEN")

	callbacks, mux := bot.SetWebhook("/webhook")
	go http.ListenAndServeTLS("0.0.0.0:8443", "cert.pem", "key.pem", mux)

	for callback := range callbacks {
		log.Printf("[%#v] %s", callback.Sender, callback.Message.Text)

		msg := mbotapi.NewMessage(callback.Message.Text)
		bot.Send(callback.Sender, msg, mbotapi.RegularNotif)
	}
}
