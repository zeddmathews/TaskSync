package server

import (
	"log"
	"net"
	"google.golang.org/grpc"
)

type GRPCServer struct {
	grpcServer *grpc.Server
	lis net.Listener
}

func NewGRPCServer() (gs *GRPCServer) {
	if err := InitDB(); err != nil {
		log.Fatalf("DB init failed: \n %v", err)
	}

	lis, err := net.Listen("tcp", GRPCPort)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	serv := grpc.NewServer()
	RegisterServices(serv)
	gs = &GRPCServer{
		grpcServer: serv,
		lis: lis,
	}
	return
}

func (serv *GRPCServer) Start() error {
	return serv.grpcServer.Serve(serv.lis)
}

func (serv *GRPCServer) Stop() {
	log.Println("Shutting down gRPC server")
	serv.grpcServer.GracefulStop()
}