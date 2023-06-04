package pkg

type TaskCreateRes struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Detail string `json:"detail"`
}

type AuthRegisterRes struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	RoleCode string `json:"role_code"`
}
