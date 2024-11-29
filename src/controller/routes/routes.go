package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/makotoiguchi/go-crud/src/controller"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/userById/:user_id", controller.FindUserById)
	r.GET("/userByEmail/:user_email", controller.FindUserByEmail)
	r.POST("/createUser", controller.CreateUser)
	r.PUT("/updateUser/:user_id", controller.UpdateUser)
	r.DELETE("/deleteUser/:user_id", controller.DeleteUser)
}
