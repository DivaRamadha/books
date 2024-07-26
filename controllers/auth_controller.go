package controllers

import (
	"main/crypt"
	"main/models"
	"main/repositories"
	"main/services"
	"main/utils"
	"net/http"
	"unicode"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthController struct {
}

func NewAuthController(db *gorm.DB) *AuthController {
	return &AuthController{}
}

func (c *AuthController) Login(ctx *gin.Context) {
	var input models.Auth

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userService := services.NewUserService(repositories.NewUserRepository(ctx.MustGet("database").(*gorm.DB)))
	user, err := userService.FindUserByUserName(input.Username)

	pass, _ := crypt.Decrypt(user.Password)
	if err != nil || pass != input.Password {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}
	token, err := utils.GenerateToken(user.ID, user.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

func (c *AuthController) Register(ctx *gin.Context) {
	var input models.User

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Username == "" || input.Password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Username and password are required"})
		return
	}

	if len(input.Username) < 6 || len(input.Username) > 20 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Username must be between 6 and 20 characters"})
		return
	}

	if len(input.Password) < 8 || !(isUppercaseCharacter(input.Password) && isLowercaseCharacter(input.Password) && isSpecialCharacter(input.Password) && isNumber(input.Password)) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Password must be more than 8 characters, including an uppercase letter, a lowercase letter, a special character, and a number"})
		return
	}

	userService := services.NewUserService(repositories.NewUserRepository(ctx.MustGet("database").(*gorm.DB)))
	user, err := userService.FindUserByUserName(input.Username)
	if err == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
		return
	}

	if user != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
		return
	}

	hasedPassword, err := crypt.EncryptPass(input.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	input.Password = string(hasedPassword)
	if err := userService.CreateUser(&input); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, input)
}

func isUppercaseCharacter(password string) bool {
	for _, char := range password {
		if unicode.IsUpper(char) {
			return true
		}
	}
	return false
}

func isLowercaseCharacter(password string) bool {
	for _, char := range password {
		if unicode.IsLower(char) {
			return true
		}
	}
	return false
}

func isSpecialCharacter(password string) bool {
	for _, char := range password {
		if !unicode.IsLetter(char) && !unicode.IsNumber(char) {
			return true
		}
	}
	return false
}

func isNumber(password string) bool {
	for _, char := range password {
		if unicode.IsNumber(char) {
			return true
		}
	}
	return false
}
