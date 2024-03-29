package views

import (
	"github.com/Art0r/psychic-invention/controllers"
	"github.com/gin-gonic/gin"
	"github.com/Art0r/psychic-invention/models"

)

func SetUsersRoutes(r *gin.Engine, userModel *models.UserModel) {
	user := r.Group("/user")
	{
		user.Use(func(ctx *gin.Context) {
			ctx.Set("userModel", userModel)
			ctx.Next()
		})

		user.GET("/", controllers.GetUsers)
		user.GET("/:id", controllers.GetUserById)
	}
}
