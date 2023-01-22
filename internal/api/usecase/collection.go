package usecase

import (
	"context"
	"fmt"

	"github.com/zd4rova/artblocks-stats/internal/api/entity"
)

type CollectionUseCase struct {
	webAPI CollectionWebAPI
}

func NewCollection(w CollectionWebAPI) *CollectionUseCase {
	return &CollectionUseCase{
		webAPI: w,
	}
}

func (uc *CollectionUseCase) СalculateStats(ctx context.Context, c entity.Collection) (entity.Collection, error) {

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

	err = collection.CountHoldersDistribution()
	if err != nil {
		return entity.Collection{}, fmt.Errorf("CollectionUseCase - Stats - collectionHoldersDistribution: %w", err)
	}

	return collection, nil
}
