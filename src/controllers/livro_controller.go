package controllers

import (
	"livraria_digital/src/models"
	"livraria_digital/src/services"
)

type LivroController struct {
	GenericRepository[models.Livro]
}

func NewLivroController(service *services.GenericServices[models.Livro]) *LivroController {
	return &LivroController{
		GenericRepository: *NewGenericController[models.Livro](service),
	}
}
