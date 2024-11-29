package controller

import "github.com/gin-gonic/gin"

func FindUserById(c *gin.Context) {
	c.Writer.Write([]byte("FindUserById"))
}

func FindUserByEmail(c *gin.Context) {
}
