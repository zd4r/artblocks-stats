package app

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/zd4rova/artblocks-holders/config"
	v1 "github.com/zd4rova/artblocks-holders/internal/controller/http/v1"
	"github.com/zd4rova/artblocks-holders/internal/usecase"
	"github.com/zd4rova/artblocks-holders/internal/usecase/webapi"
	"github.com/zd4rova/artblocks-holders/pkg/logger"
	"github.com/zd4rova/artblocks-holders/pkg/postgres"
)

func Run(cfg *config.Config) {
	// Logger
	l := logger.New(cfg.Log.Level)

	// Repository
	dbCfg := &postgres.Config{
		MaxOpenConns: cfg.MaxOpenConns,
		MaxIdleConns: cfg.MaxIdleConns,
		MaxIdleTime:  cfg.MaxIdleTime,
	}
	_, err := postgres.New(dbCfg)
	if err != nil {
		l.Fatal().Err(fmt.Errorf("app - Run - postgres.New: %w", err))
	}

	// Use case
	collectionUseCase := usecase.NewCollection(
		webapi.New(),
	)

	handler := echo.New()

	v1.NewRouter(handler, l, collectionUseCase)

	l.Fatal().Msg(handler.Start(":8080").Error())
}
