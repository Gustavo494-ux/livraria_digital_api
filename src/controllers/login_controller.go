package controllers

import (
	"livraria_digital/src/models"
	"livraria_digital/src/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ILoginController interface {
	RealizarLogin(ctx *gin.Context)
}
type LoginController struct {
	service services.ILoginService
}

// NewLoginController cria uma nova instância do controller de login
func NewLoginController(service services.ILoginService) ILoginController {
	return &LoginController{
		service: service,
	}
}

// RealizarLogin: realiza as verificações necessárias para definir se o usuário pode realizar o login
func (c *LoginController) RealizarLogin(ctx *gin.Context) {
	login := models.Login{}
	if err := ctx.ShouldBindJSON(&login); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := c.service.RealizarLogin(login)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, err.Error())
		return
	}

	resultado := map[string]string{
		"sucesso": "true",
		"token":   token,
	}
	ctx.JSON(http.StatusOK, resultado)
}
