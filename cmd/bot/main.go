package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	internal "github.com/kirillbdev/go-crypto-bot/internal/bot"
	"github.com/kirillbdev/go-crypto-bot/internal/bot/app"
	"github.com/kirillbdev/go-crypto-bot/internal/bot/di"
	"github.com/kirillbdev/go-crypto-bot/internal/bot/handler"
	"os"
)

func main() {
	di.CreateContainer(di.NewConfig(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
	))

	container := di.GetContainer()
	token := os.Getenv("BOT_TOKEN")

	if token == "" {
		panic("BOT_TOKEN variable does not set")
	}

	bot, err := tgbotapi.NewBotAPI(token)

	if err != nil {
		panic(err)
	}

	ctx := app.NewContext(bot)
	processor := internal.NewUpdatesProcessor(ctx)
	processor.AddHandler(handler.NewStartHandler(container.PortfolioRepo()))

	bot.Debug = true
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30
	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		processor.Process(update)
	}
}
