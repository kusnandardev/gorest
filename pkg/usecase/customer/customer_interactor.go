package customer

import (
	"RestGo/pkg/domain/dto/request"
	"RestGo/pkg/domain/dto/response"
	"RestGo/pkg/domain/repository"
	"RestGo/pkg/shared/util"
	"errors"
	"github.com/mitchellh/mapstructure"
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

func (i *Interactor) Authenticate(authData interface{}) (interface{}, error) {
	var (
		data request.LoginRequestDto
		rsp  response.LoginResponseDto
	)
	err := mapstructure.Decode(authData, &data)
	if err != nil {
		return nil, err
	}
	resp, err := i.customerClient.GetByUsername(data.Username)
	if err != nil {
		return nil, err
	}
	err = util.Compare(resp.Password, data.Password)
	if err != nil {
		return nil, errors.New("Wrong Password")
	}

	mapstructure.Decode(resp, &rsp)

	token, err := util.GenerateToken(resp.Username)
	if err != nil {
		return nil, err
	}
	rsp.Token = token
	i.cache.Set(token, rsp.Username, cache.DefaultExpiration)
	return rsp, nil
}
