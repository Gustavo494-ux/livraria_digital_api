package controllers

import (
	"livraria_digital/src/interfaces"
	"livraria_digital/src/models"
	"livraria_digital/src/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IUsuarioController interface {
	interfaces.Controller
}
type UsuarioController struct {
	service services.IUsuarioService
}

// NewUsuarioController cria uma nova instância do controller de usuário
func NewUsuarioController(service services.IUsuarioService) IUsuarioController {
	return &UsuarioController{
		service: service,
	}
}

// Criar: cria um novo registro no banco de dados
func (u *UsuarioController) Criar(ctx *gin.Context) {
	usuario := models.Usuario{}
	if err := ctx.ShouldBindJSON(&usuario); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error:": err.Error()})
		return
	}

	if err := u.service.Criar(&usuario); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, usuario)
}

// BuscarTodos: busca todos os registros no banco de dados
func (u *UsuarioController) BuscarTodos(ctx *gin.Context) {
	pagincacao := retornarPaginacao(ctx)

	resultado, err := u.service.BuscarTodos(pagincacao)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, resultado)
}

// BuscarPorId: busca o registro com Id fornecido
func (u *UsuarioController) BuscarPorId(ctx *gin.Context) {
	usuarioBanco, err := u.service.BuscarPorId(retornarId(ctx))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, usuarioBanco)
}

// Atualizar: atualiza um registro no banco de dados
func (u *UsuarioController) Atualizar(ctx *gin.Context) {
	id := retornarId(ctx)

	usuario := &models.Usuario{}
	if err := ctx.ShouldBindJSON(usuario); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := u.service.Atualizar(id, usuario)
	if err != nil {
		if err.Error() == "usuário não encontrado" {
			ctx.JSON(http.StatusNotFound, gin.H{"": err.Error()})
			return
		}
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, *usuario)
}

// Deletar: remove um registro do banco de dados pelo ID
func (u *UsuarioController) Deletar(ctx *gin.Context) {
	if err := u.service.Deletar(retornarId(ctx)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}
