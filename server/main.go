package main

import (
	// "fmt"
	"log"
	"net"
	"google.golang.org/grpc"
	taskpb "github.com/zeddmathews/tasksync/proto"
)
type TaskServiceServer struct {
    taskpb.UnimplementedTaskServiceServer
}
func main()  {
	port := ":50051"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := *grpc.NewServer()
	taskpb.RegisterTaskServiceServer(&grpcServer, &TaskServiceServer{})

	log.Printf("Server listening on port %v", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
