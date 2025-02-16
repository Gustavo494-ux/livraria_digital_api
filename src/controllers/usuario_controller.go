package controllers

import (
	"livraria_digital/src/models"
	"livraria_digital/src/services"
)

type UsuarioController struct {
	GenericController[models.Usuario]
}

func NewUsuarioController(service *services.GenericServices[models.Usuario]) *UsuarioController {
	return &UsuarioController{
		GenericController: *NewGenericController[models.Usuario](service),
	}
}

// ToGenericController: retorna uma instância de GenericController
func (u *UsuarioController) ToGenericController() *GenericController[models.Usuario] {
	return &u.GenericController
}
