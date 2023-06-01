package service

type TaskService interface {
	ListTask()
	CreateTask()
	DeleteTask()
	DetailTask()
	UpdateTask()
}

type taskService struct{}

func NewTaskService() TaskService {
	return &taskService{}
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
