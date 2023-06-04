package service

import (
	"github.com/blockloop/scan/v2"
	"github.com/riyan-eng/boilerplate3/internal/dto"
	"github.com/riyan-eng/boilerplate3/internal/repository"
	"github.com/riyan-eng/boilerplate3/internal/serrepconnector"
	srv "github.com/riyan-eng/boilerplate3/pkg"
	"github.com/riyan-eng/boilerplate3/pkg/util"
)

type TaskService interface {
	ListTask(dto.TaskListReq) dto.TaskListRes
	CreateTask(dto.TaskCreateReq) dto.TaskCreateRes
	DeleteTask()
	DetailTask(dto.TaskDetailReq) dto.TaskDetailRes
	UpdateTask()
}

type taskService struct {
	dao repository.DAO
}

func NewTaskService(dao repository.DAO) TaskService {
	return &taskService{
		dao: dao,
	}
}

func (t *taskService) ListTask(req dto.TaskListReq) (res dto.TaskListRes) {
	pageMeta := util.PageMeta(req.Page, req.Limit)
	sqlrows := t.dao.NewTaskQuery().ListTask(serrepconnector.TaskListReq{
		Search: req.Search,
		Limit:  pageMeta.Limit,
		Offset: pageMeta.Offset,
		Order:  req.Order,
	})
	err := scan.Rows(&res.Items, sqlrows)
	util.PanicIfNeeded(err)
	res.Page = pageMeta.Page
	res.Limit = pageMeta.Page
	if len(res.Items) > 0 {
		res.Total = res.Items[0].Total
	}
	return
}

func (t *taskService) CreateTask(req dto.TaskCreateReq) (res dto.TaskCreateRes) {
	t.dao.NewTaskQuery().CreateTask(serrepconnector.TaskCreateReq{
		ID:     req.ID,
		Name:   req.Name,
		Detail: req.Detail,
	})
	res.Data = srv.CreateTaskRes{
		ID:     req.ID,
		Name:   req.Name,
		Detail: req.Detail,
	}
	return
}

func (t *taskService) DeleteTask() {
}

func (t *taskService) DetailTask(req dto.TaskDetailReq) (res dto.TaskDetailRes) {
	sqlrows := t.dao.NewTaskQuery().DetailTask(serrepconnector.TaskDetailReq{
		ID: req.ID,
	})
	err := scan.Row(&res.Item, sqlrows)
	util.PanicIfNeeded(err)
	return
}

func (t *taskService) UpdateTask() {
}
