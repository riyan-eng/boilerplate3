package repository

import "database/sql"

type AuthQuery interface {
	Login()
	Register()
	RefreshToken()
	ResetPassword()
}

type authQuery struct {
	db *sql.DB
}

func (a *authQuery) Login() {
}

func (a *authQuery) Register() {
}

func (a *authQuery) RefreshToken() {
}

func (a *authQuery) ResetPassword() {
}
