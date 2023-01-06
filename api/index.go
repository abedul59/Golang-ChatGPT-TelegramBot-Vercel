package handler

import (
	"log"
	"os"

	
	"github.com/yanzay/tbot/v2"
)

// HandlerFunc is the signature type for the main function that will handle HTTP requests.
func HandlerFunc(w http.ResponseWriter, r *http.Request) {
bot := tbot.New(os.Getenv("TELEGRAM_BOT_TOKEN"),
			tbot.WithWebhook("https://golang-chat-gpt-telegram-bot-vercel.vercel.app/", ":8080"))
		c := bot.Client()
		bot.HandleMessage("ping", func(m *tbot.Message) {
			c.SendMessage(m.Chat.ID, "pong")
		})
		log.Fatal(bot.Start())
}



