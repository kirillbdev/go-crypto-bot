package utils

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func SendMessage(bot *tgbotapi.BotAPI, msg tgbotapi.Chattable) bool {
	_, err := bot.Send(msg)

	return err == nil
}
