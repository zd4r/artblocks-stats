package app

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/zd4rova/artblocks-holders/config"
	"github.com/zd4rova/artblocks-holders/pkg/logger"
)

func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)
	l.Info().Msg("Hello world")

	e := echo.New()

	middleware.RequestID()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Use(middleware.Logger())
	e.Logger.Fatal(e.Start(":8080"))
}
