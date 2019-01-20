package service

import (
	"context"
	"log"
	"user"
)


// GreeterService is used to implement GreeterService.
type GreeterService struct{}

// SayHello implements helloworld.GreeterServer
func (s *GreeterService) SayHello(ctx context.Context, in *user.HelloRequest) (*user.HelloReply, error) {
	log.Printf("Received: %v", in.Name)
	return &user.HelloReply{Message: "Hello " + in.Name}, nil
}
