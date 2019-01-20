package register

import (
	"google.golang.org/grpc"
	"service"
	"user"
)

//register service

// RegService is for register to register service
func RegService(server *grpc.Server) {
	user.RegisterGreeterServer(server, &service.GreeterService{})

	user.RegisterPropertyServer(server, &service.PropertyService{})
}
