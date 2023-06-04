package pkg

type TaskCreateReq struct {
	Name   string `json:"name"`
	Detail string `json:"detail"`
}

type TaskListReq struct {
	Page   uint32 `query:"page" default:"1"`
	Limit  uint32 `query:"limit" default:"10"`
	Search string `query:"search" default:"mamby"`
	Order  string `query:"order"`
}

func (l TaskListReq) Init() TaskListReq {
	l.Page = 1
	l.Limit = 10
	l.Order = "desc"
	return l
}

type TaskUpdateReq struct {
	Name   string `json:"name"`
	Detail string `json:"detail"`
}

type AuthRegister struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Gender   string `json:"gender"`
	Address  string `json:"address"`
}
