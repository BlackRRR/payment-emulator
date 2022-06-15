package server

import (
	"context"
	"github.com/BlackRRR/payment-emulator/internal/model"
	"github.com/BlackRRR/payment-emulator/internal/repository/transaction"
	"github.com/pkg/errors"
)

func (s *Server) CreatePayment(ctx context.Context, request *PaymentRequest) (*PaymentResponse, error) {
	payment := &transaction.Transaction{
		UserID:   request.UserID,
		Email:    request.Email,
		Amount:   request.Amount,
		Currency: request.Currency,
	}

	transactionID, transactionHash, status, err := s.trans.CreatePayment(ctx, payment)
	if err != nil {
		return &PaymentResponse{
			Result:  model.ResultERR,
			Payload: &TransactionHashPayload{transactionID, transactionHash, status},
			Error:   model.NewTransactionError(err),
		}, nil
	}

	return &PaymentResponse{
		Result:  model.ResultOK,
		Payload: &TransactionHashPayload{transactionID, transactionHash, status},
		Error:   nil,
	}, nil
}

func (s *Server) ChangePaymentStatus(ctx context.Context, request *PaymentStatusChangeRequest) (*PaymentStatusChangeResponse, error) {
	status, err := s.trans.ChangePaymentStatus(ctx, request.TransactionID, request.UserID, request.TransactionHash)
	if status == model.StatusFailure {
		return &PaymentStatusChangeResponse{
			Result:  model.ResultERR,
			Payload: &TransactionStatusPayload{status},
			Error:   model.NewTransactionError(err),
		}, err
	}

	if err != nil {
		return &PaymentStatusChangeResponse{
			Result:  model.ResultERR,
			Payload: &TransactionStatusPayload{status},
			Error:   model.NewTransactionError(err),
		}, err
	}

	return &PaymentStatusChangeResponse{
		Result:  model.ResultOK,
		Payload: &TransactionStatusPayload{status},
		Error:   nil,
	}, nil
}

func (s *Server) CheckPaymentStatus(ctx context.Context, request *PaymentStatusCheckRequest) (*PaymentStatusCheckResponse, error) {
	status, err := s.trans.CheckPaymentStatus(ctx, request.TransactionID)
	if err != nil {
		return &PaymentStatusCheckResponse{
			Result:  model.ResultERR,
			Payload: &TransactionStatusPayload{status},
			Error:   model.NewTransactionError(err),
		}, err
	}

	return &PaymentStatusCheckResponse{
		Result:  model.ResultOK,
		Payload: &TransactionStatusPayload{status},
		Error:   nil,
	}, nil
}

func (s *Server) GetAllPaymentsByID(ctx context.Context, request *PaymentGetFromIDRequest) (*PaymentGetFromIDResponse, error) {
	payments, err := s.trans.GetAllPaymentByID(ctx, request.UserID)
	if err != nil {
		return &PaymentGetFromIDResponse{
			Result:  model.ResultERR,
			Payload: nil,
			Error:   model.NewTransactionError(err),
		}, err
	}

	return &PaymentGetFromIDResponse{
		Result:  model.ResultOK,
		Payload: &Payments{payments},
		Error:   nil,
	}, nil

}

func (s *Server) GetAllPaymentsByEmail(ctx context.Context, request *PaymentGetFromEmailRequest) (*PaymentGetFromEmailResponse, error) {
	payments, err := s.trans.GetAllPaymentByEmail(ctx, request.Email)
	if err != nil {
		return &PaymentGetFromEmailResponse{
			Result:  model.ResultERR,
			Payload: nil,
			Error:   model.NewTransactionError(err),
		}, err
	}

	return &PaymentGetFromEmailResponse{
		Result:  model.ResultOK,
		Payload: &Payments{payments},
		Error:   nil,
	}, nil

}

func (s *Server) CancelTransaction(ctx context.Context, request *PaymentCancelRequest) (*PaymentCancelResponse, error) {
	status, err := s.trans.CancelTransaction(ctx, request.TransactionID)
	if err != nil {
		return &PaymentCancelResponse{
			Result:  model.ResultERR,
			Payload: &TransactionStatusPayload{status},
			Error:   model.NewTransactionError(err),
		}, err
	}

	if status == model.StatusSuccess {
		return &PaymentCancelResponse{
			Result:  model.ResultERR,
			Payload: &TransactionStatusPayload{status},
			Error:   model.NewTransactionError(errors.New("impossible to cancel the payment")),
		}, nil
	}

	if status == model.StatusFailure {
		return &PaymentCancelResponse{
			Result:  model.ResultERR,
			Payload: &TransactionStatusPayload{status},
			Error:   model.NewTransactionError(errors.New("impossible to cancel the payment")),
		}, nil
	}

	return &PaymentCancelResponse{
		Result:  model.ResultOK,
		Payload: &TransactionStatusPayload{status},
		Error:   nil,
	}, nil
}
