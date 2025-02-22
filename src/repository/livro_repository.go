package repository

import (
	"livraria_digital/src/interfaces"
	"livraria_digital/src/models"

	"gorm.io/gorm"
)

type LivroRepository interface {
	interfaces.Repository[models.Livro]
	ToGenericRepository() *GenericRepository[models.Livro]
}

type livroRepository struct {
	GenericRepository[models.Livro]
}

// NewLivroRepository cria uma nova instância do repositório de Livro
func NewLivroRepository(db *gorm.DB) LivroRepository {
	return &livroRepository{
		GenericRepository: *NewGenericRepository[models.Livro](db),
	}
}

// ToGenericRepository: retorna uma instância de GenericRepository
func (u *livroRepository) ToGenericRepository() *GenericRepository[models.Livro] {
	return &u.GenericRepository
}
