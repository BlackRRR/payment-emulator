package services

import (
	"payment-emulator/internal/repository"
	"payment-emulator/internal/services/transaction"
)

type Services struct {
	Trans *transaction.TransactionService
}

func InitAllServices(repository *repository.Repositories) *Services {
	trans := transaction.InitTransactionService(repository.Trans)

	services := &Services{
		Trans: trans,
	}

	return services
}
