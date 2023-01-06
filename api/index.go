package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	gogpt "github.com/sashabaranov/go-gpt3"
	"github.com/yanzay/tbot/v2"
)

// HandlerFunc is the signature type for the main function that will handle HTTP requests.
func HandlerFunc(w http.ResponseWriter, r *http.Request) {
	// Create a new instance of the GPT-3 client.
	c := gogpt.NewClient(os.Getenv("OPENAI_TOKEN"))
	ctx := context.Background()

	bot := tbot.New(os.Getenv("TELEGRAM_BOT_TOKEN"),
		tbot.WithWebhook("https://golang-chat-gpt-telegram-bot-vercel.vercel.app/", ":8080"))
	c1 := bot.Client() //Please add your Render URL between "". 請在引號中加入你的Render網址

	/////////////////

	bot.HandleMessage("ping", func(m *tbot.Message) {
			c1.SendMessage(m.Chat.ID, "pong")
		})
	log.Fatal(bot.Start())
}

func Main() {
http.HandleFunc("/", HandlerFunc)
http.ListenAndServe(":8080", nil)
}

