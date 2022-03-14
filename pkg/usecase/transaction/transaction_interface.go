package transaction

import (
	"RestGo/pkg/domain/dto/request"
	"RestGo/pkg/domain/dto/response"
)

type InputPort interface {
	Transfer(transferData request.TransferRequestDto) (response.TransferResponseDto, error)
}
