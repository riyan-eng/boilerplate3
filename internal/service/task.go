package service

import "github.com/riyan-eng/boilerplate3/internal/repository"

type TaskService interface {
	ListTask()
	CreateTask()
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

func (t *taskService) CreateTask() {
}

func (t *taskService) DeleteTask() {
}

func (t *taskService) DetailTask() {
}

func (t *taskService) UpdateTask() {
}
