package book

import (
	"errors"

	"github.com/OswinZheng/gin-web-F/internal/dto/book_dto"
	"github.com/OswinZheng/gin-web-F/internal/model/book"
)

func AddBook(info *book_dto.AddBookDto) (error, *book.Book) {
	bookModel := &book.Book{
		Name:   info.Name,
		Author: info.Author,
		Num:    info.Num,
	}
	err := bookModel.Add()
	if err != nil {
		return err, nil
	}
	return nil, bookModel
}

func UpdateBook(id int, info *book_dto.AddBookDto) (error, *book.Book) {
	bookModel := book.FindById(id)
	if bookModel.ID == 0 {
		return errors.New("book not exist"), nil
	}
	bookModel.Name = info.Name
	bookModel.Author = info.Author
	bookModel.Num = info.Num
	err := bookModel.Update()
	if err != nil {
		return err, nil
	}
	return nil, bookModel
}

func RemoveBook(id int) (error, *book.Book) {
	bookModel := book.FindById(id)
	if bookModel.ID == 0 {
		return errors.New("book not exist"), nil
	}
	err := bookModel.Remove()
	if err != nil {
		return err, nil
	}
	return nil, bookModel
}

func GetBook(id int) (error, *book.Book) {
	bookModel := book.FindById(id)
	if bookModel.ID == 0 {
		return errors.New("book not exist"), nil
	}
	return nil, bookModel
}
