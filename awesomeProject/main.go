// main.go
package main

import (
	"awesomeProject/client"
	"awesomeProject/server"
	"fmt"
	"log"
)

func main() {
	// 启动服务端（可以在 goroutine 中运行）
	go func() {
		if err := server.StartServer(":50051"); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// 创建客户端并调用服务
	client, err := client.NewHelloClient("localhost:50051")
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()
	fmt.Println("调用服务")

	response, err := client.CallHello("测试king")
	if err != nil {
		log.Fatalf("Failed to call service: %v", err)
	}

	log.Printf("Response: %s", response.GetValue())
}
