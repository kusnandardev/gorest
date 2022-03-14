package container

import (
	"RestGo/pkg/usecase/jwt"
	"github.com/sarulabs/di"
)

func JwtContainer(ctn di.Container) (interface{}, error) {
	request, _ := ctn.SubContainer()
	defer request.Delete()
	return jwt.NewJWTInteractor(), nil
}
