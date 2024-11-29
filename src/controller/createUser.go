package controller

import (
	"github.com/makotoiguchi/go-crud/src/config/rest_err"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	err := rest_err.NewBadRequestErr("This is a bad request")
	c.JSON(err.Code, err)
}
