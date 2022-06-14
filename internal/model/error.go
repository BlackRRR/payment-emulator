package model

const (
	InternalTransactionServiceError = "transaction_service_error"
)

type Error string

func (e Error) Error() string {
	return string(e)
}

type ServerError struct {
	Code    string
	Message string
}

func NewTransactionError(err error) *ServerError {
	return &ServerError{
		Code:    InternalTransactionServiceError,
		Message: err.Error(),
	}
}
