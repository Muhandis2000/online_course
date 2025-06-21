package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/Muhandis2000/online-school/internal/config"
	_ "github.com/lib/pq"
)

func NewDB(cfg config.Config) (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Database.Host, cfg.Database.Port, cfg.Database.User, os.Getenv("DB_PASSWORD"), cfg.Database.DBName)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return db, db.Ping()
}
