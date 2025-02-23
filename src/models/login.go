package models

import (
	"livraria_digital/src/utils"

	"gorm.io/gorm"
)

type Login struct {
	gorm.Model
	Email string `json:"email" validate:"required,email"`
	Senha string `json:"senha" validate:"required,min=6"`
}

// EnriptarSenha: gera um hash seguro para a senha
func (login *Login) EnriptarSenha() {
	login.Senha = utils.GerarHash(login.Senha)
}
