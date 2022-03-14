package repository

import "RestGo/pkg/domain/entity"

type CustomerRepository interface {
	GetByUsername(username string) (entity.Customer, error)
	IsCustomerExist(userid string) bool
}
