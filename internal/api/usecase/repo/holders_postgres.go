package repo

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/zd4rova/artblocks-stats/internal/api/entity"
)

type HoldersRepo struct {
	DB *sqlx.DB
}

func New(db *sqlx.DB) *HoldersRepo {
	return &HoldersRepo{db}
}

func (r *HoldersRepo) InsertHolder(holder *entity.Holder) error {
	query :=
		`INSERT INTO holders (updated_at, address, commitment_score, portfolio_score, trading_score)
		 VALUES ($1, $2, $3, $4, $5)
		 RETURNING id, updated_at, version`

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

	return r.DB.QueryRowContext(ctx, query, args...).Scan(&holder.ID, &holder.UpdatedAt, &holder.Version)
}

func (r *HoldersRepo) GetHolder(holder *entity.Holder) (*entity.Holder, error) {
	if holder.ID < 1 {
		return nil, ErrRecordNotFound
	}

	query :=
		`SELECT id, updated_at, address, commitment_score, portfolio_score, trading_score, version
		 FROM holders
		 WHERE id = $1`

	args := []interface{}{
		holder.ID,
	}

	// TODO Fix: timeout hardcode
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := r.DB.QueryRowContext(ctx, query, args...).Scan(
		&holder.ID,
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
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return holder, nil
}

func (r *HoldersRepo) UpdateHolder(holder *entity.Holder) error {
	query :=
		`UPDATE holders
		 SET updated_at = $1, commitment_score = $2, portfolio_score = $3, trading_score = $4, version = version + 1
		 WHERE id = $6 AND version = $6
		 RETURNING version`

	args := []interface{}{
		time.Now(),
		holder.CommitmentScore,
		holder.PortfolioScore,
		holder.TradingScore,
		holder.ID,
		holder.Version,
	}

	// TODO Fix: timeout hardcode
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := r.DB.QueryRowContext(ctx, query, args...).Scan(&holder.Version)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrEditConflict
		default:
			return err
		}
	}

	return nil
}
