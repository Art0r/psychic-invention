package controllers

import (
	"github.com/Art0r/psychic-invention/models"
	"github.com/gin-gonic/gin"
)

func GetUserResponse(err any, user *models.User) map[string]any {
	return gin.H{
		"error": err,
		"user":  user,
	}
}

func GetUsersResponse(err any, users []*models.User) map[string]any {
	return gin.H{
		"error": err,
		"users": users,
	}
}
