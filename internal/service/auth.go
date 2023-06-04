package service

import (
	"github.com/riyan-eng/boilerplate3/internal/dto"
	"github.com/riyan-eng/boilerplate3/internal/repository"
	"github.com/riyan-eng/boilerplate3/internal/serrepconnector"
	srv "github.com/riyan-eng/boilerplate3/pkg"
	"github.com/riyan-eng/boilerplate3/pkg/util"
)

type AuthService interface {
	Login()
	Register(dto.AuthRegisterReq) dto.AuthRegisterRes
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

func (a *authService) Register(req dto.AuthRegisterReq) (res dto.AuthRegisterRes) {
	a.dao.NewAuthQuery().Register(serrepconnector.AuthRegisterReq{
		UserID:     req.UserID,
		UserDataID: req.UserDataID,
		Username:   req.Username,
		Email:      req.Email,
		Password:   util.GenerateHash(req.Password),
		Name:       req.Name,
		Gender:     req.Gender,
		Address:    req.Address,
		RoleCode:   req.RoleCode,
		Phone:      req.Phone,
	})
	res.Data = srv.AuthRegisterRes{
		Username: req.Username,
		Email:    req.Email,
		RoleCode: req.RoleCode,
	}
	return
}

func (a *authService) RefreshToken() {
}

func (a *authService) ResetPassword() {
}
