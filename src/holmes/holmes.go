package holmes

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"holmes/conf"
	"log"
	"net"
	"register"
)

// Run is grpc server start
func Fly()  {
	lis, err := net.Listen("tcp", conf.Conf.Server.Host + ":" + conf.Conf.Server.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	register.RegService(s)

	// Register reflection service on gRPC server.
	reflection.Register(s)
	log.Printf("holmes server start listen on: %s:%s\n", conf.Conf.Server.Host, conf.Conf.Server.Port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
