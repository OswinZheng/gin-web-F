package validator

import (
	"github.com/go-playground/validator/v10"
)

type Messages map[string]string

// 验证器接口

type Validator interface {
	GetMessage() Messages
}

// 自定义验证错误信息 验证结构体需要实现 GetMessage() 方法

func GetErrorMsg(request Validator, err error) string {
	for _, v := range err.(validator.ValidationErrors) {
		if message, exist := request.GetMessage()[v.Field()+"."+v.Tag()]; exist {
			return message
		}
		return v.Error()
	}

	return "Parameter validation failed"
}
