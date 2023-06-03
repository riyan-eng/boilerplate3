package service

import (
	"github.com/riyan-eng/boilerplate3/internal/dto"
	"github.com/riyan-eng/boilerplate3/internal/repository"
	"github.com/riyan-eng/boilerplate3/internal/serrepconnector"
	srv "github.com/riyan-eng/boilerplate3/pkg"
)

type TaskService interface {
	ListTask()
	CreateTask(dto.TaskCreateReq) dto.TaskCreateRes
	DeleteTask()
	DetailTask()
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

func (t *taskService) ListTask() {
}

func (t *taskService) CreateTask(req dto.TaskCreateReq) (res dto.TaskCreateRes) {
	t.dao.NewTaskQuery().CreateTask(serrepconnector.CreateTaskReq{
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

func (t *taskService) DetailTask() {
}

func (t *taskService) UpdateTask() {
}
