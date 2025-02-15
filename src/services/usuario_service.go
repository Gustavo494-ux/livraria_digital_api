package services

import (
	"livraria_digital/src/models"
	"livraria_digital/src/repository"
)

type UsuarioService struct {
	GenericServices[models.Usuario]
}

// Alterado para aceitar um ponteiro para UsuarioRepository
func NewUsuarioService(repo *repository.GenericRepository[models.Usuario]) *UsuarioService {
	return &UsuarioService{
		GenericServices: *NewGenericServices[models.Usuario](repo),
	}
}

// ToGenericService: retorna uma instância de GenericServices
func (u *UsuarioService) ToGenericService() *GenericServices[models.Usuario] {
	return &u.GenericServices
}
