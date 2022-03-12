package container

import (
	"RestGo/pkg/adapter/client/api"
	"RestGo/pkg/usecase/customer"
	"github.com/sarulabs/di"
)

func CustomerContainer(ctn di.Container) (interface{}, error) {
	request, _ := ctn.SubContainer()
	defer request.Delete()

	cust := api.NewCustomerAPI()

	return customer.NewCustomerInteractor(cust), nil
}
