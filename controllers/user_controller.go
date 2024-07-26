package controllers

import (
	"main/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) *UserController {
	return &UserController{userService}
}

func (c *UserController) FindAllUsers(ctx *gin.Context) {
	users, err := c.userService.FindAllUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func (c *UserController) FindUserByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	user, err := c.userService.FindUserByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}
	ctx.JSON(http.StatusOK, user)
}
