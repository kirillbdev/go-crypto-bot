package repository

import "github.com/kirillbdev/go-crypto-bot/internal/bot/repository/dto"

type PortfolioRepository interface {
	Insert(portfolio dto.InsertPortfolio) int64
}
