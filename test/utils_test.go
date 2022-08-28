package test

import (
	"github.com/DolphinDong/gotools/utils"
	"log"
	"testing"
)

func TestValidate(t *testing.T) {
	type User struct {
		Name string `validate:"required,min=5"`
	}
	u := &User{
		Name: "123",
	}
	err := utils.Validate(u)
	if err != nil {
		log.Fatalf("%+v", err)
	}

}
