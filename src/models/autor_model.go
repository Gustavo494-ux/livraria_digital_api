package models

import (
	"time"

	"gorm.io/gorm"
)

type Autor struct {
	gorm.Model
	GenericModel[Autor]
	Nome           string    `json:"nome" gorm:"not null" validate:"required,min=2,max=100"`
	DataNascimento time.Time `json:"data_nascimento" gorm:"not null" validate:"required"`
}
