package container

import (
	"RestGo/pkg/adapter/client/api"
	"RestGo/pkg/usecase/transaction"
	"github.com/sarulabs/di"
)

func TransactionContainer(ctn di.Container) (interface{}, error) {
	request, _ := ctn.SubContainer()
	defer request.Delete()

	cust := api.NewCustomerAPI(request)
	trx := api.NewTransactionAPI(request)

	return transaction.NewTransactionInteractor(trx, cust), nil
}
