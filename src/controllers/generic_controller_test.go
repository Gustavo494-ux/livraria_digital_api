package controllers

import (
	"bytes"
	"encoding/json"
	"livraria_digital/src/interfaces"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// TesteCriarGenerico testa a criação de uma entidade genérica
func TesteCriarGenerico(t *testing.T, controller interfaces.Controller, payload interface{}, expectedStatus int, expectedError string) {
	payloadJSON, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(payloadJSON))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = req

	controller.Criar(ctx)

	assert.Equal(t, expectedStatus, w.Code)
	if expectedError != "" {
		assert.Contains(t, w.Body.String(), expectedError)
	}
}

// TestGetAllGeneric testa a listagem de todas as entidades genéricas
func TesteBuscarTodosGenerico(t *testing.T, controller interfaces.Controller, expectedStatus int, expectedCount int) {
	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = req

	controller.BuscarTodos(ctx)

	assert.Equal(t, expectedStatus, w.Code)

	var response []map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, expectedCount, len(response))
}

// TesteBuscarPorIdGenerico testa a busca de uma entidade genérica por ID
func TesteBuscarPorIdGenerico(t *testing.T, controller interfaces.Controller, entityID string, expectedStatus int, expectedError string) {
	req, _ := http.NewRequest("GET", "/"+entityID, nil)
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = req
	ctx.Params = []gin.Param{{Key: "id", Value: entityID}}

	controller.BuscarPorId(ctx)

	assert.Equal(t, expectedStatus, w.Code)
	if expectedError != "" {
		assert.Contains(t, w.Body.String(), expectedError)
	}
}

// TesteAtualizarGenerico testa a atualização de uma entidade genérica
func TesteAtualizarGenerico(t *testing.T, controller interfaces.Controller, entityID string, payload interface{}, expectedStatus int, expectedError string) {
	payloadJSON, _ := json.Marshal(payload)
	req, _ := http.NewRequest("PUT", "/"+entityID, bytes.NewBuffer(payloadJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = req
	ctx.Params = []gin.Param{{Key: "id", Value: entityID}}

	controller.Atualizar(ctx)

	assert.Equal(t, expectedStatus, w.Code)
	if expectedError != "" {
		assert.Contains(t, w.Body.String(), expectedError)
	}
}

// TesteDeletarGenerico testa a exclusão de uma entidade genérica
func TesteDeletarGenerico(t *testing.T, controller interfaces.Controller, entityID string, expectedStatus int, expectedError string) {
	req, _ := http.NewRequest("DELETE", "/"+entityID, nil)
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = req
	ctx.Params = []gin.Param{{Key: "id", Value: entityID}}

	controller.Deletar(ctx)

	assert.Equal(t, expectedStatus, w.Code)
	if expectedError != "" {
		assert.Contains(t, w.Body.String(), expectedError)
	}
}
