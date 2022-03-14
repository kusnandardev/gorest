package customer

import (
	"RestGo/pkg/domain/dto/request"
	"RestGo/pkg/domain/dto/response"
	"RestGo/pkg/domain/repository"
	"RestGo/pkg/shared/util"
	"errors"
	"github.com/patrickmn/go-cache"
	"time"
)

type Interactor struct {
	customerClient repository.CustomerRepository
	cache          *cache.Cache
}

func NewCustomerInteractor(cust repository.CustomerRepository) *Interactor {
	c := cache.New(5*time.Minute, 10*time.Minute)
	return &Interactor{
		customerClient: cust,
		cache:          c,
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
