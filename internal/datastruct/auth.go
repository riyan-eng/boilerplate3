package datastruct

type Login struct {
	UserID   string `db:"id" json:"-"`
	Username string `db:"username" json:"username"`
	Email    string `db:"email" json:"email"`
	Password string `db:"password" json:"-"`
	RoleCode string `db:"role_code" json:"-"`
	RoleName string `db:"role_name" json:"role"`
}
