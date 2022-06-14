package transaction

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	"time"
)

type Transactioner interface {
	CreatePayment(ctx context.Context, transactionID, userID, amount int64, transactionHash, email, currency, status string) error
	ChangeStatus(ctx context.Context, transactionID int64, status string) (string, error)
	CheckStatus(ctx context.Context, transactionID int64) (string, error)
	GetTransactionHashFromID(ctx context.Context, transactionID, userID int64) (string, error)
	GetPaymentsByID(ctx context.Context, userID int64) ([]*Transaction, error)
	GetPaymentsByEmail(ctx context.Context, email string) ([]*Transaction, error)
	CancelTransaction(ctx context.Context, transactionID int64) error
}

type TransactionRepository struct {
	ConnPool *pgxpool.Pool
}

func NewTransactionRepository(ctx context.Context, pool *pgxpool.Pool) (*TransactionRepository, error) {
	transactionRep := TransactionRepository{
		ConnPool: pool,
	}

	rows, err := transactionRep.ConnPool.Query(ctx, `
CREATE TABLE IF NOT EXISTS transaction(
	transaction_id bigint UNIQUE, 
	transaction_hash text UNIQUE, 
	user_id bigint, 
	email text,
	amount bigint,
	currency text,
	date_of_creation timestamp NOT NULL DEFAULT TRANSACTION_TIMESTAMP(),
	date_of_last_change timestamp NOT NULL DEFAULT TRANSACTION_TIMESTAMP(),
	status text);`)
	if err != nil {
		return nil, errors.Wrap(err, "create transaction table")
	}
	rows.Close()

	return &transactionRep, nil
}

type Transaction struct {
	TransactionID    int64     `json:"transaction_id"`
	TransactionHash  string    `json:"transaction_hash"`
	UserID           int64     `json:"user_id"`
	Email            string    `json:"email"`
	Amount           int64     `json:"amount"`
	Currency         string    `json:"currency"`
	DateOfCreation   time.Time `json:"date_of_creation"`
	DateOfLastChange time.Time `json:"date_of_last_change"`
	Status           string    `json:"status"`
}
