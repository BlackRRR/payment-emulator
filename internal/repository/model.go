package repository

import (
	"context"
	"github.com/BlackRRR/payment-emulator/internal/repository/transaction"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"

	_ "github.com/lib/pq"
)

type Repositories struct {
	Trans transaction.Transactioner
}

func InitRepositories(ctx context.Context, pool *pgxpool.Pool) (*Repositories, error) {
	trans, err := transaction.NewTransactionRepository(ctx, pool)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create repository")
	}

	repositories := &Repositories{
		Trans: trans,
	}

	return repositories, nil
}
