package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/riyan-eng/boilerplate3/internal/serrepconnector"
	"github.com/riyan-eng/boilerplate3/pkg/util"
)

type AuthQuery interface {
	Login(serrepconnector.AuthLoginReq) *sql.Rows
	Register(serrepconnector.AuthRegisterReq)
	RefreshToken()
	ResetPassword()
}

type authQuery struct {
	db *sql.DB
}

func (a *authQuery) Login(req serrepconnector.AuthLoginReq) *sql.Rows {
	query := fmt.Sprintf(`
		select u.id, u.username, u.email, u.password, g.name as gender, r.code as role_code, 
		r.name as role_name, ud.name as name, ud.address as address
		from public.users u
		left join public.roles r on u.role = r.code
		left join public.user_datas ud on u.user_data = ud.id
		left join public.genders g on ud.gender = g.code
		where u.username = '%v'
	`, req.Username)
	sqlrows, err := a.db.Query(query)
	util.PanicIfNeeded(err)
	return sqlrows
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
