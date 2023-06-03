package pkg

type CreateTaskReq struct {
	Name   string `json:"name"`
	Detail string `json:"detail"`
}

type ListTaskReq struct {
	Page   uint32 `query:"page" default:"1"`
	Limit  uint32 `query:"limit" default:"10"`
	Search string `query:"search" default:"mamby"`
	Order  string `query:"order"`
}

func (l ListTaskReq) Init() ListTaskReq {
	l.Page = 1
	l.Limit = 10
	l.Order = "desc"
	return l
}
