package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	pingTimeOut = 5
)

type Config struct {
	PostgresDB       string
	PostgresUser     string
	PostgresPassword string
	PostgresHost     string

	MaxOpenConns int
	MaxIdleConns int
	MaxIdleTime  string
}

type Postgres struct {
	DB *sqlx.DB
}

func New(cfg *Config) (*Postgres, error) {
	pgUrl := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
		cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresHost, cfg.PostgresDB)

	db, err := sqlx.Open("postgres", pgUrl)
	if err != nil {
		return nil, fmt.Errorf("postgres - New - sqlx.Open: %w", err)
	}

	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)

	duration, err := time.ParseDuration(cfg.MaxIdleTime)
	if err != nil {
		return nil, fmt.Errorf("postgres - New - time.ParseDuration: %w", err)
	}
	db.SetConnMaxIdleTime(duration)

	ctx, cancel := context.WithTimeout(context.Background(), pingTimeOut*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("postgres - New - pg.Pool.PingContext: %w", err)
	}

	return &Postgres{
		DB: db,
	}, nil
}
