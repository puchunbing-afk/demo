// server/server.go
package server

import (
	"awesomeProject/hello"
	"net"

	"google.golang.org/grpc"
)

func StartServer(port string) error {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()
	hello.RegisterHelloServiceServer(grpcServer, &HelloServiceImpl{})

	return grpcServer.Serve(lis)
}
