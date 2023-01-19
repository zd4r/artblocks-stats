package usecase

import (
	"context"
	"fmt"

	"github.com/zd4rova/artblocks-holders/internal/entity"
)

type CollectionUseCase struct {
	webAPI CollectionWebAPI
}

func NewCollection(w CollectionWebAPI) *CollectionUseCase {
	return &CollectionUseCase{
		webAPI: w,
	}
}

func (uc *CollectionUseCase) Stats(ctx context.Context, c entity.Collection) (entity.Collection, error) {

	collection, err := uc.webAPI.GetHoldersCount(c)
	if err != nil {
		return entity.Collection{}, fmt.Errorf("CollectionUseCase - GetHoldersCount - uc.webAPI.GetHolders: %w", err)
	}

	collection.Holders = make([]entity.Holder, collection.HoldersCount)

	collection, err = uc.webAPI.GetHolders(collection)
	if err != nil {
		return entity.Collection{}, fmt.Errorf("CollectionUseCase - Stats - uc.webAPI.GetHolders: %w", err)
	}

	for i, h := range collection.Holders {
		holder, err := uc.webAPI.GetHolderScores(h)
		if err != nil {
			return entity.Collection{}, fmt.Errorf("CollectionUseCase - Stats - uc.webAPI.GetHolderScores: %w", err)
		}

		collection.Holders[i] = holder
	}

	err = collectionHoldersDistribution(&collection)
	if err != nil {
		return entity.Collection{}, fmt.Errorf("CollectionUseCase - Stats - collectionHoldersDistribution: %w", err)
	}

	return collection, nil
}

func collectionHoldersDistribution(c *entity.Collection) error {
	for _, h := range c.Holders {
		switch {
		case h.CommitmentScore < 3.5:
			c.HoldersDistribution.ByCommitmentScore["[3 - 3.5)"] += 1
		case h.CommitmentScore >= 3.5 && h.CommitmentScore < 4:
			c.HoldersDistribution.ByCommitmentScore["[3.5 - 4)"] += 1
		case h.CommitmentScore >= 4 && h.CommitmentScore < 4.5:
			c.HoldersDistribution.ByCommitmentScore["[4 - 4.5)"] += 1
		case h.CommitmentScore >= 4.5 && h.CommitmentScore <= 5:
			c.HoldersDistribution.ByCommitmentScore["[4 - 4.5]"] += 1
		}

		switch {
		case h.PortfolioScore < 3.5:
			c.HoldersDistribution.ByPortfolioScore["[3 - 3.5)"] += 1
		case h.PortfolioScore >= 3.5 && h.PortfolioScore < 4:
			c.HoldersDistribution.ByPortfolioScore["[3.5 - 4)"] += 1
		case h.PortfolioScore >= 4 && h.PortfolioScore < 4.5:
			c.HoldersDistribution.ByPortfolioScore["[4 - 4.5)"] += 1
		case h.PortfolioScore >= 4.5 && h.PortfolioScore <= 5:
			c.HoldersDistribution.ByPortfolioScore["[4 - 4.5]"] += 1
		}

		switch {
		case h.TradingScore < 3.5:
			c.HoldersDistribution.ByTradingScore["[3 - 3.5)"] += 1
		case h.TradingScore >= 3.5 && h.TradingScore < 4:
			c.HoldersDistribution.ByTradingScore["[3.5 - 4)"] += 1
		case h.TradingScore >= 4 && h.TradingScore < 4.5:
			c.HoldersDistribution.ByTradingScore["[4 - 4.5)"] += 1
		case h.TradingScore >= 4.5 && h.TradingScore <= 5:
			c.HoldersDistribution.ByTradingScore["[4 - 4.5]"] += 1
		}
	}

	return nil
}
