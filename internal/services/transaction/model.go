package transaction

import "payment-emulator/internal/repository/transaction"

type TransactionService struct {
	rep transaction.Transactioner
}

func InitTransactionService(tr transaction.Transactioner) *TransactionService {
	service := &TransactionService{
		rep: tr,
	}

	return service
}
