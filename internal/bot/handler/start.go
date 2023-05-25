package handler

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/kirillbdev/go-crypto-bot/internal/bot/app"
	"github.com/kirillbdev/go-crypto-bot/internal/bot/repository"
	"github.com/kirillbdev/go-crypto-bot/internal/bot/repository/dto"
	"github.com/kirillbdev/go-crypto-bot/internal/bot/utils"
)

type StartHandler struct {
	portfolioRepository repository.PortfolioRepository
}

func NewStartHandler(portfolioRepository repository.PortfolioRepository) *StartHandler {
	return &StartHandler{portfolioRepository: portfolioRepository}
}

func (h *StartHandler) Handle(update tgbotapi.Update, ctx *app.Context) {
	portfolioId := h.portfolioRepository.FindDefaultPortfolioId(update.Message.Chat.ID)

	if portfolioId == 0 {
		h.portfolioRepository.Insert(dto.InsertPortfolio{
			ChatId:    update.Message.Chat.ID,
			Name:      "My Portfolio",
			IsDefault: true,
		})

		utils.SendMessage(
			ctx.Bot(),
			tgbotapi.NewMessage(update.Message.Chat.ID, "Hello, I'm simple crypto monitoring bot"),
		)
	} else {
		utils.SendMessage(
			ctx.Bot(),
			tgbotapi.NewMessage(update.Message.Chat.ID, "You are already registered. Start to monitor crypto."),
		)
	}
}

func (h *StartHandler) CanHandle(command string) bool {
	return command == "start"
}
