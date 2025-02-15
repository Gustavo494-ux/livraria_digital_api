package repository

import (
	"livraria_digital/src/models"

	"gorm.io/gorm"
)

type UsuarioRepository struct {
	GenericRepository[models.Usuario]
}

// NewUsuarioRepository cria uma nova instância do repositório de usuario
func NewUsuarioRepository(db *gorm.DB) *UsuarioRepository {
	return &UsuarioRepository{
		GenericRepository: *NewGenericRepository[models.Usuario](db),
	}
}

// ToGenericRepository: retorna uma instância de GenericRepository
func (u *UsuarioRepository) ToGenericRepository() *GenericRepository[models.Usuario] {
	return &u.GenericRepository
}
