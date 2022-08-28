package utils

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"strings"
)

// 校验数据
func Validate(struc interface{}) error {
	v := validator.New()
	err := v.Struct(struc)
	if err != nil {
		msg := "Invalid parameter: "
		es := err.(validator.ValidationErrors)
		errFields := []string{}
		for _, e := range es {
			errFields = append(errFields, fmt.Sprintf("%v(%v)", e.StructField(), e.ActualTag()))
		}
		// 拼接不符合要求的字段
		msg += strings.Join(errFields, ",")
		return errors.New(msg)
	}
	return nil
}
