package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"github.com/zd4r/artblocks-stats/cmd/api/config"
	"github.com/zd4r/artblocks-stats/internal/api/controller/http/v1"
	"github.com/zd4r/artblocks-stats/internal/api/usecase"
	"github.com/zd4r/artblocks-stats/internal/api/usecase/repo"
	"github.com/zd4r/artblocks-stats/internal/api/usecase/webapi"
	"github.com/zd4r/artblocks-stats/pkg/httpserver"
	"github.com/zd4r/artblocks-stats/pkg/logger"
	"github.com/zd4r/artblocks-stats/pkg/postgres"
)

// Run is main function continuation, which starts app
func Run(cfg *config.Config) {
	// Logger
	l := logger.New(cfg.Log.Level, cfg.Log.Structured)
	l.Info().Msg(fmt.Sprintf("⚡ init app [ name: %v, version: %v ]", cfg.Name, cfg.Version))

	// Repository
	dbCfg := &postgres.Config{
		URL:          cfg.PG.URL,
		MaxOpenConns: cfg.PG.MaxOpenConns,
		MaxIdleConns: cfg.PG.MaxIdleConns,
		MaxIdleTime:  cfg.PG.MaxIdleTime,
	}
	pg, err := postgres.New(dbCfg)
	if err != nil {
		l.Fatal().Msg(fmt.Errorf("app - Run - postgres.New: %w", err).Error())
	}

	// Use case
	collectionUseCase := usecase.NewCollection(
		repo.New(pg.DB),
		webapi.New(),
	)

	// HTTP server
	handler := echo.New()
	v1.NewRouter(handler, l, collectionUseCase)
	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))
	l.Info().Msg(fmt.Sprintf("🌏 http server started on :%v", cfg.HTTP.Port))

	// Graceful shutdown
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGTERM, os.Interrupt)

	select {
	case s := <-interrupt:
		l.Info().Msg(fmt.Sprintf("app - Run - signal: %v", s.String()))
	case err = <-httpServer.Notify():
		l.Err(err).Msg("app - Run - httpServer.Notify")
	}

	err = httpServer.Shutdown()
	if err != nil {
		l.Err(err).Msg("app - Run - httpServer.Shutdown")
	}
}
