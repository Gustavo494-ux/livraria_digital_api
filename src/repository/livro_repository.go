package repository

import (
	"livraria_digital/src/models"

	"gorm.io/gorm"
)

type LivroRepository struct {
	DB *gorm.DB
}

func NewLivroRepository(db *gorm.DB) *LivroRepository {
	return &LivroRepository{DB: db}
}

// Create: cria um novo livro
func (r *LivroRepository) Create(livro *models.Livro) error {
	return r.DB.Create(livro).Error
}

// FindAll: busca todos os livros
func (r *LivroRepository) FindAll() ([]models.Livro, error) {
	var livros []models.Livro
	err := r.DB.Find(&livros).Error
	return livros, err
}

// BuscarLivros: realiza uma busca genérica por livros que possuem dados iguais aos definidos para o parametro livro
func (r *LivroRepository) BuscarLivros(livro models.Livro) (livros []*models.Livro) {
	r.DB.Find(&livros, livro)
	return
}

// BuscarLivro: realiza uma busca genérica por um livro que possui dados iguais aos definidos para o parametro livro
func (r *LivroRepository) BuscarLivro(livro models.Livro) (livroBanco *models.Livro) {
	r.DB.First(&livroBanco, livro)
	return
}

// Update: atualiza os dados de um livro
func (r *LivroRepository) Update(livro *models.Livro) error {
	return r.DB.Save(livro).Error
}

// Delete: deleta um livro
func (r *LivroRepository) Delete(id uint) error {
	return r.DB.Delete(&models.Livro{}, id).Error
}
