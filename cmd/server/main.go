package main

import (
	"log"

	"github.com/zeddmathews/tasksync/internal/server"
)
// type TaskServiceServer struct {
//     taskpb.UnimplementedTaskServiceServer
// }
func main()  {
	serv := server.NewGRPCServer()

	go func() {
		log.Printf("Server listening on port: %v", server.GRPCPort)
		if err := serv.Start(); err != nil {
			log.Fatalf("Server error: \n%v", err)
		}
	}()
	server.ShutdownProcess(serv)
}
