package transaction

import (
	"context"
	"github.com/pkg/errors"
	"time"
)

func (r *TransactionRepository) CreatePayment(ctx context.Context, payment *Transaction) error {
	_, err := r.ConnPool.Exec(ctx, `
INSERT INTO transaction 
		(transaction_id, 
		transaction_hash, 
		user_id, 
		email, 
		amount, 
		currency,
		status)
VALUES ($1,$2,$3,$4,$5,$6,$7);`,
		payment.TransactionID,
		payment.TransactionHash,
		payment.UserID,
		payment.Email,
		payment.Amount,
		payment.Currency,
		payment.Currency)
	if err != nil {
		return errors.Wrap(err, "create transaction")
	}

	return nil
}

func (r *TransactionRepository) ChangeStatus(ctx context.Context, transactionID int64, status string) (string, error) {
	var newStatus string
	now := time.Now()
	err := r.ConnPool.QueryRow(ctx, `
UPDATE transaction SET status = $1, date_of_last_change = $2
	WHERE transaction_id = $3 RETURNING status;`,
		status,
		now,
		transactionID).Scan(&newStatus)
	if err != nil {
		return "", errors.Wrap(err, "change status")
	}

	return newStatus, nil
}

func (r *TransactionRepository) CheckStatus(ctx context.Context, transactionID int64) (string, error) {
	var status string
	err := r.ConnPool.QueryRow(ctx, "SELECT status FROM transaction WHERE transaction_id = $1;",
		transactionID).Scan(&status)
	if err != nil {
		return "", errors.Wrap(err, "check status")
	}

	return status, nil
}

func (r *TransactionRepository) GetTransactionHashFromID(ctx context.Context, transactionID, userID int64) (string, error) {
	var transactionHash string

	err := r.ConnPool.QueryRow(ctx, `
SELECT transaction_hash FROM transaction 
	WHERE transaction_id = $1 AND user_id = $2;`,
		transactionID,
		userID).Scan(&transactionHash)
	if err != nil {
		return "", errors.Wrap(err, "select transaction hash")
	}

	return transactionHash, nil
}

func (r *TransactionRepository) GetPaymentsByID(ctx context.Context, userID int64) ([]*Transaction, error) {
	payment := &Transaction{}
	var payments []*Transaction

	rows, err := r.ConnPool.Query(ctx, `
SELECT * FROM transaction 
	WHERE user_id = $1;`, userID)
	if err != nil {
		return nil, errors.Wrap(err, "check status")
	}
	defer rows.Close()

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

	return payments, nil
}

func (r *TransactionRepository) GetPaymentsByEmail(ctx context.Context, email string) ([]*Transaction, error) {
	payment := &Transaction{}
	var payments []*Transaction

	rows, err := r.ConnPool.Query(ctx, `
SELECT * FROM transaction
	WHERE email = $1;`, email)
	if err != nil {
		return nil, errors.Wrap(err, "check status")
	}
	defer rows.Close()

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

	return payments, nil
}

func (r *TransactionRepository) CancelTransaction(ctx context.Context, transactionID int64) error {
	_, err := r.ConnPool.Exec(ctx, `
DELETE FROM transaction 
	WHERE transaction_id = $1;`, transactionID)
	if err != nil {
		return errors.Wrap(err, "delete transaction")
	}

	return nil
}
