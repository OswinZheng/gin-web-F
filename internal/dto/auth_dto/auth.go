package auth_dto

import "github.com/OswinZheng/gin-web-F/pkg/validator"

type AddAuthDto struct {
	UserName string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type ParamIdDto struct {
	Id int `uri:"id" binding:"required"`
}

type RspAuth struct {
	Id       int    `json:"id"`
	UserName string `json:"username"`
}

func (authDto AddAuthDto) GetMessage() validator.Messages {
	return validator.Messages{
		"UserName.required": "用户名不能为空",
		"Password.required": "密码不能为空",
	}
}
