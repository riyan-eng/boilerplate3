package dto

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/riyan-eng/boilerplate3/internal/datastruct"
	srv "github.com/riyan-eng/boilerplate3/pkg"
)

type AuthRegisterReq struct {
	UserID     string
	UserDataID string
	Username   string
	Email      string
	Password   string
	Name       string
	Gender     string
	Address    string
	RoleCode   string
	Phone      string
}

type AuthRegisterRes struct {
	Data srv.AuthRegisterRes
}

type AuthLoginReq struct {
	Username string
	Password string
}

type AuthLoginRes struct {
	Data         datastruct.Login
	AccessToken  string
	RefreshToken string
	ExpiredAt    *jwt.NumericDate
	Match        bool
}
