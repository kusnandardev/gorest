package api

import (
	"RestGo/pkg/domain/entity"
	"encoding/json"
	"errors"
	"github.com/mitchellh/mapstructure"
	"io/ioutil"
	"os"
)

type CustomerAPI struct {
}

func NewCustomerAPI() *CustomerAPI {
	return &CustomerAPI{}
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
		return entity.Customer{}, errors.New("not found")
	}
	return cust, nil
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
