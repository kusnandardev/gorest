package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type (
	Response struct {
		Context *gin.Context
	}

	ResponseData struct {
		Message string      `json:"message,omitempty"`
		Content interface{} `json:"content,omitempty"`
	}
)

func NewResponse(ctx *gin.Context) *Response {
	return &Response{Context: ctx}
}

func (r *Response) Created(content interface{}) {
	r.Context.JSON(http.StatusCreated, content)
	return
}

func (r *Response) BadRequest(msg string) {
	r.Context.JSON(http.StatusBadRequest, ResponseData{
		Message: msg,
	})
	return
}

func (r *Response) InternalServerError(msg string) {
	r.Context.JSON(http.StatusInternalServerError, ResponseData{
		Message: msg,
	})
	return
}
