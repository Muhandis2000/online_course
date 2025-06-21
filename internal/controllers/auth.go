package controllers

import (
	"net/http"

	"github.com/Muhandis2000/online-school/internal/models"
	"github.com/Muhandis2000/online-school/internal/services"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	service *services.AuthService
}

func NewAuthController(service *services.AuthService) *AuthController {
	return &AuthController{service}
}

func (c *AuthController) Register(ctx *gin.Context) {
	var user models.User
	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	if _, err := c.service.Register(user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Registration failed"})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "User registered"})
}

func (c *AuthController) Login(ctx *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	token, err := c.service.Login(req.Username, req.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Login failed"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
