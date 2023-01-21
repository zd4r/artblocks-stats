package usecase

import (
	"context"

	"github.com/zd4rova/artblocks-stats/internal/api/entity"
)

type (
	CollectionWebAPI interface {
		GetHoldersCount(entity.Collection) (entity.Collection, error)
		GetHolders(entity.Collection) (entity.Collection, error)
		GetHolderScores(entity.Holder) (entity.Holder, error)
	}

	Collection interface {
		Stats(context.Context, entity.Collection) (entity.Collection, error)
	}
)
