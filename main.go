package main

import (
	"github.com/thinkpaduser/av/internal/config"
	"log"
	"os"
)

func main() {
	conf := config.NewConfig("./conf.yaml")
	bot, err := tgbotapi.NewBotAPI(conf.Token)
	if err != nil {
		log.Panic(err)
	}
	log.Printf("Authorized on account %s", bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 300
	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		switch update.Message.Command() {
		case "help":
			msg.Text = "1. Set a category, type /cat; 2. Queue: type /q X1 Carbon"
		case "cat":
			msg.Text = "Categories available (for now): Laptops, PCs, Phones. Type /cat Laptops"
		case "q":
			msg.Text = "Quering ..."
		default:
			msg.Text = "Usage: '/help'"
		}
		bot.Send(msg)
	}
}
