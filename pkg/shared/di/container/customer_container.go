package container

import (
	"RestGo/pkg/adapter/client/api"
	"RestGo/pkg/adapter/db/inmemory"
	"RestGo/pkg/usecase/customer"
	"github.com/sarulabs/di"
)

func CustomerContainer(ctn di.Container) (interface{}, error) {
	request, _ := ctn.SubContainer()
	defer request.Delete()

	cust := api.NewCustomerAPI(request)
	cc := inmemory.DefaultCacheHandler()

	return customer.NewCustomerInteractor(cust, cc), nil
}
