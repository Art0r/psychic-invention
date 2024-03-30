package controllers

import (
	"net/http"

	"github.com/Art0r/psychic-invention/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

func CreateUser(ctx *gin.Context) {
	userModel := ctx.MustGet("userModel").(*models.UserModel)

	var newUser *models.User
	var id string

	for {
		id = uuid.NewString()

		user, _ := userModel.GetUserById(id)
		if user == nil {
			break
		}
	}

	if err := ctx.ShouldBindJSON(&newUser); err != nil {
		ctx.JSON(http.StatusBadRequest, GetUserResponse(err, nil))
		return
	}

	user, _ := userModel.GetUserByEmail(newUser.Email)
	if user != nil {
		ctx.JSON(http.StatusConflict, GetUserResponse("Um usuário com esse email já existe", nil))
		return
	}

	newUser.ID = id
	err := userModel.CreateUser(newUser)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, GetUserResponse(err.Error(), nil))
		return
	}

	ctx.JSON(http.StatusOK, GetUserResponse(nil, newUser))
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
