package serrepconnector

type TaskCreateReq struct {
	ID     string
	Name   string
	Detail string
}

type TaskCreateRes struct{}

type TaskListReq struct {
	Search string
	Limit  uint32
	Offset uint32
	Order  string
}

type TaskListRes struct{}

type TaskDetailReq struct {
	ID string
}

type TaskDeleteReq struct {
	ID string
}

type TaskUpdateReq struct {
	ID     string
	Name   string
	Detail string
}
