package config

import (
	"context"
	"fmt"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func SetupDB(ctx context.Context) error {
	if DB != nil {
		return nil // already initialized
	}

	dsn := GetEnv("DB_URL")

	db, err := sqlx.ConnectContext(ctx, "postgres", dsn)
	if err != nil {
		return err
	}

	// Pool configuration
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(time.Hour)
	db.SetConnMaxIdleTime(30 * time.Minute)

	// Verify connection
	if err := db.PingContext(ctx); err != nil {
		return err
	}

	fmt.Println("DB Connected ...")

	DB = db
	return nil
}

func GetDB(ctx context.Context) (*sqlx.DB, error) {
	dsn := GetEnv("DB_URL")

	db, err := sqlx.ConnectContext(ctx, "pgx", dsn)
	if err != nil {
		return nil, err
	}

	// Pool configuration
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(time.Hour)
	db.SetConnMaxIdleTime(30 * time.Minute)

	if err := db.PingContext(ctx); err != nil {
		_ = db.Close()
		return nil, err
	}

	fmt.Println("DB Connected ...")
	return db, nil
}

func CloseDB() {
	if DB != nil {
		_ = DB.Close()
	}
}
