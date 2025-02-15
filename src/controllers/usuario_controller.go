package controllers

import (
	"livraria_digital/src/models"
	"livraria_digital/src/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UsuarioController struct {
	Service *services.UsuarioService
}

func NewUsuarioController(service *services.UsuarioService) *UsuarioController {
	return &UsuarioController{Service: service}
}

func (c *UsuarioController) Criar(ctx *gin.Context) {
	var usuario models.Usuario
	if err := ctx.ShouldBindJSON(&usuario); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.Service.Criar(&usuario); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, usuario)
}

func (c *UsuarioController) BuscarTodos(ctx *gin.Context) {
	usuario, err := c.Service.BuscarTodos()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, usuario)
}

func (c *UsuarioController) BuscarPorId(ctx *gin.Context) {
	usuario := models.Usuario{}
	usuario.ID = uint(retornarId(ctx))

	usuario, err := c.Service.BuscarPrimeiro(usuario)
	if err != nil {
		ctx.Error(err)
		return
	}

	if !usuario.IsID() {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
		return
	}

	ctx.JSON(http.StatusOK, usuario)
}

func (c *UsuarioController) Atualizar(ctx *gin.Context) {
	id := retornarId(ctx)

	var usuario models.Usuario
	if err := ctx.ShouldBindJSON(&usuario); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	usuario.ID = uint(id)
	if err := c.Service.Atualizar(id, &usuario); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, usuario)
}

func (c *UsuarioController) Deletar(ctx *gin.Context) {
	if err := c.Service.Deletar(retornarId(ctx)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
