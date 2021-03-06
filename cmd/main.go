package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	godotenv.Load()
	token := os.Getenv("TOKEN")
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.UpdateConfig{
		Timeout: 60,
	}

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}

	for update := range updates {
		if update.Message == nil { // If we got a message
			continue
		}
		switch update.Message.Command() {
		case "help":
			HelpCommand(bot, update.Message)
		default:
			DefaultBehavior(bot, update.Message)
		}

	}
}

func HelpCommand(bot *tgbotapi.BotAPI, InputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(InputMessage.Chat.ID, "/help - help")
	bot.Send(msg)
}

func DefaultBehavior(bot *tgbotapi.BotAPI, InputMessage *tgbotapi.Message) {
	log.Printf("[%s] %s", InputMessage.From.UserName, InputMessage.Text)

	msg := tgbotapi.NewMessage(InputMessage.Chat.ID, "Моя бейби написала: "+InputMessage.Text)
	//msg.ReplyToMessageID = update.Message.MessageID

	bot.Send(msg)
}
