package di

import (
	"github.com/kirillbdev/go-crypto-bot/internal/bot/repository"
)

type Container struct {
	portfolioRepo repository.PortfolioRepository
}

func (c *Container) PortfolioRepo() repository.PortfolioRepository {
	return c.portfolioRepo
}

var containerInstance *Container

func CreateContainer(config Config) {
	if containerInstance != nil {
		panic("Container already created")
	}

	portfolioRepo := repository.NewMySqlPortfolioRepository(
		config.dbHost,
		config.dbUser,
		config.dbPass,
		config.dbName,
	)

	containerInstance = &Container{
		portfolioRepo: portfolioRepo,
	}
}

func GetContainer() *Container {
	if containerInstance == nil {
		panic("You cannot retrieve container without initialization")
	}

	return containerInstance
}
