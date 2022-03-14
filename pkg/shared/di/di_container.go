package di

import (
	"RestGo/pkg/shared/di/container"
	"RestGo/pkg/shared/enum"
	"github.com/sarulabs/di"
)

type Container struct {
	ctn di.Container
}

func NewContainer() *Container {
	builder, _ := di.NewBuilder()

	_ = builder.Add([]di.Def{
		{Name: string(enum.CustomerContainer), Build: container.CustomerContainer},
		{Name: string(enum.CacheContainer), Scope: di.Request, Build: container.MainCacheContainer},
		{Name: string(enum.JWTContainer), Build: container.JwtContainer},
		{Name: string(enum.TransactionContainer), Build: container.TransactionContainer},
	}...)

	return &Container{
		ctn: builder.Build(),
	}
}

func (c *Container) Resolve(name enum.ContainerName) interface{} {
	return c.ctn.Get(string(name))
}
