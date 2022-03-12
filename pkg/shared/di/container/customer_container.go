package container

import (
	"RestGo/pkg/usecase/customer"
	"github.com/sarulabs/di"
)

func CustomerContainer(ctn di.Container) (interface{}, error) {
	request, _ := ctn.SubContainer()
	defer request.Delete()

	return customer.NewCustomerInteractor(), nil
}
