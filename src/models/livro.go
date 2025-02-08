package models

import (
	"time"

	"gorm.io/gorm"
)

type Livro struct {
	gorm.Model
	Titulo         string    `json:"titulo" gorm:"not null" validate:"required,min=2,max=100"`
	Autor          string    `json:"autor" gorm:"not null" validate:"required,min=2,max=100"`
	Estoque        int       `json:"estoque" gorm:"not null" validate:"required,min=0"`
	ISBN           string    `json:"isbn" gorm:"unique;not null" validate:"required,min=10,max=13"`
	Preco          float64   `json:"preco" gorm:"not null" validate:"required,min=0"`
	DataPublicacao time.Time `json:"data_publicacao" gorm:"not null" validate:"required"`
}
