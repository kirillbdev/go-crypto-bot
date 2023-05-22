package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/kirillbdev/go-crypto-bot/internal/bot/di"
	"github.com/kirillbdev/go-crypto-bot/internal/bot/repository/dto"
	"os"
)

func main() {
	container := di.GetContainer()
	token := os.Getenv("BOT_TOKEN")

	if token == "" {
		panic("BOT_TOKEN variable does not set")
	}

	bot, err := tgbotapi.NewBotAPI(token)

	if err != nil {
		panic(err)
	}

	bot.Debug = true
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30
	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		var msg tgbotapi.MessageConfig

		if update.Message.IsCommand() && update.Message.Command() == "start" {
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Hello, I'm simple crypto monitoring bot")

			container.PortfolioRepo().Insert(dto.InsertPortfolio{
				UserId:    update.Message.Chat.ID,
				Name:      "My Portfolio",
				IsDefault: true,
			})
		}

		if _, err := bot.Send(msg); err != nil {
			panic(err)
		}
	}
}
