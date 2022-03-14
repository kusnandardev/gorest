package api

import (
	"RestGo/pkg/domain/entity"
	"RestGo/pkg/shared/enum"
	"fmt"
	"github.com/patrickmn/go-cache"
	"github.com/sarulabs/di"
	"os"
)

type TransactionAPI struct {
	cache *cache.Cache
}

func NewTransactionAPI(ctn di.Container) *TransactionAPI {
	cc := ctn.Get(string(enum.CacheContainer)).(*cache.Cache)
	return &TransactionAPI{cache: cc}
}

func (t TransactionAPI) WriteTransaction(data entity.Transaction) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	f, err := os.OpenFile(wd+"/trx_log/transaction.log", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(fmt.Sprintf("%d | %s | %s | %s | %d | %d \n", data.Id, data.TransactionType, data.SourceId, data.DestinationId, data.Amount, data.TransactionDate))
	if err != nil {
		return err
	}
	return nil
}

func (t TransactionAPI) MoveBalance(data entity.BalanceMovement) (int, error) {
	wd, err := os.Getwd()
	if err != nil {
		return 0, err
	}
	f, err := os.OpenFile(wd+"/trx_log/movement.log", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	_, err = f.WriteString(fmt.Sprintf("%d | %s | %d | %d | %d | %d | %d \n", data.Id, data.UserId, data.TransactionId, data.Drcr, data.Amount, data.BalanceBefore, data.BalanceAfter))
	if err != nil {
		return 0, err
	}
	t.cache.Set("balance_"+data.UserId, data.BalanceAfter, cache.NoExpiration)
	return 0, nil
}
