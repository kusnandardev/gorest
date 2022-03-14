package transaction

import (
	"RestGo/pkg/domain/dto/request"
	"RestGo/pkg/domain/dto/response"
	"RestGo/pkg/domain/entity"
	"RestGo/pkg/domain/repository"
	"errors"
	"math/rand"
	"time"
)

type Interactor struct {
	transactionClient repository.TransactionRepository
	customerClient    repository.CustomerRepository
}

func NewTransactionInteractor(trx repository.TransactionRepository, cust repository.CustomerRepository) *Interactor {
	return &Interactor{
		transactionClient: trx,
		customerClient:    cust,
	}
}

func (i *Interactor) Transfer(data request.TransferRequestDto) (response.TransferResponseDto, error) {

	// check is customer exist and get their data
	srcCust, err := i.customerClient.GetByUsername(data.SourceId)
	if err != nil {
		return response.TransferResponseDto{}, errors.New("source id is invalid")
	}
	dstCust, err := i.customerClient.GetByUsername(data.DestinationId)
	if err != nil {
		return response.TransferResponseDto{}, errors.New("source id is invalid")
	}

	// check is cust balance more than transfer amount
	//if srcCust.Balance < trfData.Amount {
	//	return nil, errors.New("insufficient balance")
	//}

	// write transaction and balance movement

	trx := entity.Transaction{
		Id:              rand.Int(),
		TransactionType: "trf",
		SourceId:        data.SourceId,
		DestinationId:   data.DestinationId,
		Amount:          data.Amount,
		TransactionDate: time.Now().Unix(),
	}

	err = i.transactionClient.WriteTransaction(trx)
	if err != nil {
		return response.TransferResponseDto{}, err
	}

	srcMov := entity.BalanceMovement{
		Id:            rand.Int(),
		UserId:        srcCust.Username,
		TransactionId: trx.Id,
		Drcr:          1,
		Amount:        data.Amount,
		BalanceBefore: srcCust.Balance,
		BalanceAfter:  (srcCust.Balance - data.Amount),
	}

	_, err = i.transactionClient.MoveBalance(srcMov)

	dstMov := entity.BalanceMovement{
		Id:            rand.Int(),
		UserId:        dstCust.Username,
		TransactionId: trx.Id,
		Drcr:          2,
		Amount:        data.Amount,
		BalanceBefore: dstCust.Balance,
		BalanceAfter:  (dstCust.Balance + data.Amount),
	}

	_, err = i.transactionClient.MoveBalance(dstMov)

	return response.TransferResponseDto{
		TransactionId:   trx.Id,
		DestinationId:   dstCust.Username,
		DestinationName: dstCust.Name,
		Amount:          data.Amount,
		Balance:         srcMov.BalanceAfter,
	}, nil
}
