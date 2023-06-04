package dto

import (
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
