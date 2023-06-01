package app

import "github.com/riyan-eng/boilerplate3/internal/service"

type ServiceServer struct {
	taskService service.TaskService
}

func NewService(taskService service.TaskService) *ServiceServer {
	return &ServiceServer{
		taskService: taskService,
	}
}
