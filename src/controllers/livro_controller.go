package controllers

import (
	"livraria_digital/src/models"
	"livraria_digital/src/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LivroController struct {
	Service *services.LivroService
}

func NewLivroController(service *services.LivroService) *LivroController {
	return &LivroController{Service: service}
}

func (c *LivroController) CreateLivro(ctx *gin.Context) {
	var livro models.Livro
	if err := ctx.ShouldBindJSON(&livro); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.Service.CreateLivro(&livro); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, livro)
}

func (c *LivroController) GetAllLivros(ctx *gin.Context) {
	livros, err := c.Service.GetAllLivros()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, livros)
}

func (c *LivroController) GetLivroByID(ctx *gin.Context) {
	id := retornarId(ctx)

	livro := c.Service.GetLivroByID(uint(id))
	if !livro.IsID() {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Livro não encontrado"})
		return
	}

	ctx.JSON(http.StatusOK, livro)
}

func (c *LivroController) UpdateLivro(ctx *gin.Context) {
	id := retornarId(ctx)

	var livro models.Livro
	if err := ctx.ShouldBindJSON(&livro); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	livro.ID = uint(id)
	if err := c.Service.UpdateLivro(&livro); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, livro)
}

func (c *LivroController) DeleteLivro(ctx *gin.Context) {
	id := retornarId(ctx)

	if err := c.Service.DeleteLivro(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
