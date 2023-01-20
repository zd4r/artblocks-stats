package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	"github.com/zd4rova/artblocks-holders/internal/usecase"
)

func NewRouter(handler *echo.Echo, l *zerolog.Logger, c usecase.Collection) {

	// Middleware
	handler.Use(middleware.Recover())
	handler.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			l.Info().
				Str("URI", v.URI).
				Int("status", v.Status).
				Str("latency", v.Latency.String()).
				Msg("request")

			return nil
		},
	}))

	// Routers
	h := handler.Group("/v1")
	{
		newCollectionRoutes(h, c, l)
	}
}
