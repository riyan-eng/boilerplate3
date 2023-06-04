package app

import "github.com/riyan-eng/boilerplate3/internal/service"

type ServiceServer struct {
	taskService service.TaskService
	authService service.AuthService
}

func NewService(taskService service.TaskService, authService service.AuthService) *ServiceServer {
	return &ServiceServer{
		taskService: taskService,
		authService: authService,
	}
}
