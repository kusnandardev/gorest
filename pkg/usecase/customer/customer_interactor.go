package customer

import (
	"RestGo/pkg/domain/dto/request"
	"RestGo/pkg/domain/dto/response"
	"RestGo/pkg/domain/repository"
	"RestGo/pkg/shared/util"
	"errors"
	"fmt"
)

type Interactor struct {
	customerClient repository.CustomerRepository
	cache          repository.Inmemory
}

func NewCustomerInteractor(cust repository.CustomerRepository, cc repository.Inmemory) *Interactor {
	return &Interactor{
		customerClient: cust,
		cache:          cc,
	}
}

func (i *Interactor) Authenticate(data request.LoginRequestDto) (response.LoginResponseDto, error) {
	resp, err := i.customerClient.GetByUsername(data.Username)
	if err != nil {
		return response.LoginResponseDto{}, err
	}
	err = util.Compare(resp.Password, data.Password)
	if err != nil {
		return response.LoginResponseDto{}, errors.New("Wrong Password")
	}
	return response.LoginResponseDto{
		Username: resp.Username,
		Name:     resp.Name,
	}, nil
}

func (i *Interactor) EndSession(token string) error {
	err := i.cache.Delete(token)
	if err != nil {
		return err
	}
	value, err := i.cache.Get(token)
	fmt.Println(value)
	return nil
}
