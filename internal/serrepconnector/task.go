package serrepconnector

type CreateTaskReq struct {
	UserID string
	ID     string
	Name   string
	Detail string
}

type CreateTaskRes struct{}
