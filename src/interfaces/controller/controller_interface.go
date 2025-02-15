package interfaces

import "github.com/gin-gonic/gin"

type Controller[T any] interface {
	Criar(ctx *gin.Context)
	BuscarTodos(ctx *gin.Context)
	BuscarPorId(ctx *gin.Context)
	Atualizar(ctx *gin.Context)
	Deletar(ctx *gin.Context)
}
