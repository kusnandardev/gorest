package router

import (
	"RestGo/pkg/adapter/rest/handler"
	"RestGo/pkg/infrastucture/rest/middleware"
	"RestGo/pkg/shared/di"
	"RestGo/pkg/shared/enum"
	"RestGo/pkg/usecase/customer"
	"RestGo/pkg/usecase/jwt"
	"RestGo/pkg/usecase/transaction"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	ctn := di.NewContainer()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.Default())

	gin.SetMode("debug")

	// health check
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "ok"})
	})

	r.POST("/customer/login", handler.NewCustomerHandler(ctn.Resolve(enum.CustomerContainer).(*customer.Interactor), ctn.Resolve(enum.JWTContainer).(*jwt.Interactor)).Login)

	r.POST("/transaction/transfer", middleware.AuthorizeJWT(), handler.NewTransactionHandler(ctn.Resolve(enum.TransactionContainer).(*transaction.Interactor)).Transfer)

	r.POST("/customer/logout", middleware.AuthorizeJWT(), handler.NewCustomerHandler(ctn.Resolve(enum.CustomerContainer).(*customer.Interactor), ctn.Resolve(enum.JWTContainer).(*jwt.Interactor)).Logout)

	return r
}
