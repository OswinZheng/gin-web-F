package auth

import (
	"net/http"

	"github.com/OswinZheng/gin-web-F/internal/dto/auth_dto"
	"github.com/OswinZheng/gin-web-F/internal/services/auth"
	"github.com/OswinZheng/gin-web-F/pkg/response"
	"github.com/OswinZheng/gin-web-F/pkg/validator"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	param := auth_dto.AddAuthDto{}
	if err := c.ShouldBind(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": validator.GetErrorMsg(param, err)})
		return
	}
	err, authInfo := auth.Register(&param)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, response.Success.WithData(authInfo))
}

func Login(c *gin.Context) {
	username, password := c.PostForm("username"), c.PostForm("password")
	err, token := auth.Login(username, password)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}
	c.JSON(http.StatusOK, response.Success.WithData(token))
}
