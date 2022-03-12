package handler

import (
	"RestGo/pkg/shared/util"
	"RestGo/pkg/usecase/customer"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

type CustomerHandler struct {
	customerService customer.InputPort
}

func NewCustomerHandler(uc customer.InputPort) *CustomerHandler {
	return &CustomerHandler{customerService: uc}
}

func (h *CustomerHandler) Login(c *gin.Context) {
	authData, err := util.NewMapper(c).GetAuthData()
	if err != nil {
		util.NewResponse(c).BadRequest(err.Error())
		return
	}

	validator := validation.Validation{}
	check, err := validator.Valid(authData)
	if err != nil {
		util.NewResponse(c).InternalServerError(err.Error())
		return
	}
	if !check {
		util.NewResponse(c).BadRequest(util.MarkErrors(validator.Errors))
		return
	}

	resp, err := h.customerService.Authenticate(authData)
	if err != nil {
		util.NewResponse(c).InternalServerError(err.Error())
		return
	}
	util.NewResponse(c).Ok(resp)
	return
}
