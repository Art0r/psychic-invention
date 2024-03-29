package views

import (
	"github.com/Art0r/psychic-invention/controllers"
	"github.com/gin-gonic/gin"
)

func SetUsersRoutes(r *gin.Engine) {
	user := r.Group("/user")
	{
		user.GET("/ping", controllers.Ping)
	}
}
