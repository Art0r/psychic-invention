package controllers

import (
	"net/http"

	"github.com/Art0r/psychic-invention/models"
	"github.com/Art0r/psychic-invention/utils"
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

	if !utils.IsUUID(id) {
		ctx.JSON(http.StatusNotFound, GetUserResponse("Usuário não encontrado", nil))
		return
	}

	user, err := userModel.GetUserById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, GetUserResponse("Usuário não encontrado", nil))
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

func UpdateUser(ctx *gin.Context) {
	userModel := ctx.MustGet("userModel").(*models.UserModel)
	var userUpdates *models.User

	id, idDefined := ctx.Params.Get("id")
	if !idDefined {
		ctx.JSON(http.StatusNotFound, GetUserResponse("ID não foi definido", nil))
		return
	}

	if err := ctx.ShouldBindJSON(&userUpdates); err != nil {
		ctx.JSON(http.StatusBadRequest, GetUserResponse(err, nil))
		return
	}

	userRetrieved, err := userModel.GetUserById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, GetUserResponse(err, nil))
		return
	}

	err = whichFieldUpdate(userUpdates, userRetrieved, userModel)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, GetUserResponse(err.Error(), nil))
		return
	}

	if userRetrieved, err = userModel.GetUserById(id); err != nil {
		ctx.JSON(http.StatusNotFound, GetUserResponse(err, nil))
		return
	}

	ctx.JSON(http.StatusOK, GetUserResponse(nil, userRetrieved))
}

func DeleteUser(ctx *gin.Context) {
	userModel := ctx.MustGet("userModel").(*models.UserModel)

	id, idDefined := ctx.Params.Get("id")

	if !idDefined {
		ctx.JSON(http.StatusNotFound, DeleteResponse("ID não foi definido"))
		return
	}

	if !utils.IsUUID(id) {
		ctx.JSON(http.StatusNotFound, DeleteResponse("Usuário não encontrado"))
		return
	}

	user, err := userModel.GetUserById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, GetUserResponse("Usuário não encontrado", nil))
		return
	}
	if user == nil {
		ctx.JSON(http.StatusNotFound, DeleteResponse("Usuário não encontrado"))
		return
	}

	err = userModel.DeleteUser(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, DeleteResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, DeleteResponse(nil))
}

func whichFieldUpdate(userUpdates, userRetrieved *models.User, userModel *models.UserModel) error {
	columns := map[string]string{}

	if len(userUpdates.Email) != 0 {
		columns["email"] = userUpdates.Email
	}

	if len(userUpdates.Name) != 0 {
		columns["name"] = userUpdates.Name
	}

	err := userModel.UpdateUser(userRetrieved.ID, columns)
	if err != nil {
		return err
	}
	return nil
}

// userModel.DeleteUser("2")
// users, _ = userModel.GetAllUsers()
// fmt.Println(users)
