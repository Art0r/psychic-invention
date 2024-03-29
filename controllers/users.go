package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

// user, _ := userModel.GetUserById("1")
// fmt.Println(user)

// fmt.Println("-------------------------")

// users, _ := userModel.GetAllUsers()
// fmt.Println(users)

// fmt.Println("-------------------------")

// userModel.UpdateUserEmail("2", "asf@asf.com")
// user, _ = userModel.GetUserById("2")
// fmt.Println(user)

// fmt.Println("-------------------------")

// userModel.UpdateUserName("2", "Asf")
// user, _ = userModel.GetUserById("2")
// fmt.Println(user)

// fmt.Println("-------------------------")

// userModel.DeleteUser("2")
// users, _ = userModel.GetAllUsers()
// fmt.Println(users)

// fmt.Println("-------------------------")

// user, _ = userModel.GetUserByName("Art0r")
// fmt.Println(user)

// fmt.Println("-------------------------")

// user, _ = userModel.GetUserByEmail("art0r@art0r.com")
// fmt.Println(user
