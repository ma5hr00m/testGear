package main

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var bot *tgbotapi.BotAPI

const (
	CHAITID = 5920766613
	TOKEN   = "5866674937:AAGuitaCM28NCPmlb_7f9A7nq6H9Vhv8tRM"
)

func main() {
	var err error
	bot, err = tgbotapi.NewBotAPI(TOKEN)

	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false

	message := fmt.Sprintf("Hello,Ma5hr00m!")

	sendMsg(message)
}

func sendMsg(msg string) {
	newMsg := tgbotapi.NewMessage(CHAITID, msg)

	_, err := bot.Send(newMsg)

	if err == nil {
		log.Printf("Send telegram message success")
	} else {
		log.Printf("Send telegram message error")
	}
}
