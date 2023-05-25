package handler

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/kirillbdev/go-crypto-bot/internal/bot/app"
)

type CommandHandler interface {
	Handle(update tgbotapi.Update, ctx *app.Context)
	CanHandle(command string) bool
}
