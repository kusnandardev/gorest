package router

import (
	"RestGo/pkg/adapter/rest/handler"
	"RestGo/pkg/shared/di"
	"RestGo/pkg/shared/enum"
	"RestGo/pkg/usecase/customer"
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

	r.POST("/customer/login", handler.NewCustomerHandler(ctn.Resolve(enum.CustomerContainer).(*customer.Interactor)).Login)

	r.POST("/customer/logout")

	return r
}
