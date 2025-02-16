package controllers

import (
	"livraria_digital/src/models"
	"livraria_digital/src/services"
)

type LivroController struct {
	GenericController[models.Livro]
}

func NewLivroController(service *services.GenericServices[models.Livro]) *LivroController {
	return &LivroController{
		GenericController: *NewGenericController[models.Livro](service),
	}
}
