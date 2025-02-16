package models

import (
	"gorm.io/gorm"
)

type Categoria struct {
	gorm.Model
	GenericModel[Categoria]
	Nome string `json:"nome" gorm:"not null" validate:"required,min=2,max=100"`
}
