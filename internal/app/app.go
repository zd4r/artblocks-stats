package app

import (
	"github.com/labstack/echo/v4"
	"github.com/zd4rova/artblocks-holders/config"
	v1 "github.com/zd4rova/artblocks-holders/internal/controller/http/v1"
	"github.com/zd4rova/artblocks-holders/internal/usecase"
	"github.com/zd4rova/artblocks-holders/internal/usecase/webapi"
	"github.com/zd4rova/artblocks-holders/pkg/logger"
)

func Run(cfg *config.Config) {
	// Logger
	l := logger.New(cfg.Log.Level)
	//l.Info().Msg("Hello world")
	//
	//e := echo.New()
	//
	//middleware.RequestID()
	//
	//e.GET("/", func(c echo.Context) error {
	//	return c.String(http.StatusOK, "Hello, World!")
	//})
	//e.Use(middleware.Logger())
	//e.Logger.Fatal(e.Start(":8080"))

	// Use case
	collectionUseCase := usecase.NewCollection(
		webapi.New(),
	)

	handler := echo.New()
	v1.NewRouter(handler, l, collectionUseCase)
	l.Fatal().Msg(handler.Start("8080").Error())
}
