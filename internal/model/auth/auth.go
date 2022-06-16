package auth

import (
	"github.com/OswinZheng/gin-web-F/internal/repository/postgres"
	"github.com/OswinZheng/gin-web-F/pkg/util"
	"gorm.io/gorm"
)

type Auth struct {
	gorm.Model
	ID       int    `gorm:"primary_key" json:"id"`
	UserName string `json:"username" gorm:"column:username"`
	Password string `json:"password"`
}

func (m *Auth) TableName() string {
	return "user"
}

func (m *Auth) Add() error {
	m.Password = util.EncodeMD5(m.Password)
	return postgres.Db.Create(m).Error
}

func Get(username, password string) *Auth {
	var info Auth
	postgres.Db.Where("username = ? and password = ?", username, util.EncodeMD5(password)).First(&info)
	return &info
}
