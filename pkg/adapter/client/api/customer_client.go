package api

import (
	"RestGo/pkg/domain/entity"
	"RestGo/pkg/shared/enum"
	"encoding/json"
	"errors"
	"github.com/mitchellh/mapstructure"
	"github.com/patrickmn/go-cache"
	"github.com/sarulabs/di"
	"io/ioutil"
	"os"
)

type CustomerAPI struct {
	cache *cache.Cache
}

func NewCustomerAPI(ctn di.Container) *CustomerAPI {
	cc := ctn.Get(string(enum.CacheContainer)).(*cache.Cache)
	return &CustomerAPI{cache: cc}
}

func (c *CustomerAPI) GetByUsername(username string) (entity.Customer, error) {
	var (
		cust  entity.Customer
		found = false
	)
	customers, err := c.getAll()
	if err != nil {
		return entity.Customer{}, err
	}
	for i := range customers {
		if customers[i].Username == username {
			cust = customers[i]
			found = true
		}
	}
	if !found {
		return entity.Customer{}, errors.New("username not found")
	}
	balance, found := c.cache.Get("balance_" + username)
	if found {
		cust.Balance = balance.(int)
	}

	return cust, nil
}

func (c *CustomerAPI) IsCustomerExist(userid string) bool {
	customers, err := c.getAll()
	if err != nil {
		return false
	}
	for i := range customers {
		if customers[i].Username == userid {
			return true
		}
	}
	return false
}

func (c *CustomerAPI) getAll() ([]entity.Customer, error) {
	var (
		resp []map[string]interface{}
		cust []entity.Customer
	)
	wd, _ := os.Getwd()
	jsonFile, err := os.Open(wd + "/mock_data/user1.json")
	if err != nil {
		return []entity.Customer{}, err
	}
	defer jsonFile.Close()
	bCust, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(bCust, &resp)
	err = mapstructure.Decode(resp, &cust)
	return cust, nil
}
