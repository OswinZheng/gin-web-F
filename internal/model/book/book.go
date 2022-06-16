package book

import (
	"github.com/OswinZheng/gin-web-F/internal/repository/postgres"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name   string `json:"name"`
	Author string `json:"author"`
	Num    int    `json:"num"`
}

func (m *Book) TableName() string {
	return "book"
}

func (m *Book) Add() error {
	return postgres.Db.Create(m).Error
}

func (m *Book) Update() error {
	return postgres.Db.Save(m).Error
}

func (m *Book) Remove() error {
	return postgres.Db.Delete(m).Error
}

func FindById(id int) *Book {
	var info Book
	postgres.Db.Where("id = ?", id).First(&info)
	return &info
}
