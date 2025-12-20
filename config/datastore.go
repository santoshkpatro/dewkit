package config

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func SetupDB(ctx context.Context) error {
	if DB != nil {
		return nil // already initialized
	}

	config, err := pgxpool.ParseConfig(GetEnv("DB_URL"))
	if err != nil {
		return  err
	}

	config.MaxConns = 20
	config.MinConns = 5
	config.MaxConnLifetime = time.Hour
	config.MaxConnIdleTime = 30 * time.Minute

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return err
	}

	// Verify connection
	if err := pool.Ping(ctx); err != nil {
		return err
	}

	fmt.Println("DB Connected ...")

	DB = pool
	return nil
}

func GetDB(ctx context.Context) (*pgx.Conn, error) {
	return pgx.Connect(ctx, GetEnv("DB_URL"))
}

func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}