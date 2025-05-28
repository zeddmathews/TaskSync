package main

import (
	// "fmt"
	"log"
	"net"
	"google.golang.org/grpc"
	// taskpb "github.com/zeddmathews/tasksync/proto"
	"github.com/zeddmathews/tasksync/internal/server"
)
// type TaskServiceServer struct {
//     taskpb.UnimplementedTaskServiceServer
// }
func main()  {
	const port = ":50051"

	if err := server.InitDB(); err != nil {
		log.Fatalf("DB connection failed:\n %v", err)
	}
	defer server.DB.Close()
	log.Println("DB connection success")

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := *grpc.NewServer()
	server.RegisterServices(&grpcServer)

	log.Printf("Server listening on port %v", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
