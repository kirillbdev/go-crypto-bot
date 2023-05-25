package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kirillbdev/go-crypto-bot/internal/bot/repository/dto"
	"log"
	"time"
)

type MySqlPortfolioRepository struct {
	db *sql.DB
}

func NewMySqlPortfolioRepository(host string, user string, pass string, name string) *MySqlPortfolioRepository {
	db, err := sql.Open(
		"mysql",
		fmt.Sprintf("%s:%s@(%s)/%s", user, pass, host, name),
	)

	if err != nil {
		panic("Unable to open database connection" + err.Error())
	}

	return &MySqlPortfolioRepository{
		db: db,
	}
}

func (repo *MySqlPortfolioRepository) Insert(dto dto.InsertPortfolio) int64 {
	now := float64(time.Now().UnixMicro()) / 1000000

	query := "INSERT INTO `portfolio` (`chat_id`, `name`, `default`, `created_at`, `updated_at`) VALUES (?, ?, ?, ?, ?)"
	result, err := repo.db.Exec(query, dto.ChatId, dto.Name, dto.IsDefault, now, now)

	if err != nil {
		log.Fatalf("Database insert error: %s", err)
	}

	id, err := result.LastInsertId()

	if err != nil {
		log.Fatalf("Database insert error: %s", err)
	}

	return id
}

func (repo *MySqlPortfolioRepository) FindDefaultPortfolioId(chatId int64) int64 {
	row := repo.db.QueryRow("SELECT `id` FROM `portfolio` WHERE `chat_id` = ? AND `default` = 1", chatId)

	var id int64

	if err := row.Scan(&id); err != nil {
		return 0
	}

	return id
}
