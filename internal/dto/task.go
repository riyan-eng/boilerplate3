package dto

import (
	"github.com/riyan-eng/boilerplate3/internal/datastruct"
	srv "github.com/riyan-eng/boilerplate3/pkg"
)

type TaskCreateReq struct {
	ID     string
	Name   string
	Detail string
}

type TaskCreateRes struct {
	Data srv.CreateTaskRes
}

type TaskListReq struct {
	Search string
	Limit  uint32
	Page   uint32
	Order  string
}

type TaskListRes struct {
	Items []datastruct.Task
	Page  uint32
	Limit uint32
	Total uint32
}

type TaskDetailReq struct {
	ID string
}

type TaskDetailRes struct {
	Item datastruct.Task
}

type TaskDeleteReq struct {
	ID string
}
