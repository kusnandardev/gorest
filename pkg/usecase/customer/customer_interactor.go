package customer

import (
	"RestGo/pkg/domain/dto/request"
	"RestGo/pkg/domain/repository"
	"github.com/mitchellh/mapstructure"
)

type Interactor struct {
	customerClient repository.CustomerRepository
}

func NewCustomerInteractor(cust repository.CustomerRepository) *Interactor {
	return &Interactor{
		customerClient: cust,
	}
}

func (i *Interactor) Authenticate(authData interface{}) (interface{}, error) {
	var data request.LoginRequestDto
	err := mapstructure.Decode(authData, &data)
	if err != nil {
		return nil, err
	}
	resp, err := i.customerClient.GetByUsername(data.Username)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
