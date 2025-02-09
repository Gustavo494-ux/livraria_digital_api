package controllers

import (
	"bytes"
	"encoding/json"
	"livraria_digital/src/models"
	"livraria_digital/src/repository"
	"livraria_digital/src/services"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// setupTest inicializa o banco de dados em memória e retorna o controller para testes
func setupTest() (*LivroController, *gorm.DB) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("Falha ao conectar ao banco de dados de teste")
	}
	db.AutoMigrate(&models.Livro{})

	livroRepo := repository.NewLivroRepository(db)
	livroService := services.NewLivroService(livroRepo)
	livroController := NewLivroController(livroService)

	return livroController, db
}

// parseTime converte uma string no formato RFC3339 para time.Time
func parseTime(timeStr string) time.Time {
	t, err := time.Parse(time.RFC3339, timeStr)
	if err != nil {
		panic("Falha ao converter string para time.Time")
	}
	return t
}

// TestCreateLivro testa a criação de um livro com diferentes cenários
func TestCreateLivro(t *testing.T) {
	livroController, db := setupTest()

	tests := []struct {
		name           string
		setup          func() // Função de setup específica para cada cenário
		payload        interface{}
		expectedStatus int
		expectedError  string
	}{
		{
			name: "Sucesso - Criação de livro válido",
			setup: func() {
				// Nenhum setup necessário, pois não há pré-condições
			},
			payload: models.Livro{
				Titulo:         "1984",
				Autor:          "George Orwell",
				Estoque:        10,
				ISBN:           "9780451524935",
				Preco:          29.90,
				DataPublicacao: parseTime("1949-06-08T00:00:00Z"),
			},
			expectedStatus: http.StatusCreated,
			expectedError:  "",
		},
		{
			name: "Falha - JSON inválido",
			setup: func() {
				// Nenhum setup necessário
			},
			payload:        `{"titulo": "1984", "autor": "George Orwell", "estoque": "dez"}`, // Estoque como string
			expectedStatus: http.StatusBadRequest,
			expectedError:  "error",
		},
		{
			name: "Falha - Campos obrigatórios faltando",
			setup: func() {
				// Nenhum setup necessário
			},
			payload: models.Livro{
				Titulo: "1984", // Faltam Autor, Estoque, ISBN, Preco, DataPublicacao
			},
			expectedStatus: http.StatusBadRequest,
			expectedError:  "error",
		},
		{
			name: "Falha - ISBN duplicado",
			setup: func() {
				// Insere um livro com ISBN duplicado para o cenário de falha
				db.Create(&models.Livro{
					Titulo:         "1984",
					Autor:          "George Orwell",
					Estoque:        10,
					ISBN:           "9780451524935",
					Preco:          29.90,
					DataPublicacao: parseTime("1949-06-08T00:00:00Z"),
				})
			},
			payload: models.Livro{
				Titulo:         "Outro Livro",
				Autor:          "Autor",
				Estoque:        5,
				ISBN:           "9780451524935", // ISBN duplicado
				Preco:          19.90,
				DataPublicacao: parseTime("2020-01-01T00:00:00Z"),
			},
			expectedStatus: http.StatusBadRequest,
			expectedError:  "error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Executa o setup específico para o cenário atual
			tt.setup()

			payloadJSON, _ := json.Marshal(tt.payload)
			req, _ := http.NewRequest("POST", "/livros", bytes.NewBuffer(payloadJSON))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = req

			livroController.CreateLivro(ctx)

			assert.Equal(t, tt.expectedStatus, w.Code)
			if tt.expectedError != "" {
				assert.Contains(t, w.Body.String(), tt.expectedError)
			}
		})
	}
}

// TestGetAllLivros testa a listagem de todos os livros
func TestGetAllLivros(t *testing.T) {
	livroController, db := setupTest()

	tests := []struct {
		name           string
		setup          func()
		expectedStatus int
		expectedCount  int
	}{
		{
			name: "Sucesso - Lista com livros cadastrados",
			setup: func() {
				db.Create(&models.Livro{
					Titulo:         "1984",
					Autor:          "George Orwell",
					Estoque:        10,
					ISBN:           "9780451524935",
					Preco:          29.90,
					DataPublicacao: parseTime("1949-06-08T00:00:00Z"),
				})
				db.Create(&models.Livro{
					Titulo:         "A Revolução dos Bichos",
					Autor:          "George Orwell",
					Estoque:        5,
					ISBN:           "9780451526342",
					Preco:          24.90,
					DataPublicacao: parseTime("1945-08-17T00:00:00Z"),
				})
			},
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name: "Sucesso - Lista vazia",
			setup: func() {
				db.Where("1=1").Delete(&models.Livro{})
			},
			expectedStatus: http.StatusOK,
			expectedCount:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()

			req, _ := http.NewRequest("GET", "/livros", nil)
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = req

			livroController.GetAllLivros(ctx)

			assert.Equal(t, tt.expectedStatus, w.Code)

			var responseLivros []models.Livro
			json.Unmarshal(w.Body.Bytes(), &responseLivros)
			assert.Equal(t, tt.expectedCount, len(responseLivros))
		})
	}
}

