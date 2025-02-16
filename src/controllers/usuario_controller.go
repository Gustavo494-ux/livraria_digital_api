package controllers

import (
	"livraria_digital/src/models"
	"livraria_digital/src/services"
)

type UsuarioController struct {
	GenericRepository[models.Usuario]
}

func NewUsuarioController(service *services.GenericServices[models.Usuario]) *UsuarioController {
	return &UsuarioController{
		GenericRepository: *NewGenericController[models.Usuario](service),
	}
}
