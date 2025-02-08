package services

import (
	"errors"
	"livraria_digital/src/models"
	"livraria_digital/src/repository"

	"github.com/go-playground/validator/v10"
)

type LivroService struct {
	Repo      *repository.LivroRepository
	Validator *validator.Validate
}

func NewLivroService(repo *repository.LivroRepository) *LivroService {
	return &LivroService{
		Repo:      repo,
		Validator: validator.New(),
	}
}

func (s *LivroService) validateLivro(livro *models.Livro) error {
	// Validações básicas usando o pacote validator
	if err := s.Validator.Struct(livro); err != nil {
		return err
	}

	// Verificação adicional: estoque não pode ser negativo
	if livro.Estoque < 0 {
		return errors.New("estoque não pode ser negativo")
	}

	// Verificação adicional: preço não pode ser negativo
	if livro.Preco < 0 {
		return errors.New("preço não pode ser negativo")
	}

	// Verificação adicional: ISBN deve ser único
	existingLivro, err := s.Repo.FindByISBN(livro.ISBN)
	if err == nil && existingLivro.ID != 0 && existingLivro.ID != livro.ID {
		return errors.New("ISBN já está em uso")
	}

	return nil
}

func (s *LivroService) CreateLivro(livro *models.Livro) error {
	if err := s.validateLivro(livro); err != nil {
		return err
	}
	return s.Repo.Create(livro)
}

func (s *LivroService) UpdateLivro(livro *models.Livro) error {
	if err := s.validateLivro(livro); err != nil {
		return err
	}
	return s.Repo.Update(livro)
}

func (s *LivroService) GetAllLivros() ([]models.Livro, error) {
	return s.Repo.FindAll()
}

func (s *LivroService) GetLivroByID(id uint) (*models.Livro, error) {
	return s.Repo.FindByID(id)
}

func (s *LivroService) DeleteLivro(id uint) error {
	return s.Repo.Delete(id)
}
