package postgres


import (
	"database/sql"
	"fmt"
	"sustainability-service/config"

	_ "github.com/lib/pq"
)

func ConnectDb() (*sql.DB, error) {
	cfg := config.Load()

	conn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		cfg.DB_HOST, cfg.DB_PORT, cfg.DB_USER, cfg.DB_NAME, cfg.DB_PASSWORD)
	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
