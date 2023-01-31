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

// newCollectionsRoutes creates new routes based on collectionRoutes
func newCollectionsRoutes(handler *echo.Group, c usecase.Collection, l *zerolog.Logger) {
	r := &collectionRoutes{c, l}

	handler.GET("/collections/:id/stats", r.collectionStats())
	handler.GET("/collections/:id/holders", r.collectionHolders())
}

// envelope is helper type for neat response
type envelope map[string]interface{}

// collectionStats gather collection holders with scores and
// calculate collection holders distribution based on scores
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

		collection, err = cr.c.СalculateStats(c.Request().Context(), collection)
		if err != nil {
			cr.l.Err(err).Msg("http - v1 - collectionStats - cr.c.Stats")
			return echo.NewHTTPError(http.StatusInternalServerError, "failed getting collection stats")
		}

		return c.JSON(http.StatusOK, envelope{"collection": struct {
			ID                  int                        `json:"id"`
			HoldersCount        int                        `json:"holders_count"`
			HoldersDistribution entity.HoldersDistribution `json:"holders_distribution"`
		}{collection.ID, collection.HoldersCount, collection.HoldersDistribution}})
	}
}

// collectionHolders gather collection holders with scores
func (cr *collectionRoutes) collectionHolders() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		collectionID, err := strconv.Atoi(id)
		if err != nil {
			cr.l.Err(err).Msg("http - v1 - collectionStats - strconv.Atoi")
			return echo.NewHTTPError(http.StatusBadRequest, "invalid collection id")
		}

		collection := entity.Collection{
			ID: collectionID,
		}

		collection, err = cr.c.GetHolders(c.Request().Context(), collection)
		if err != nil {
			cr.l.Err(err).Msg("http - v1 - collectionHolders - cr.c.GetHolders")
			return echo.NewHTTPError(http.StatusInternalServerError, "failed getting collection holders")
		}

		return c.JSON(http.StatusOK, envelope{"holders": collection.Holders})
	}
}
