package serrepconnector

type CreateTaskReq struct {
	UserID string
	ID     string
	Name   string
	Detail string
}

type CreateTaskRes struct{}

type ListTaskReq struct {
	Search string
	Limit  uint32
	Offset uint32
	Order  string
}

type ListTaskRes struct{}
