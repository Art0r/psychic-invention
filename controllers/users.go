package controllers

import (
	"net/http"

	"github.com/Art0r/psychic-invention/models"
	"github.com/gin-gonic/gin"
)

func GetUsers(ctx *gin.Context) {
	userModel := ctx.MustGet("userModel").(*models.UserModel)

	users, err := userModel.GetAllUsers()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, GetUsersResponse(err.Error(), nil))
		return
	}

	ctx.JSON(http.StatusOK, GetUsersResponse(nil, users))
}

func GetUserById(ctx *gin.Context) {
	userModel := ctx.MustGet("userModel").(*models.UserModel)

	id, idDefined := ctx.Params.Get("id")

	if !idDefined {
		ctx.JSON(http.StatusNotFound, GetUserResponse("ID não foi definido", nil))
		return
	}

	user, err := userModel.GetUserById(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, GetUserResponse(err.Error(), nil))
		return
	}

	if user == nil {
		ctx.JSON(http.StatusNotFound, GetUserResponse("Usuário não encontrado", nil))
		return
	}

	ctx.JSON(http.StatusOK, GetUserResponse(nil, user))
}

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
