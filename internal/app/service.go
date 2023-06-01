package app

type ServiceServer struct{}

func NewService() *ServiceServer {
    return &ServiceServer{}
}