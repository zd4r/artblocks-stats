package v1

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/zd4r/artblocks-stats/internal/api/entity"
	"github.com/zd4r/artblocks-stats/internal/api/usecase"
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

// errResp is helper type for swagger docs
type errResp struct {
	Message string `json:"message"`
}

// collectionStatsResponse is response structure from collectionStats method
type collectionStatsResponse struct {
	Collection struct {
		ID                  int                        `json:"id"`
		HoldersCount        int                        `json:"holders_count"`
		HoldersDistribution entity.HoldersDistribution `json:"holders_distribution"`
	} `json:"collection"`
}

// collectionStats gather calculate collection holders distribution based on their scores
// @Summary     Show collection stats
// @Description Show collection holders distribution based on artacle scores
// @ID          stats
// @Accept      json
// @Produce     json
// @Param       id path int true "Collection ID from Artacle"
// @Success     200 {object} collectionStatsResponse
// @Failure     400 {object} errResp
// @Failure     500 {object} errResp
// @Router      /collections/{id}/stats [get]
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

		var resp collectionStatsResponse
		resp.Collection = struct {
			ID                  int                        `json:"id"`
			HoldersCount        int                        `json:"holders_count"`
			HoldersDistribution entity.HoldersDistribution `json:"holders_distribution"`
		}{
			collection.ID,
			collection.HoldersCount,
			collection.HoldersDistribution}

		return c.JSON(http.StatusOK, resp)
	}
}

type collectionHoldersResponse struct {
	Holders []entity.Holder `json:"holders"`
}

// collectionHolders gather collection holders with scores
// @Summary     Show collection holders
// @Description Show collection holders with scores from artacle API
// @ID          holders
// @Accept      json
// @Produce     json
// @Param       id path int true "Collection ID from Artacle"
// @Success     200 {object} collectionHoldersResponse
// @Failure     400 {object} errResp
// @Failure     500 {object} errResp
// @Router      /collections/{id}/holders [get]
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

		var resp collectionHoldersResponse
		resp.Holders = collection.Holders
		return c.JSON(http.StatusOK, resp)
	}
}
