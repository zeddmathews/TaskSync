package main

import (
	"context"
	"fmt"
	"log"
	"time"

	taskpb "github.com/zeddmathews/tasksync/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	// grpc.Dial("localhost:50051", grpc.WithInsecure()) **deprecated
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := taskpb.NewTaskServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &taskpb.TaskRequest{
		Title: "Test task",
		Description: "Instance of task testing",
	}

	res, err := client.CreateTask(ctx, req)
	if err != nil {
		log.Fatalf("Error creating task: %v", err)
	}

	log.Printf("Task created. ID: %s, Status: %s\n", res.Id, res.Status)

	fmt.Println("gRPC client running...")
}