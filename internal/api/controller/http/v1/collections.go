package v1

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/zd4rova/artblocks-stats/internal/api/entity"
	"github.com/zd4rova/artblocks-stats/internal/api/usecase"
)

type collectionRoutes struct {
	c usecase.Collection
	l *zerolog.Logger
}

func newCollectionsRoutes(handler *echo.Group, c usecase.Collection, l *zerolog.Logger) {
	r := &collectionRoutes{c, l}

	handler.GET("/collections/:id", r.collectionStats())
}

type envelope map[string]interface{}

func (cr *collectionRoutes) collectionStats() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		collectionID, err := strconv.Atoi(id)
		if err != nil {
			cr.l.Err(err).Msg("http - v1 - collectionStats - strconv.Atoi")
			return echo.NewHTTPError(http.StatusBadRequest, "invalid collection id")
		}

		collection := entity.Collection{
			ID: collectionID,
			HoldersDistribution: entity.HoldersDistribution{
				ByCommitmentScore: make(map[string]int),
				ByTradingScore:    make(map[string]int),
				ByPortfolioScore:  make(map[string]int),
			},
		}

		collectionStats, err := cr.c.Stats(c.Request().Context(), collection)
		if err != nil {
			cr.l.Err(err).Msg("http - v1 - collectionStats - cr.c.Stats")
			return echo.NewHTTPError(http.StatusInternalServerError, "failed getting collection stats")
		}

		return c.JSON(http.StatusOK, envelope{"collection": collectionStats})
	}
}
