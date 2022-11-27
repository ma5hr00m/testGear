package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	TOKEN = "5866674937:AAGuitaCM28NCPmlb_7f9A7nq6H9Vhv8tRM"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(TOKEN)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 30

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		if update.Message == nil {
			continue
		}

		if !update.Message.IsCommand() {
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		switch update.Message.Command() {
		case "help":
			msg.Text = "You can use the following command:\n/hello\n/ping"
			log.Printf("Successfully execute /help command.")
		case "hello":
			msg.Text = ("Hello, " + update.Message.From.UserName + "!")
			log.Printf("Successfully execute /hello command.")
		case "ping":
			msg.Text = "pong!"
			log.Printf("Successfully execute /ping command.")
		default:
			continue
		}

		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}
