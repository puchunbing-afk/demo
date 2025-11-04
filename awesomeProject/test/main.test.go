// test_client.go
package main

import (
	"context"
	"log"
	"time"

	pb "git.huoys.com/oversea/grpc/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewUserClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// 测试 GetUserVipTag
	//resp, err := client.GetUserVipTag(ctx, &pb.GetUserVipTagRequest{
	//	PlayerId: 1001375,
	//})

	//resp, err := client.BathGetUserVipTags(ctx, &pb.BathGetUserVipTagRequest{
	//	PlayerIds: []int64{1001375, 1089397, 1101206},
	//})

	// 测试 GetPlayerInfosByIdFieldReform
	//resp, err := client.GetPlayerInfosByIdFieldReform(ctx, &pb.GetPlayerInfosByIdFieldRequest{
	//	PlayerIds: []int64{1001375, 1089397, 1101206},
	//	Fields:    []string{"Nick", "Sex", "OS"},
	//})

	resp, err := client.GetUserVipBlueTagIds(ctx, &pb.GetUserVipBlueTagIdsRequest{
		Sum: 5,
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Response: %v", resp)
}
