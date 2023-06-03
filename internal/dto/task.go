package dto

import srv "github.com/riyan-eng/boilerplate3/pkg"

type TaskCreateReq struct {
	ID     string
	Name   string
	Detail string
}

type TaskCreateRes struct {
	Data srv.CreateTaskRes
}
