package repository

import "RestGo/pkg/domain/entity"

type TransactionRepository interface {
	WriteTransaction(transaction entity.Transaction) error
	MoveBalance(movement entity.BalanceMovement) (int, error)
}
