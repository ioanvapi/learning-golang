package main

import (
	"strings"

	log "github.com/sirupsen/logrus"
	"gopkg.in/telegram-bot-api.v4"
)

func handleHelp(bot *tgbotapi.BotAPI, ud tgbotapi.Update) {
	msg := tgbotapi.NewMessage(ud.Message.Chat.ID, "欢迎使用量投社机器人!")
	msg.ReplyToMessageID = ud.Message.MessageID

	bot.Send(msg)
}

func main() {
	bot, err := tgbotapi.NewBotAPI(Opts.TgToken)
	if err != nil {
		log.Fatalf("Create Telegram API failed, err: %v", err)
	}

	bot.Debug = false

	log.Infof("Telegram robot name: %v", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 0

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		// ignore edited messages
		if update.Message == nil {
			continue
		}

		// tokenize the update
		tokens := strings.Split(update.Message.Text, " ")
		tokens = strings.Split(tokens[0], "@")
		command := strings.ToLower(tokens[0])

		switch command {
		case "help", "/help":
			go handleHelp(bot, update)
		default:
		}

		log.Infof("msg: %v", update.Message.Text)
	}
}
