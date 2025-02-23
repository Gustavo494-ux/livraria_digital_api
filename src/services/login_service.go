package services

import (
	"errors"
	"livraria_digital/src/models"
	"livraria_digital/src/repository"
	"livraria_digital/src/utils"
)

type ILoginService interface {
	repository.ILoginRepository
	BuscarUsuarioPorEmail(email string) (usuario models.Usuario, err error)
	RealizarLogin(login models.Login) (token string, err error)
}

type LoginService struct {
	repo *repository.LoginRepository
}

// NewLoginService cria uma nova instância do serviço de login
func NewLoginService(repo *repository.LoginRepository) *LoginService {
	return &LoginService{
		repo: repo,
	}
}

// BuscarUsuario: busca o primeiro usuário que atender aos requisitos
func (r *LoginService) BuscarUsuario(usuarioFiltro models.Usuario) (usuario models.Usuario, err error) {
	return r.repo.BuscarUsuario(usuarioFiltro)
}

// BuscarUsuario: busca o usuário pelo e-mail
func (r *LoginService) BuscarUsuarioPorEmail(email string) (usuario models.Usuario, err error) {
	usuarioFiltro := models.Usuario{}
	usuarioFiltro.Email = email
	return r.BuscarUsuario(usuarioFiltro)
}

// RealizarLogin: verifica as credenciais de login estão corretas
func (r *LoginService) RealizarLogin(login models.Login) (token string, err error) {
	usuario, err := r.BuscarUsuarioPorEmail(login.Email)
	if err != nil {
		return
	}

	if !usuario.IsID() || !utils.ValidarHash(login.Senha, usuario.Senha) {
		err = errors.New("credenciais inválidas")
		return
	}

	if !usuario.IsAtivo() {
		err = errors.New("usuário inativo")
		return
	}

	return usuario.GerarToken()
}
