package util

import (
	"RestGo/pkg/domain/dto/request"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

type Mapper struct {
	Context *gin.Context
}

func NewMapper(c *gin.Context) *Mapper {
	return &Mapper{Context: c}
}

func (m *Mapper) GetAuthData() (interface{}, error) {
	authKey := strings.Split(m.Context.Request.Header.Get("Authorization"), " ")
	if len(authKey) != 2 {
		return nil, errors.New("basic auth can't be empty")
	}
	if authKey[0] != "Basic" {
		return errors.New("use basic auth as authentication method"), nil
	}
	data, _ := base64.StdEncoding.DecodeString(authKey[1])
	decodedData := fmt.Sprintf("%q", data)
	splitedData := strings.Split(decodedData[1:len(decodedData)-1], ":")
	return &request.LoginRequestDto{
		Username: splitedData[0],
		Password: splitedData[1],
	}, nil
}
