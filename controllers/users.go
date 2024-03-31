package controllers

import (
	"net/http"

	"github.com/Art0r/psychic-invention/models"
	"github.com/Art0r/psychic-invention/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

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

func getUserVerifyId(ctx *gin.Context) (*models.User, *models.UserModel) {
	userModel := ctx.MustGet("userModel").(*models.UserModel)
	id, idDefined := ctx.Params.Get("id")

	if !idDefined {
		ctx.JSON(http.StatusNotFound, GetUserResponse("ID não foi definido", nil))
		return nil, userModel
	}

	if !utils.IsUUID(id) {
		ctx.JSON(http.StatusNotFound, GetUserResponse("Usuário não encontrado", nil))
		return nil, userModel
	}

	user, err := userModel.GetUserById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, GetUserResponse("Usuário não encontrado", nil))
		return nil, userModel
	}

	if user == nil {
		ctx.JSON(http.StatusNotFound, GetUserResponse("Usuário não encontrado", nil))
		return nil, userModel
	}

	return user, userModel
}

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
	user, _ := getUserVerifyId(ctx)
	if user == nil {
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

	if len(newUser.Email) == 0 || len(newUser.Name) == 0 {
		ctx.JSON(http.StatusBadRequest, GetUserResponse("Algum campo necessário não foi preenchido", nil))
		return	
	}

	if !utils.IsValidEmail(newUser.Email) {
		ctx.JSON(http.StatusBadRequest, GetUserResponse("Endereço de Email Inválido", nil))
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
	var userUpdates *models.User

	userRetrieved, userModel := getUserVerifyId(ctx)
	if userRetrieved == nil { return }

	if err := ctx.ShouldBindJSON(&userUpdates); err != nil {
		ctx.JSON(http.StatusBadRequest, GetUserResponse(err, nil))
		return
	}

	err := whichFieldUpdate(userUpdates, userRetrieved, userModel)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, GetUserResponse(err.Error(), nil))
		return
	}

	if userRetrieved, err = userModel.GetUserById(userRetrieved.ID); err != nil {
		ctx.JSON(http.StatusNotFound, GetUserResponse(err, nil))
		return
	}

	ctx.JSON(http.StatusOK, GetUserResponse(nil, userRetrieved))
}

func DeleteUser(ctx *gin.Context) {
	user, userModel := getUserVerifyId(ctx)
	if user == nil {
		return
	}

	err := userModel.DeleteUser(user.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, DeleteResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, DeleteResponse(nil))
}
