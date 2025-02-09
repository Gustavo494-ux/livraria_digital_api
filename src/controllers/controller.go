package controllers

import (
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
