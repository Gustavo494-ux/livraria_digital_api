package models

import (
	"time"

	"gorm.io/gorm"
)

type Livro struct {
	gorm.Model
	GenericModel[Livro]
	Titulo         string    `json:"titulo" gorm:"not null" validate:"required,min=2,max=100"`
	Autor          string    `json:"autor" gorm:"not null" validate:"required,min=2,max=100"`
	Estoque        int       `json:"estoque" gorm:"not null" validate:"required,min=0"`
	ISBN           string    `json:"isbn" gorm:"unique;not null" validate:"required,min=10,max=13"`
	Preco          float64   `json:"preco" gorm:"not null" validate:"required,min=0"`
	DataPublicacao time.Time `json:"data_publicacao" gorm:"not null" validate:"required"`
}

// IsID: verifica se o Id é válido. ID != 0
func (livro *Livro) IsID() bool {
	return livro.ID != 0
}

// IsEstoque: verifica se o estoque é positivo
func (livro *Livro) IsEstoque() bool {
	return livro.Estoque > 0
}

// IsEstoqueNegativo: verifica se o estoque é negativo
func (livro *Livro) IsEstoqueNegativo() bool {
	return livro.Estoque < 0
}

// IsPreco: verifica se o preço é positivo
func (livro *Livro) IsPreco() bool {
	return livro.Preco > 0
}

// IsPrecoNegativo: verifica se o preço é negativo
func (livro *Livro) IsPrecoNegativo() bool {
	return livro.Preco < 0
}
