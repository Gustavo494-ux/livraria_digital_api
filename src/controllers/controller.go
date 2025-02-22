package controllers

import (
	"livraria_digital/src/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func retornarId(ctx *gin.Context) (id int) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	return
}

// retornaPaginacao: retorna uma instância de paginação com os dados da requisição
func retornarPaginacao(ctx *gin.Context) (paginacao models.Paginacao) {
	paginacao.PaginaAtual, _ = strconv.Atoi(ctx.DefaultQuery("pagina", "1"))
	paginacao.Limite, _ = strconv.Atoi(ctx.DefaultQuery("limite", "100"))

	return paginacao
}
