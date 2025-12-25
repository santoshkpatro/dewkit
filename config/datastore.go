package config

import (
	"context"
	"fmt"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

var DB *sqlx.DB

var Cache *redis.Client

func SetupDB(ctx context.Context) (*sqlx.DB, error) {
	if DB != nil {
		return DB, nil // already initialized
	}

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

	// Verify connection
	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}

	fmt.Println("DB Connected ...")

	DB = db
	return db, nil
}

func SetupCache(ctx context.Context) (*redis.Client, error) {
	if Cache != nil {
		return Cache, nil
	}

	cacheURL := GetEnv("CACHE_URL")
	opts, err := redis.ParseURL(cacheURL)
	if err != nil {
		return nil, err
	}

	opts.PoolSize = 20
	opts.MinIdleConns = 5
	opts.DialTimeout = 5 * time.Second
	opts.ReadTimeout = 3 * time.Second
	opts.WriteTimeout = 3 * time.Second
	opts.PoolTimeout = 4 * time.Second

	rdb := redis.NewClient(opts)

	if err := rdb.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	fmt.Println("Redis Connected ...")

	Cache = rdb
	return rdb, nil
}

func GetCache(ctx context.Context) (*redis.Client, error) {
	if Cache == nil {
		return SetupCache(ctx)
	}
	return Cache, nil
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

func CloseCache() {
	if Cache != nil {
		_ = Cache.Close()
	}
}
