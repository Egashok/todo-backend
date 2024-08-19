package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMODE  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {

	db, error := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMODE))
	if error != nil {
		return nil, error
	}
	error = db.Ping()
	if error != nil {
		return nil, error
	}
	return db, nil
}