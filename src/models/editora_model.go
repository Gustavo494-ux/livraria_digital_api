package models

import (
	"gorm.io/gorm"
)

type Editora struct {
	gorm.Model
	GenericModel[Editora]
	Nome     string `json:"nome" gorm:"not null" validate:"required,min=2,max=100"`
	Endereco string `json:"endereco" gorm:"not null" validate:"required,min=2,max=255"`
	Telefone string `json:"telefone" gorm:"not null" validate:"required,min=10,max=15"`
	Email    string `json:"email" gorm:"not null" validate:"required,email"`
}
