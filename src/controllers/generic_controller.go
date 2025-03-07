package controllers

import (
	Generics "livraria_digital/src/pkg"
	"livraria_digital/src/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GenericController é uma estrutura genérica para operações de CRUD na camada controller
type GenericController[T any] struct {
	Service *services.GenericServices[T]
}

// NewGenericController cria uma nova instância do controller genérico
func NewGenericController[T any](service *services.GenericServices[T]) *GenericController[T] {
	return &GenericController[T]{Service: service}
}

// Criar: cria um novo registro no banco de dados
func (c *GenericController[T]) Criar(ctx *gin.Context) {
	var entidade T
	if err := ctx.ShouldBindJSON(&entidade); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.Service.Criar(&entidade); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, entidade)
}

// BuscarTodos: busca todos os registros no banco de dados
func (c *GenericController[T]) BuscarTodos(ctx *gin.Context) {
	pagincao := retornarPaginacao(ctx)

	entidades, err := c.Service.BuscarTodos(pagincao)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, entidades)
}

// Buscar: busca o registro com Id fornecido
func (c *GenericController[T]) BuscarPorId(ctx *gin.Context) {
	entidadeBanco, err := c.Service.BuscarPorId(retornarId(ctx))
	if err != nil {
		ctx.Error(err)
		return
	}

	if Generics.RetornarCampo(entidadeBanco, "ID", uint(0)).(uint) < 1 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "registro não encontrado"})
		return
	}

	ctx.JSON(http.StatusOK, entidadeBanco)
}

// Atualizar: atualiza um registro no banco de dados
func (c *GenericController[T]) Atualizar(ctx *gin.Context) {
	id := retornarId(ctx)

	var entidade T
	if err := ctx.ShouldBindJSON(&entidade); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := c.Service.Atualizar(id, &entidade)

	if err != nil {
		if err.Error() == "nenhum registro encontrado" {
			ctx.JSON(http.StatusNotFound, gin.H{"": err.Error()})
			return
		}
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, entidade)
}

// Deletar: remove um registro do banco de dados pelo ID
func (c *GenericController[T]) Deletar(ctx *gin.Context) {
	if err := c.Service.Deletar(retornarId(ctx)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