// TestGetLivroByID testa a busca de um livro por ID
func TestGetLivroByID(t *testing.T) {
	livroController, db := setupTest()

	tests := []struct {
		name           string
		setup          func()
		livroID        string
		expectedStatus int
		expectedError  string
	}{
		{
			name: "Sucesso - Livro encontrado",
			setup: func() {
				db.Create(&models.Livro{
					Titulo:         "1984",
					Autor:          "George Orwell",
					Estoque:        10,
					ISBN:           "9780451524935",
					Preco:          29.90,
					DataPublicacao: parseTime("1949-06-08T00:00:00Z"),
				})
			},
			livroID:        "1",
			expectedStatus: http.StatusOK,
			expectedError:  "",
		},
		{
			name:           "Falha - Livro não encontrado",
			setup:          func() {}, // Nenhum livro cadastrado
			livroID:        "999",
			expectedStatus: http.StatusNotFound,
			expectedError:  "error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()

			req, _ := http.NewRequest("GET", "/livros/"+tt.livroID, nil)
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = req
			ctx.Params = []gin.Param{{Key: "id", Value: tt.livroID}}

			livroController.GetLivroByID(ctx)

			assert.Equal(t, tt.expectedStatus, w.Code)
			if tt.expectedError != "" {
				assert.Contains(t, w.Body.String(), tt.expectedError)
			}
		})
	}
}

// TestUpdateLivro testa a atualização de um livro
func TestUpdateLivro(t *testing.T) {
	livroController, db := setupTest()

	tests := []struct {
		name           string
		setup          func()
		livroID        string
		payload        interface{}
		expectedStatus int
		expectedError  string
	}{
		{
			name: "Sucesso - Atualização de livro válido",
			setup: func() {
				db.Create(&models.Livro{
					Titulo:         "1984",
					Autor:          "George Orwell",
					Estoque:        10,
					ISBN:           "9780451524935",
					Preco:          29.90,
					DataPublicacao: parseTime("1949-06-08T00:00:00Z"),
				})
			},
			livroID: "1",
			payload: models.Livro{
				Titulo:         "1984 - Edição Especial",
				Autor:          "George Orwell",
				Estoque:        15,
				ISBN:           "9780451524935",
				Preco:          34.90,
				DataPublicacao: parseTime("1949-06-08T00:00:00Z"),
			},
			expectedStatus: http.StatusOK,
			expectedError:  "",
		},
		{
			name:    "Falha - Livro não encontrado",
			setup:   func() {}, // Nenhum livro cadastrado
			livroID: "999",
			payload: models.Livro{
				Titulo:         "1984 - Edição Especial",
				Autor:          "George Orwell",
				Estoque:        15,
				ISBN:           "9780451524935",
				Preco:          34.90,
				DataPublicacao: parseTime("1949-06-08T00:00:00Z"),
			},
			expectedStatus: http.StatusBadRequest,
			expectedError:  "error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()

			payloadJSON, _ := json.Marshal(tt.payload)
			req, _ := http.NewRequest("PUT", "/livros/"+tt.livroID, bytes.NewBuffer(payloadJSON))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = req
			ctx.Params = []gin.Param{{Key: "id", Value: tt.livroID}}

			livroController.UpdateLivro(ctx)

			assert.Equal(t, tt.expectedStatus, w.Code)
			if tt.expectedError != "" {
				assert.Contains(t, w.Body.String(), tt.expectedError)
			}
		})
	}
}

// TestDeleteLivro testa a exclusão de um livro
func TestDeleteLivro(t *testing.T) {
	livroController, db := setupTest()

	tests := []struct {
		name           string
		setup          func()
		livroID        string
		expectedStatus int
		expectedError  string
	}{
		{
			name: "Sucesso - Exclusão de livro válido",
			setup: func() {
				db.Create(&models.Livro{
					Titulo:         "1984",
					Autor:          "George Orwell",
					Estoque:        10,
					ISBN:           "9780451524935",
					Preco:          29.90,
					DataPublicacao: parseTime("1949-06-08T00:00:00Z"),
				})
			},
			livroID:        "1",
			expectedStatus: http.StatusNoContent,
			expectedError:  "",
		},
		{
			name:           "Falha - Livro não encontrado",
			setup:          func() {}, // Nenhum livro cadastrado
			livroID:        "999",
			expectedStatus: http.StatusInternalServerError,
			expectedError:  "error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()

			req, _ := http.NewRequest("DELETE", "/livros/"+tt.livroID, nil)
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = req
			ctx.Params = []gin.Param{{Key: "id", Value: tt.livroID}}

			livroController.DeleteLivro(ctx)

			assert.Equal(t, tt.expectedStatus, w.Code)
			if tt.expectedError != "" {
				assert.Contains(t, w.Body.String(), tt.expectedError)
			}
		})
	}
}
