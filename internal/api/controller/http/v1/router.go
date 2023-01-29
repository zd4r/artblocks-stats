package v1

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	"github.com/zd4rova/artblocks-stats/internal/api/usecase"
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
				Str("response time", time.Now().Sub(v.StartTime).String()).
				Msg("request")
			return nil
		},
	}))

	// Routers
	h := handler.Group("/v1")
	{
		newCollectionsRoutes(h, c, l)
	}
}
