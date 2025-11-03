// server/service.go
package server

import (
	"awesomeProject/hello"
	"context"
)

type HelloServiceImpl struct {
	hello.UnimplementedHelloServiceServer
}

func (s *HelloServiceImpl) Hello(ctx context.Context, req *hello.String) (*hello.String, error) {
	return &hello.String{Value: "Hello, " + req.GetValue()}, nil
}
