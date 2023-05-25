package app

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type Context struct {
	bot *tgbotapi.BotAPI
}

func NewContext(bot *tgbotapi.BotAPI) *Context {
	return &Context{bot: bot}
}

func (ctx *Context) Bot() *tgbotapi.BotAPI {
	return ctx.bot
}
