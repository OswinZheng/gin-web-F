package auth

import (
	"errors"

	"github.com/OswinZheng/gin-web-F/internal/dto/auth_dto"
	"github.com/OswinZheng/gin-web-F/internal/model/auth"
	"github.com/OswinZheng/gin-web-F/pkg/util"
)

func Register(info *auth_dto.AddAuthDto) (error, *auth_dto.RspAuth) {
	authModel := &auth.Auth{
		UserName: info.UserName,
		Password: info.Password,
	}
	err := authModel.Add()
	if err != nil {
		return err, nil
	}
	return nil, &auth_dto.RspAuth{
		Id:       authModel.ID,
		UserName: authModel.UserName,
	}
}

func Login(username, password string) (error, string) {
	authModel := auth.Get(username, password)
	if authModel.ID == 0 {
		return errors.New("username error"), ""
	}

	token, err := util.GenerateToken(username, password)
	if err != nil {
		return err, ""
	}

	return nil, token
}
