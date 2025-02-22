package services

import (
	"errors"
	"livraria_digital/src/interfaces"
	"livraria_digital/src/models"
	"livraria_digital/src/repository"
)

type LivroService interface {
	interfaces.Services[models.Livro]
	ToGenericService() *GenericServices[models.Livro]
}
type livroService struct {
	GenericServices[models.Livro]
}

func NewLivroService(repo *repository.GenericRepository[models.Livro]) LivroService {
	return &livroService{
		GenericServices: *NewGenericServices[models.Livro](repo),
	}
}

// ToGenericService: retorna uma instância de GenericServices
func (u *livroService) ToGenericService() *GenericServices[models.Livro] {
	return &u.GenericServices
}

// validar: verifica se o livro possui dados válidos
func (s *livroService) Validar(livro *models.Livro) (err error) {
	if err := s.Validator.Struct(livro); err != nil {
		return err
	}

	if livro.IsEstoqueNegativo() {
		return errors.New("estoque não pode ser negativo")
	}

	if livro.IsPrecoNegativo() {
		return errors.New("preço não pode ser negativo")
	}

	livroBanco, err := s.Repo.BuscarPrimeiro(models.Livro{ISBN: livro.ISBN})
	if err != nil {
		return err
	}

	if livroBanco.ID != livro.ID {
		return errors.New("ISBN já está em uso")
	}

	return
}
