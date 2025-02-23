package repository

import (
	"livraria_digital/src/interfaces"
	"livraria_digital/src/models"

	"gorm.io/gorm"
)

type IUsuarioRepository interface {
	interfaces.Repository[models.Usuario]
}

type UsuarioRepository struct {
	GenericRepository[models.Usuario]
}

// NewUsuarioRepository cria uma nova instância do repositório de usuario
func NewUsuarioRepository(db *gorm.DB) *UsuarioRepository {
	return &UsuarioRepository{
		GenericRepository: *NewGenericRepository[models.Usuario](db),
	}
}
