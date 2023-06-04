package service

import "github.com/riyan-eng/boilerplate3/internal/repository"

type AuthService interface {
	Login()
	Register()
	RefreshToken()
	ResetPassword()
}

type authService struct {
	dao repository.DAO
}

func NewAuthService(dao repository.DAO) AuthService {
	return &authService{
		dao: dao,
	}
}

func (a *authService) Login() {
}

func (a *authService) Register() {
}

func (a *authService) RefreshToken() {
}

func (a *authService) ResetPassword() {
}
