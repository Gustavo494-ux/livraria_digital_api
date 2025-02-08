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

func (r *LivroRepository) Create(livro *models.Livro) error {
	return r.DB.Create(livro).Error
}

func (r *LivroRepository) FindAll() ([]models.Livro, error) {
	var livros []models.Livro
	err := r.DB.Find(&livros).Error
	return livros, err
}

func (r *LivroRepository) FindByID(id uint) (*models.Livro, error) {
	var livro models.Livro
	err := r.DB.First(&livro, id).Error
	return &livro, err
}

func (r *LivroRepository) FindByISBN(isbn string) (*models.Livro, error) {
	var livro models.Livro
	err := r.DB.Where("isbn = ?", isbn).First(&livro).Error
	return &livro, err
}

func (r *LivroRepository) Update(livro *models.Livro) error {
	return r.DB.Save(livro).Error
}

func (r *LivroRepository) Delete(id uint) error {
	return r.DB.Delete(&models.Livro{}, id).Error
}
