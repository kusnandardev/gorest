package util

import (
	"fmt"
	"github.com/astaxie/beego/validation"
)

func MarkErrors(errors []*validation.Error) string {
	res := ""
	for _, err := range errors {
		res += fmt.Sprintf("%s %s, \n", err.Key, err.Message)
	}

	return res
}
