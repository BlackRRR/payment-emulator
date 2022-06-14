package transaction

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"math/rand"
	"payment-emulator/internal/model"
	"payment-emulator/internal/repository/transaction"
	"payment-emulator/internal/utils"
	"time"
)

func (s *TransactionService) CreatePayment(ctx context.Context, userID, email, amount, currency string) (int64, string, error) {
	rand.Seed(time.Now().UnixNano())
	id := int64(rand.Intn(10000000))
	hash := utils.GetHash()

	fmt.Println(id)

	err := s.rep.CreatePayment(ctx, id, hash, userID, email, amount, currency, model.StatusNew)
	if err != nil {
		return 0, "", errors.Wrap(err, "service error: failed to create payment")
	}

	return id, hash, nil
}

func (s *TransactionService) ChangePaymentStatus(ctx context.Context, transactionID int64, transactionHash, userID string) (string, error) {
	hash, err := s.rep.GetTransactionHashFromID(ctx, transactionID, userID)
	if err != nil {
		return "", errors.Wrap(err, "service error: get transaction hash")
	}

	if hash != transactionHash {
		err = s.rep.ChangeStatus(ctx, transactionID, model.StatusFailure)
		if err != nil {
			return model.StatusFailure, errors.Wrap(err, "service error: change payment status")
		}

		return model.StatusFailure, nil
	}

	err = s.rep.ChangeStatus(ctx, transactionID, model.StatusSuccess)
	if err != nil {
		return "", errors.Wrap(err, "service error: failed to change status")
	}

	return model.StatusSuccess, nil
}

func (s *TransactionService) CheckPaymentStatus(ctx context.Context, transactionID int64) (string, error) {
	status, err := s.rep.CheckStatus(ctx, transactionID)
	if err != nil {
		return "", errors.Wrap(err, "service error: check status")
	}

	return status, nil
}

func (s *TransactionService) GetAllPaymentByID(ctx context.Context, UserID string) ([]*transaction.Transaction, error) {
	payments, err := s.rep.GetPaymentsByID(ctx, UserID)
	if err != nil {
		return nil, errors.Wrap(err, "service error: get payments by ID")
	}

	return payments, nil

}

func (s *TransactionService) GetAllPaymentByEmail(ctx context.Context, email string) ([]*transaction.Transaction, error) {
	payments, err := s.rep.GetPaymentsByEmail(ctx, email)
	if err != nil {
		return nil, errors.Wrap(err, "service error: get payments by email")
	}

	return payments, nil

}

func (s *TransactionService) CancelTransaction(ctx context.Context, transactionID int64) (string, error) {
	status, err := s.rep.CheckStatus(ctx, transactionID)
	if err != nil {
		return "", errors.Wrap(err, "service error: check payment status")
	}

	if status == model.StatusSuccess {
		return model.StatusSuccess, nil
	}

	err = s.rep.CancelTransaction(ctx, transactionID)
	if err != nil {
		return "", errors.Wrap(err, "service error: cancel payment")
	}

	return "", nil

}
