package container

import (
	"RestGo/pkg/adapter/db/inmemory"
	"RestGo/pkg/usecase/jwt"
	"github.com/sarulabs/di"
)

func JwtContainer(ctn di.Container) (interface{}, error) {
	request, _ := ctn.SubContainer()
	defer request.Delete()
	cc := inmemory.DefaultCacheHandler()
	return jwt.NewJWTInteractor(cc), nil
}
