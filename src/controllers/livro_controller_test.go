package controllers

import (
	"livraria_digital/src/models"
	"livraria_digital/src/repository"
	"livraria_digital/src/services"
	"net/http"
	"testing"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// setupTest inicializa o banco de dados em memória e retorna o controller para testes
func setupTest() (*GenericRepository[models.Livro], *gorm.DB) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("Falha ao conectar ao banco de dados de teste")
	}
	db.AutoMigrate(&models.Livro{})

	repo := repository.NewGenericRepository[models.Livro](db)
	service := services.NewGenericServices[models.Livro](repo)
	controller := NewGenericController[models.Livro](service)

	return controller, db
}

// parseTime converte uma string no formato RFC3339 para time.Time
func parseTime(timeStr string) time.Time {
	t, err := time.Parse(time.RFC3339, timeStr)
	if err != nil {
		panic("Falha ao converter string para time.Time")
	}
	return t
}

func TestLivroController(t *testing.T) {
	livroController, _ := setupTest()

	// Teste de criação de livro
	t.Run("Criar Livro", func(t *testing.T) {
		payload := models.Livro{
			Titulo:         "1984",
			Autor:          "George Orwell",
			Estoque:        10,
			ISBN:           "9780451524935",
			Preco:          29.90,
			DataPublicacao: parseTime("1949-06-08T00:00:00Z"),
		}
		TesteCriarGenerico(t, livroController, payload, http.StatusCreated, "")
	})

	// Teste de listagem de todos os livros
	t.Run("Buscar Todos os Livros", func(t *testing.T) {
		TesteBuscarTodosGenerico(t, livroController, http.StatusOK, 1)
	})

	// Teste de busca de livro por ID
	t.Run("Buscar Livro por ID", func(t *testing.T) {
		TesteBuscarPorIdGenerico(t, livroController, "1", http.StatusOK, "")
	})

	// Teste de atualização de livro
	t.Run("Atualizar Livro", func(t *testing.T) {
		payload := models.Livro{
			Titulo:         "1984 - Edição Especial",
			Autor:          "George Orwell",
			Estoque:        15,
			ISBN:           "9780451524935",
			Preco:          34.90,
			DataPublicacao: parseTime("1949-06-08T00:00:00Z"),
		}
		TesteAtualizarGenerico(t, livroController, "1", payload, http.StatusOK, "")
	})

	// Teste de exclusão de livro
	t.Run("Deletar Livro", func(t *testing.T) {
		TesteDeletarGenerico(t, livroController, "1", http.StatusNoContent, "")
	})
}
