package repository

import (
	"livraria_digital/src/models"

	"gorm.io/gorm"
)

type ILoginRepository interface {
	BuscarUsuario(usuarioFiltro models.Usuario) (usuario models.Usuario, err error)
}

type LoginRepository struct {
	GenericRepository[models.Login]
}

// NewLoginRepository cria uma nova instância do repositório de Login
func NewLoginRepository(db *gorm.DB) *LoginRepository {
	return &LoginRepository{
		GenericRepository: *NewGenericRepository[models.Login](db),
	}
}

// BuscarUsuario: busca o usuário
func (r *LoginRepository) BuscarUsuario(usuarioFiltro models.Usuario) (usuario models.Usuario, err error) {
	usuarioRepository := NewUsuarioRepository(r.DB)
	usuario, err = usuarioRepository.BuscarPrimeiro(usuarioFiltro)
	return
}
