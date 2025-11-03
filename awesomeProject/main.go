// main.go
package main

import (
	"awesomeProject/client"
	"awesomeProject/server"
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

	response, err := client.CallHello("测试king")
	if err != nil {
		log.Fatalf("Failed to call service: %v", err)
	}

	log.Printf("Response: %s", response.GetValue())
}

// BatchGetPlayerVipTags 批量获取玩家VIP标签
func (p *playerRepo) BatchGetPlayerVipTags(playerIds []int64) (tagsMap map[int64]common.HPlayerVipTags, err error) {
	// 复用批量查询逻辑
	playerInfos, err := p.GetPlayerInfosByIdField(context.Background(), playerIds, []string{"VipTags"})
	if err != nil {
		return nil, err
	}
}

// GetPlayerVipTag 获取单个玩家VIP标签（基于批量方法）
func (p *playerRepo) GetPlayerVipTag(playerId int64) (tags common.HPlayerVipTags, err error) {
	tagsMap, err := p.BatchGetPlayerVipTags([]int64{playerId})
	if err != nil {
		return common.HPlayerVipTagsUnKnow, err
	}

	if tag, exists := tagsMap[playerId]; exists {
		return tag, nil
	}

	return common.HPlayerVipTagsUnKnow, nil
}
