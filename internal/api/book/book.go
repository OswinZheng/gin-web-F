package book

import (
	"net/http"

	"github.com/OswinZheng/gin-web-F/internal/dto/book_dto"
	"github.com/OswinZheng/gin-web-F/internal/services/book"
	"github.com/OswinZheng/gin-web-F/pkg/response"
	"github.com/gin-gonic/gin"
)

func AddBook(c *gin.Context) {
	param := book_dto.AddBookDto{}
	if err := c.ShouldBind(&param); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	err, bookInfo := book.AddBook(&param)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}
	c.JSON(http.StatusOK, bookInfo)
}

func UpdateBook(c *gin.Context) {
	idParam := book_dto.ParamIdDto{}
	if err := c.ShouldBindUri(idParam); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	param := book_dto.AddBookDto{}
	if err := c.ShouldBind(&param); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	err, bookInfo := book.UpdateBook(idParam.Id, &param)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}
	c.JSON(http.StatusOK, bookInfo)
}

func RemoveBook(c *gin.Context) {
	idParam := book_dto.ParamIdDto{}
	if err := c.ShouldBindUri(idParam); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	err, bookInfo := book.RemoveBook(idParam.Id)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}
	c.JSON(http.StatusOK, bookInfo)
}

func GetBook(c *gin.Context) {
	idParam := book_dto.ParamIdDto{}
	if err := c.ShouldBindUri(&idParam); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	err, bookInfo := book.GetBook(idParam.Id)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}
	c.JSON(http.StatusOK, response.Success.WithData(bookInfo))
}
