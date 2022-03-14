package customer

import (
	"RestGo/pkg/domain/dto/request"
	"RestGo/pkg/domain/dto/response"
)

type InputPort interface {
	Authenticate(data request.LoginRequestDto) (response.LoginResponseDto, error)
	EndSession(token string) error
}
