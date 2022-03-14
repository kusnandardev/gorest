package handler

import (
	"RestGo/pkg/domain/dto/request"
	"RestGo/pkg/shared/util"
	"RestGo/pkg/usecase/transaction"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
)

type TransactionHandler struct {
	transactionService transaction.InputPort
}

func NewTransactionHandler(uc transaction.InputPort) *TransactionHandler {
	return &TransactionHandler{transactionService: uc}
}

func (t *TransactionHandler) Transfer(c *gin.Context) {
	var (
		claim map[string]interface{}
		req   = request.TransferRequestDto{}
	)
	err := mapstructure.Decode(c.MustGet("claims"), &claim)

	if err != nil {
		util.NewResponse(c).InternalServerError(err.Error())
		return
	}
	c.ShouldBind(&req)
	req.SourceId = claim["jti"].(string)

	res, err := t.transactionService.Transfer(req)
	if err != nil {
		util.NewResponse(c).InternalServerError(err.Error())
		return
	}
	util.NewResponse(c).Ok(res)
	return
}
