package transaction

import (
	"context"
	"github.com/pkg/errors"
	"time"
)

func (r *TransactionRepository) CreatePayment(ctx context.Context, transactionID int64, transactionHash, userID, email, amount, currency, status string) error {
	rows, err := r.ConnPool.Query(ctx, "INSERT INTO transaction (transaction_id, "+
		"transaction_hash, "+
		"user_id, "+
		"email, "+
		"amount, "+
		"currency, "+
		"status"+
		") "+
		"VALUES ($1,$2,$3,$4,$5,$6,$7);",
		transactionID,
		transactionHash,
		userID,
		email,
		amount,
		currency,
		status)
	if err != nil {
		return errors.Wrap(err, "create transaction")
	}

	defer rows.Close()

	return nil
}

func (r *TransactionRepository) ChangeStatus(ctx context.Context, transactionID int64, status string) error {
	now := time.Now()
	rows, err := r.ConnPool.Query(ctx, "UPDATE transaction SET transaction_id = $1, status = $2, date_of_last_change = $3;",
		transactionID,
		status,
		now)
	if err != nil {
		return errors.Wrap(err, "change status")
	}

	defer rows.Close()

	return nil
}

func (r *TransactionRepository) CheckStatus(ctx context.Context, transactionID int64) (string, error) {
	var status string
	rows, err := r.ConnPool.Query(ctx, "SELECT status FROM transaction WHERE transaction_id = $1;", transactionID)
	if err != nil {
		return "", errors.Wrap(err, "check status")
	}

	for rows.Next() {
		err = rows.Scan(
			&status,
		)
	}

	if err != nil {
		return "", errors.Wrap(err, "failed to scan")
	}

	defer rows.Close()

	return status, nil
}

func (r *TransactionRepository) GetTransactionHashFromID(ctx context.Context, transactionID int64, userID string) (string, error) {
	var transactionHash string

	rows, err := r.ConnPool.Query(ctx, "SELECT transaction_hash FROM transaction WHERE transaction_id = $1 AND user_id = $2;",
		transactionID,
		userID)
	if err != nil {
		return "", errors.Wrap(err, "select transaction hash")
	}

	for rows.Next() {
		err = rows.Scan(
			&transactionHash,
		)
	}

	if err != nil {
		return "", errors.Wrap(err, "failed to scan")
	}

	defer rows.Close()

	return transactionHash, nil
}

func (r *TransactionRepository) GetPaymentsByID(ctx context.Context, userID string) ([]*Transaction, error) {
	payment := &Transaction{
		TransactionID:    0,
		TransactionHash:  "",
		UserID:           "",
		Email:            "",
		Amount:           "",
		Currency:         "",
		DateOfCreation:   time.Time{},
		DateOfLastChange: time.Time{},
		Status:           "",
	}

	var payments []*Transaction
	rows, err := r.ConnPool.Query(ctx, "SELECT * FROM transaction WHERE user_id = $1;", userID)
	if err != nil {
		return nil, errors.Wrap(err, "check status")
	}

	for rows.Next() {
		err = rows.Scan(
			&payment.TransactionID,
			&payment.TransactionHash,
			&payment.UserID,
			&payment.Email,
			&payment.Amount,
			&payment.Currency,
			&payment.DateOfCreation,
			&payment.DateOfLastChange,
			&payment.Status,
		)
		if err != nil {
			return nil, errors.Wrap(err, "failed to scan")
		}

		payments = append(payments, payment)
	}

	defer rows.Close()

	return payments, nil
}

func (r *TransactionRepository) GetPaymentsByEmail(ctx context.Context, email string) ([]*Transaction, error) {
	payment := &Transaction{
		TransactionID:    0,
		TransactionHash:  "",
		UserID:           "",
		Email:            "",
		Amount:           "",
		Currency:         "",
		DateOfCreation:   time.Time{},
		DateOfLastChange: time.Time{},
		Status:           "",
	}

	var payments []*Transaction
	rows, err := r.ConnPool.Query(ctx, "SELECT * FROM transaction WHERE email = $1;", email)
	if err != nil {
		return nil, errors.Wrap(err, "check status")
	}

	for rows.Next() {
		err = rows.Scan(
			&payment.TransactionID,
			&payment.TransactionHash,
			&payment.UserID,
			&payment.Email,
			&payment.Amount,
			&payment.Currency,
			&payment.DateOfCreation,
			&payment.DateOfLastChange,
			&payment.Status,
		)
		if err != nil {
			return nil, errors.Wrap(err, "failed to scan")
		}

		payments = append(payments, payment)
	}

	defer rows.Close()

	return payments, nil
}

func (r *TransactionRepository) CancelTransaction(ctx context.Context, transactionID int64) error {
	row, err := r.ConnPool.Query(ctx, "DELETE FROM transaction WHERE transaction_id = $1;", transactionID)
	if err != nil {
		return errors.Wrap(err, "delete transaction")
	}

	defer row.Close()

	return nil
}
