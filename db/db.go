package db

import (
	"context"
	"fmt"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	Name   string
}

func Connect(ctx context.Context, c Config) (*sqlx.DB, error) {
	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		c.Host, c.Port, c.User, c.Password, c.Name)

	db, err := sqlx.ConnectContext(ctx, "pgx", connString)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	return db, nil
}
