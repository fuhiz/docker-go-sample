package controller

import (
	"github.com/gin-gonic/gin"
)

func Setup(r *gin.RouterGroup) {
	users := r.Group("/users")
	{
		u := UserController{}
		users.GET("", u.Index)
		users.GET("/:id", u.GetUser)
		users.POST("", u.CreateUser)
		users.PATCH("/:id", u.UpdateUser)
		users.DELETE("/:id", u.DeleteUser)
	}
}
