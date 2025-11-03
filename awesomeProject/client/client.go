// client/client.go
package client

import (
	"awesomeProject/hello"
	"context"
	"time"

	"google.golang.org/grpc"
)

type HelloClient struct {
	client hello.HelloServiceClient
	conn   *grpc.ClientConn
}

func NewHelloClient(address string) (*HelloClient, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return &HelloClient{
		client: hello.NewHelloServiceClient(conn),
		conn:   conn,
	}, nil
}

func (c *HelloClient) CallHello(message string) (*hello.String, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	return c.client.Hello(ctx, &hello.String{Value: message})
}

func (c *HelloClient) Close() error {
	return c.conn.Close()
}
