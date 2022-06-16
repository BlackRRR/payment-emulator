package transaction

import (
	"context"
	"github.com/BlackRRR/payment-emulator/internal/model"
	"github.com/BlackRRR/payment-emulator/internal/repository/transaction"
	"github.com/BlackRRR/payment-emulator/internal/utils"
	"github.com/pkg/errors"
	"math/rand"
	"strings"
	"time"
)

func (s *TransactionService) CreatePayment(ctx context.Context, payment *transaction.Payment) (int64, string, string, error) {
	rand.Seed(time.Now().UnixNano())
	id := rand.Int63n(10000000)
	hash := utils.GetUUID()

	//Random number of payments goes to error status
	if rand.Intn(2) == 0 {
		err := s.rep.CreatePayment(ctx, &transaction.Payment{
			TransactionID:   id,
			TransactionHash: hash,
			UserID:          payment.UserID,
			Status:          model.StatusError,
		})
		if err != nil {
			return 0, "", model.StatusError, errors.Wrap(err, "service error: failed to create payment")
		}

		return id, "", model.StatusError, errors.New("Random error status")
	}

	if !strings.Contains(payment.Email, "@") {
		err := s.rep.CreatePayment(ctx, &transaction.Payment{
			TransactionID:   id,
			TransactionHash: hash,
			UserID:          payment.UserID,
			Status:          model.StatusError,
		})
		if err != nil {
			return 0, "", model.StatusError, errors.Wrap(err, "service error: failed to create payment")
		}

		return id, "", model.StatusError, errors.New("Incorrect email")
	}

	if payment.Amount < 0 {
		err := s.rep.CreatePayment(ctx, &transaction.Payment{
			TransactionID:   id,
			TransactionHash: hash,
			UserID:          payment.UserID,
			Status:          model.StatusError,
		})
		if err != nil {
			return 0, "", model.StatusError, errors.Wrap(err, "service error: failed to create payment")
		}

		return id, "", model.StatusError, errors.New("Incorrect amount")
	}

	if payment.Currency == "RUB" || payment.Currency == "DOLLAR" || payment.Currency == "EURO" {
		err := s.rep.CreatePayment(ctx, &transaction.Payment{
			TransactionID:   id,
			TransactionHash: hash,
			UserID:          payment.UserID,
			Email:           payment.Email,
			Amount:          payment.Amount,
			Currency:        payment.Currency,
			Status:          model.StatusNew,
		})
		if err != nil {
			return 0, "", model.StatusError, errors.Wrap(err, "service error: failed to create payment")
		}

		return id, hash, model.StatusNew, nil
	}

	err := s.rep.CreatePayment(ctx, &transaction.Payment{
		TransactionID:   id,
		TransactionHash: hash,
		UserID:          payment.UserID,
		Status:          model.StatusError,
	})
	if err != nil {
		return 0, "", model.StatusError, errors.Wrap(err, "service error: failed to create payment")
	}

	return id, "", model.StatusError, errors.New("Incorrect currency")
}

func (s *TransactionService) ChangePaymentStatus(ctx context.Context, transactionID, userID int64, transactionHash string) (string, error) {
	hash, err := s.rep.GetTransactionHashFromID(ctx, transactionID, userID)
	if err != nil {
		status, err := s.rep.ChangeStatus(ctx, transactionID, model.StatusFailure)
		if err != nil {
			return status, errors.Wrap(err, "service error: change payment status")
		}

		return status, errors.Wrap(err, "service error: change payment status")
	}

	if hash != transactionHash {
		status, err := s.rep.ChangeStatus(ctx, transactionID, model.StatusFailure)
		if err != nil {
			return status, errors.Wrap(err, "service error: change payment status")
		}

		return status, nil
	}

	status, err := s.rep.ChangeStatus(ctx, transactionID, model.StatusSuccess)
	if err != nil {
		status, err = s.rep.ChangeStatus(ctx, transactionID, model.StatusFailure)
		if err != nil {
			return status, errors.Wrap(err, "service error: change payment status")
		}
	}

	return status, nil
}

func (s *TransactionService) CheckPaymentStatus(ctx context.Context, transactionID int64) (string, error) {
	status, err := s.rep.CheckStatus(ctx, transactionID)
	if err != nil {
		return "", errors.Wrap(err, "service error: check status")
	}

	return status, nil
}

func (s *TransactionService) GetAllPaymentByID(ctx context.Context, UserID int64) ([]*transaction.Payment, error) {
	payments, err := s.rep.GetPaymentsByID(ctx, UserID)
	if err != nil {
		return nil, errors.Wrap(err, "service error: get payments by ID")
	}

	return payments, nil
}

func (s *TransactionService) GetAllPaymentByEmail(ctx context.Context, email string) ([]*transaction.Payment, error) {
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

	if status == model.StatusFailure {
		return model.StatusFailure, nil
	}

	err = s.rep.CancelTransaction(ctx, transactionID)
	if err != nil {
		return "", errors.Wrap(err, "service error: cancel payment")
	}

	status, err = s.rep.ChangeStatus(ctx, transactionID, model.StatusCanceled)
	if err != nil {
		return "", errors.Wrap(err, "service error: cancel payment")
	}

	return status, nil
}
