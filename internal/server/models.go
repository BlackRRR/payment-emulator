package server

import (
	"payment-emulator/internal/model"
	"payment-emulator/internal/repository/transaction"
)

//////////////////////////
//Payment Request
//////////////////////////

type PaymentRequest struct {
	UserID   string `json:"user_id"`
	Email    string `json:"email"`
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

type PaymentResponse struct {
	Result  model.Result            `json:"result"`
	Payload *TransactionHashPayload `json:"payload"`
	Error   *model.ServerError      `json:"error"`
}

type TransactionHashPayload struct {
	TransactionID   int64  `json:"transaction_id"`
	TransactionHash string `json:"transaction_hash"`
}

//////////////////////////
//PaymentStatus Request
//////////////////////////

type PaymentStatusChangeRequest struct {
	UserID          string `json:"user_id"`
	TransactionID   int64  `json:"transaction_id"`
	TransactionHash string `json:"transaction_hash"`
}

type PaymentStatusChangeResponse struct {
	Result  model.Result              `json:"result"`
	Payload *TransactionStatusPayload `json:"payload"`
	Error   *model.ServerError        `json:"error"`
}

//////////////////////////
//PaymentStatusCheck Request
//////////////////////////

type PaymentStatusCheckRequest struct {
	TransactionID int64 `json:"transaction_id"`
}

type PaymentStatusCheckResponse struct {
	Result  model.Result              `json:"result"`
	Payload *TransactionStatusPayload `json:"payload"`
	Error   *model.ServerError        `json:"error"`
}

type TransactionStatusPayload struct {
	TransactionStatus string `json:"transaction_status"`
}

//////////////////////////
//PaymentGetFromID Request
//////////////////////////

type PaymentGetFromIDRequest struct {
	UserID string `json:"user_id"`
}

type PaymentGetFromIDResponse struct {
	Result  model.Result       `json:"result"`
	Payload *Payments          `json:"payload"`
	Error   *model.ServerError `json:"error"`
}

//////////////////////////
//PaymentGetFromEmail Request
//////////////////////////

type PaymentGetFromEmailRequest struct {
	Email string `json:"email"`
}

type PaymentGetFromEmailResponse struct {
	Result  model.Result       `json:"result"`
	Payload *Payments          `json:"payload"`
	Error   *model.ServerError `json:"error"`
}

type Payments struct {
	Payments []*transaction.Transaction
}

//////////////////////////
//PaymentCancel Request
//////////////////////////

type PaymentCancelRequest struct {
	TransactionID int64 `json:"transaction_id"`
}

type PaymentCancelResponse struct {
	Result  model.Result              `json:"result"`
	Payload *TransactionStatusPayload `json:"payload"`
	Error   *model.ServerError        `json:"error"`
}
