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
	// Configura o banco de dados em memória
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("Falha ao conectar ao banco de dados de teste")
	}

	// Auto-migrate (cria as tabelas no banco de dados)
	db.AutoMigrate(&models.Livro{})

	// Inicializa o repository, service e controller
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

// TestCreateLivro testa a criação de um livro
func TestCreateLivro(t *testing.T) {
	livroController, _ := setupTest()

	// Cria um livro de exemplo
	livro := models.Livro{
		Titulo:         "1984",
		Autor:          "George Orwell",
		Estoque:        10,
		ISBN:           "9780451524935",
		Preco:          29.90,
		DataPublicacao: parseTime("1949-06-08T00:00:00Z"),
	}

	// Converte o livro para JSON
	livroJSON, _ := json.Marshal(livro)

	// Cria uma requisição HTTP POST
	req, _ := http.NewRequest("POST", "/livros", bytes.NewBuffer(livroJSON))
	req.Header.Set("Content-Type", "application/json")

	// Grava a resposta
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = req

	// Executa o método do controller
	livroController.CreateLivro(ctx)

	// Verifica o status code
	assert.Equal(t, http.StatusCreated, w.Code)

	// Verifica o corpo da resposta
	var responseLivro models.Livro
	json.Unmarshal(w.Body.Bytes(), &responseLivro)
	assert.Equal(t, livro.Titulo, responseLivro.Titulo)
	assert.Equal(t, livro.Autor, responseLivro.Autor)
}

// TestGetAllLivros testa a listagem de todos os livros
func TestGetAllLivros(t *testing.T) {
	livroController, db := setupTest()

	// Cria dois livros de exemplo
	livro1 := models.Livro{
		Titulo:         "1984",
		Autor:          "George Orwell",
		Estoque:        10,
		ISBN:           "9780451524935",
		Preco:          29.90,
		DataPublicacao: parseTime("1949-06-08T00:00:00Z"),
	}
	livro2 := models.Livro{
		Titulo:         "A Revolução dos Bichos",
		Autor:          "George Orwell",
		Estoque:        5,
		ISBN:           "9780451526342",
		Preco:          24.90,
		DataPublicacao: parseTime("1945-08-17T00:00:00Z"),
	}

	// Insere os livros no banco de dados
	db.Create(&livro1)
	db.Create(&livro2)

	// Cria uma requisição HTTP GET
	req, _ := http.NewRequest("GET", "/livros", nil)
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = req

	// Executa o método do controller
	livroController.GetAllLivros(ctx)

	// Verifica o status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Verifica o corpo da resposta
	var responseLivros []models.Livro
	json.Unmarshal(w.Body.Bytes(), &responseLivros)
	assert.Equal(t, 2, len(responseLivros))
}

// TestGetLivroByID testa a busca de um livro por ID
func TestGetLivroByID(t *testing.T) {
	livroController, db := setupTest()

	// Cria um livro de exemplo
	livro := models.Livro{
		Titulo:         "1984",
		Autor:          "George Orwell",
		Estoque:        10,
		ISBN:           "9780451524935",
		Preco:          29.90,
		DataPublicacao: parseTime("1949-06-08T00:00:00Z"),
	}

	// Insere o livro no banco de dados
	db.Create(&livro)

	// Cria uma requisição HTTP GET
	req, _ := http.NewRequest("GET", "/livros/1", nil)
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = req
	ctx.Params = []gin.Param{{Key: "id", Value: "1"}}

	// Executa o método do controller
	livroController.GetLivroByID(ctx)

	// Verifica o status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Verifica o corpo da resposta
	var responseLivro models.Livro
	json.Unmarshal(w.Body.Bytes(), &responseLivro)
	assert.Equal(t, livro.Titulo, responseLivro.Titulo)
}

// TestUpdateLivro testa a atualização de um livro
func TestUpdateLivro(t *testing.T) {
	livroController, db := setupTest()

	// Cria um livro de exemplo
	livro := models.Livro{
		Titulo:         "1984",
		Autor:          "George Orwell",
		Estoque:        10,
		ISBN:           "9780451524935",
		Preco:          29.90,
		DataPublicacao: parseTime("1949-06-08T00:00:00Z"),
	}

	// Insere o livro no banco de dados
	db.Create(&livro)

	// Atualiza o livro
	livroAtualizado := models.Livro{
		Titulo:         "1984 - Edição Especial",
		Autor:          "George Orwell",
		Estoque:        15,
		ISBN:           "9780451524935",
		Preco:          34.90,
		DataPublicacao: parseTime("1949-06-08T00:00:00Z"),
	}

	// Converte o livro atualizado para JSON
	livroJSON, _ := json.Marshal(livroAtualizado)

	// Cria uma requisição HTTP PUT
	req, _ := http.NewRequest("PUT", "/livros/1", bytes.NewBuffer(livroJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = req
	ctx.Params = []gin.Param{{Key: "id", Value: "1"}}

	// Executa o método do controller
	livroController.UpdateLivro(ctx)

	// Verifica o status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Verifica o corpo da resposta
	var responseLivro models.Livro
	json.Unmarshal(w.Body.Bytes(), &responseLivro)
	assert.Equal(t, livroAtualizado.Titulo, responseLivro.Titulo)
}

// TestDeleteLivro testa a exclusão de um livro
func TestDeleteLivro(t *testing.T) {
	livroController, db := setupTest()

	// Cria um livro de exemplo
	livro := models.Livro{
		Titulo:         "1984",
		Autor:          "George Orwell",
		Estoque:        10,
		ISBN:           "9780451524935",
		Preco:          29.90,
		DataPublicacao: parseTime("1949-06-08T00:00:00Z"),
	}

	// Insere o livro no banco de dados
	db.Create(&livro)

	// Cria uma requisição HTTP DELETE
	req, _ := http.NewRequest("DELETE", "/livros/1", nil)
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = req
	ctx.Params = []gin.Param{{Key: "id", Value: "1"}}

	// Executa o método do controller
	livroController.DeleteLivro(ctx)

	// Verifica o status code
	assert.Equal(t, http.StatusNoContent, w.Code)

	// Verifica se o livro foi excluído
	var livroExcluido models.Livro
	db.First(&livroExcluido, 1)
	assert.Equal(t, uint(0), livroExcluido.ID)
}
