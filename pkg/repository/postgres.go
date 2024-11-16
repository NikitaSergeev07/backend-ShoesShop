package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"time"
)

const (
	usersTable = "users"
	maxRetries = 10
	retryDelay = 2 * time.Second
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	var db *sqlx.DB
	var err error

	for i := 0; i < maxRetries; i++ {
		db, err = sqlx.Connect("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DBName, cfg.SSLMode))
		if err == nil {
			err = db.Ping()
			if err == nil {
				log.Println("Успешное подключение к базе данных")
				return db, nil
			}
		}

		log.Printf("Ошибка подключения к базе данных: %v. Повторная попытка через %v...\n", err, retryDelay)
		time.Sleep(retryDelay)
	}

	return nil, fmt.Errorf("не удалось подключиться к базе данных после %d попыток: %v", maxRetries, err)
}
