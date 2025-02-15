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

// validateLivro: verifica se o livro possui dados válidos
func (s *LivroService) validateLivro(livro *models.Livro) (err error) {
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

// CreateLivro: adiciona um novo livro no banco
func (s *LivroService) CreateLivro(livro *models.Livro) error {
	if err := s.validateLivro(livro); err != nil {
		return err
	}
	return s.Repo.Criar(livro)
}

// UpdateLivro: atualiza os dados de um livro
func (s *LivroService) UpdateLivro(livro *models.Livro) error {
	if err := s.validateLivro(livro); err != nil {
		return err
	}
	return s.Repo.Atualizar(livro)
}

// GetAllLivros: busca todos os livros do banco
func (s *LivroService) GetAllLivros() ([]models.Livro, error) {
	return s.Repo.BuscarTodos()
}

// GetLivroByID: Busca um livro pelo ID
func (s *LivroService) GetLivroByID(id uint) (livro models.Livro) {
	livroParametro := models.Livro{}
	livroParametro.ID = id
	livro, _ = s.Repo.BuscarPrimeiro(livroParametro)
	return
}

// DeleteLivro: deleta um livro do banco
func (s *LivroService) DeleteLivro(id uint) (err error) {
	livroBanco := models.Livro{}
	livroBanco.ID = id

	livroBanco, err = s.Repo.BuscarPrimeiro(livroBanco)
	if err != nil {
		return err
	}

	if !livroBanco.IsID() {
		return errors.New("livro não encontrado")
	}

	return s.Repo.Deletar(livroBanco)
}
