package repo

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/jmoiron/sqlx"

	"github.com/zd4r/artblocks-stats/internal/api/entity"
)

type HoldersRepo struct {
	DB *sqlx.DB
}

// New creates new repo based on provided db
func New(db *sqlx.DB) *HoldersRepo {
	return &HoldersRepo{db}
}

// Insert adds new holder into db
func (r *HoldersRepo) Insert(holder entity.Holder) (entity.Holder, error) {
	query :=
		`INSERT INTO holders (updated_at, address, commitment_score, portfolio_score, trading_score)
		 VALUES ($1, $2, $3, $4, $5)
		 RETURNING updated_at, version`

	args := []interface{}{
		time.Now(),
		holder.Address,
		holder.CommitmentScore,
		holder.PortfolioScore,
		holder.TradingScore,
	}

	// TODO Fix: timeout hardcode
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := r.DB.QueryRowContext(ctx, query, args...).Scan(&holder.UpdatedAt, &holder.Version)
	if err != nil {
		return entity.Holder{}, err
	}

	return holder, nil
}

// Get returns holder based on provided holder address
func (r *HoldersRepo) Get(holder entity.Holder) (entity.Holder, error) {
	query :=
		`SELECT updated_at, address, commitment_score, portfolio_score, trading_score, version
		 FROM holders
		 WHERE address = $1`

	args := []interface{}{
		holder.Address,
	}

	// TODO Fix: timeout hardcode
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := r.DB.QueryRowContext(ctx, query, args...).Scan(
		&holder.UpdatedAt,
		&holder.Address,
		&holder.CommitmentScore,
		&holder.PortfolioScore,
		&holder.TradingScore,
		&holder.Version,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return entity.Holder{}, entity.ErrHolderNotFound
		default:
			return entity.Holder{}, err
		}
	}

	return holder, nil
}

// UpdateScores updates holder's scores based on provided holder address
func (r *HoldersRepo) UpdateScores(holder entity.Holder) (entity.Holder, error) {
	query :=
		`UPDATE holders
		 SET updated_at = $1, commitment_score = $2, portfolio_score = $3, trading_score = $4, version = version + 1
		 WHERE address = $5 AND version = $6
		 RETURNING version`

	args := []interface{}{
		time.Now(),
		holder.CommitmentScore,
		holder.PortfolioScore,
		holder.TradingScore,
		holder.Address,
		holder.Version,
	}

	// TODO Fix: timeout hardcode
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := r.DB.QueryRowContext(ctx, query, args...).Scan(&holder.Version)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return entity.Holder{}, entity.ErrUpdateScoresConflict
		default:
			return entity.Holder{}, err
		}
	}

	return holder, nil
}
