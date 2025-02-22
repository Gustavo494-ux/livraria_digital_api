package models

import (
	"time"

	"gorm.io/gorm"
)

type Usuario struct {
	gorm.Model
	GenericModel[Usuario]
	Nome        string    `json:"nome" gorm:"not null" validate:"required,min=2,max=100"`
	Email       string    `json:"email" gorm:"unique;not null" validate:"required,email"`
	Senha       string    `json:"senha" gorm:"not null" validate:"required,min=6"`
	Ativo       bool      `json:"ativo" gorm:"not null" validate:"required"`
	UltimoLogin time.Time `json:"ultimo_login"`
}

// IsID: verifica se o Id é válido. ID != 0
func (usuario *Usuario) IsID() bool {
	return usuario.ID != 0
}

// IsAtivo: verifica se o usuário está ativo
func (usuario *Usuario) IsAtivo() bool {
	return usuario.Ativo
}

// IsEmailValido: verifica se o email é válido
func (usuario *Usuario) IsEmailValido() bool {
	return usuario.Email != ""
}

// IsSenhaValida: verifica se a senha tem pelo menos 6 caracteres
func (usuario *Usuario) IsSenhaValida() bool {
	return len(usuario.Senha) >= 6
}
