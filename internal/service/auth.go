package service

import (
	"database/sql"

	"github.com/blockloop/scan/v2"
	"github.com/riyan-eng/boilerplate3/internal/dto"
	"github.com/riyan-eng/boilerplate3/internal/repository"
	"github.com/riyan-eng/boilerplate3/internal/serrepconnector"
	srv "github.com/riyan-eng/boilerplate3/pkg"
	"github.com/riyan-eng/boilerplate3/pkg/util"
)

type AuthService interface {
	Login(dto.AuthLoginReq) dto.AuthLoginRes
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

func (a *authService) Login(req dto.AuthLoginReq) (res dto.AuthLoginRes) {
	sqlrows := a.dao.NewAuthQuery().Login(serrepconnector.AuthLoginReq{
		Username: req.Username,
	})
	err := scan.Row(&res.Data, sqlrows)
	if err == sql.ErrNoRows {
		return
	} else {
		util.PanicIfNeeded(err)
	}

	ok := util.VerifyHash(res.Data.Password, req.Password)
	if ok {
		res.Match = true
		genJwt := util.GenerateJwt(res.Data.UserID, res.Data.RoleCode, "issuer")
		res.AccessToken = genJwt.AccessToken
		res.RefreshToken = genJwt.RefreshToken
		res.ExpiredAt = genJwt.ExpiredAt
		return
	}
	return
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
