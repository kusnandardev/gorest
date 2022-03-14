package handler

import (
	"RestGo/pkg/domain/dto/request"
	"RestGo/pkg/shared/util"
	"RestGo/pkg/usecase/customer"
	"RestGo/pkg/usecase/jwt"
	"github.com/gin-gonic/gin"
)

type CustomerHandler struct {
	customerService customer.InputPort
	jwtService      jwt.InputPort
}

func NewCustomerHandler(uc customer.InputPort, jwt jwt.InputPort) *CustomerHandler {
	return &CustomerHandler{customerService: uc, jwtService: jwt}
}

func (h *CustomerHandler) Login(c *gin.Context) {
	user, password, hasAuth := c.Request.BasicAuth()
	authData := request.LoginRequestDto{
		Username: user,
		Password: password,
	}

	if !hasAuth {
		util.NewResponse(c).BadRequest("use basic auth")
		return
	}

	result, err := h.customerService.Authenticate(authData)
	if err != nil {
		util.NewResponse(c).Unauthorize(err.Error())
		return
	}

	token, err := h.jwtService.GenerateToken(authData.Username)
	if err != nil {
		util.NewResponse(c).InternalServerError(err.Error())
		return
	}

	result.Token = token

	util.NewResponse(c).Ok(result)
	return
}

func (h *CustomerHandler) Logout(c *gin.Context) {

	err := h.customerService.EndSession(util.GetToken(c))
	if err != nil {
		util.NewResponse(c).InternalServerError(err.Error())
		return
	}
	util.NewResponse(c).Ok(gin.H{"message": "success logout"})
	return
}
