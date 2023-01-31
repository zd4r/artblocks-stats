package usecase

import (
	"context"

	"github.com/zd4rova/artblocks-stats/internal/api/entity"
)

type (
	HoldersRepo interface {
		Insert(entity.Holder) (entity.Holder, error)
		Get(entity.Holder) (entity.Holder, error)
		UpdateScores(entity.Holder) (entity.Holder, error)
	}

	CollectionWebAPI interface {
		GetHoldersCount(entity.Collection) (entity.Collection, error)
		GetHolders(entity.Collection) (entity.Collection, error)
		GetHolderScores(entity.Holder) (entity.Holder, error)
	}

	Collection interface {
		СalculateStats(context.Context, entity.Collection) (entity.Collection, error)
		GetHolders(context.Context, entity.Collection) (entity.Collection, error)
	}
)
