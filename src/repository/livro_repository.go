package repository

import (
	"livraria_digital/src/models"

	"gorm.io/gorm"
)

type LivroRepository struct {
	GenericRepository[models.Livro]
}

// NewLivroRepository cria uma nova instância do repositório de Livro
func NewLivroRepository(db *gorm.DB) *LivroRepository {
	return &LivroRepository{
		GenericRepository: *NewGenericRepository[models.Livro](db), // Inicializa o repositório genérico
	}
}
