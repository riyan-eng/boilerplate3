package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/riyan-eng/boilerplate3/internal/serrepconnector"
	"github.com/riyan-eng/boilerplate3/pkg/util"
)

type AuthQuery interface {
	Login()
	Register(serrepconnector.AuthRegisterReq)
	RefreshToken()
	ResetPassword()
}

type authQuery struct {
	db *sql.DB
}

func (a *authQuery) Login() {
}

func (a *authQuery) Register(req serrepconnector.AuthRegisterReq) {
	queryUserData := fmt.Sprintf(`
		insert into public.user_datas (id, name, gender, address) values ('%v', '%v', '%v', '%v')
	`, req.UserDataID, req.Name, req.Gender, req.Address)
	queryUser := fmt.Sprintf(`
	insert into public.users (id, username, email, password, role, phone, user_data) values ('%v', '%v', '%v', '%v', '%v', nullif('%v', ''), '%v')
	`, req.UserID, req.Username, req.Email, req.Password, req.RoleCode, req.Phone, req.UserDataID)

	ctx := context.Background()
	tx, err := a.db.BeginTx(ctx, nil)
	util.PanicIfNeeded(err)
	_, err = tx.ExecContext(ctx, queryUserData)
	util.PanicIfNeeded(err)
	_, err = tx.ExecContext(ctx, queryUser)
	err = tx.Commit()
	util.PanicIfNeeded(err)
}

func (a *authQuery) RefreshToken() {
}

func (a *authQuery) ResetPassword() {
}
