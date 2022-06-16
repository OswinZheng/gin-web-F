package book_dto

type AddBookDto struct {
	Name   string `form:"name" binding:"required"`
	Author string `form:"author" binding:"required"`
	Num    int    `form:"num" binding:"required"`
}

type ParamIdDto struct {
	Id int `uri:"id" binding:"required"`
}
