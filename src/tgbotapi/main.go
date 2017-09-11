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
func handleChat(bot *tgbotapi.BotAPI, ud tgbotapi.Update) {
	tokens := strings.SplitN(ud.Message.Text, " ", 2)
	var text string
	if len(tokens) > 1 {
		text = tokens[1]
	} else {
		text = tokens[0]
	}
	resp, err := SendTuling(text, string(ud.Message.Chat.ID))
	if err != nil {
		log.Error(err)
	}

	msg := tgbotapi.NewMessage(ud.Message.Chat.ID, resp.Text)
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
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		// ignore edited messages
		if update.Message == nil {
			continue
		}

		if len(update.Message.Text) == 0 {
			continue
		}

		// tokenize the update
		tokens := strings.Split(update.Message.Text, " ")
		commands := strings.Split(tokens[0], "@")
		command := strings.ToLower(commands[0])

		switch command {
		case "help", "/help":
			go handleHelp(bot, update)
		case "chat", "/chat", "/c":
			go handleChat(bot, update)
		default:
			go handleChat(bot, update)
		}

		log.Infof("msg: %v", update.Message.Text)
	}
}
