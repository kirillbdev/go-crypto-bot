package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/kirillbdev/go-crypto-bot/internal/bot/app"
	"github.com/kirillbdev/go-crypto-bot/internal/bot/handler"
)

type UpdatesProcessor struct {
	handlers []handler.CommandHandler
	ctx      *app.Context
}

func NewUpdatesProcessor(ctx *app.Context) *UpdatesProcessor {
	return &UpdatesProcessor{ctx: ctx}
}

func (proc *UpdatesProcessor) AddHandler(h handler.CommandHandler) {
	proc.handlers = append(proc.handlers, h)
}

func (proc *UpdatesProcessor) Process(update tgbotapi.Update) {
	if !update.Message.IsCommand() {
		return
	}

	for _, h := range proc.handlers {
		if h.CanHandle(update.Message.Command()) {
			h.Handle(update, proc.ctx)
		}
	}
}
