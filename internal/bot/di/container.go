package di

import (
	"github.com/kirillbdev/go-crypto-bot/internal/bot/repository"
	"os"
)

type Container struct {
	portfolioRepo repository.PortfolioRepository
}

func (c *Container) PortfolioRepo() repository.PortfolioRepository {
	return c.portfolioRepo
}

var containerInstance *Container

func createContainer() *Container {
	portfolioRepo := repository.NewMySqlPortfolioRepository(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
	)

	container := Container{
		portfolioRepo: portfolioRepo,
	}

	return &container
}

func GetContainer() *Container {
	if containerInstance == nil {
		containerInstance = createContainer()
	}

	return containerInstance
}
