package server

import (
	"os"
	"os/signal"
	"syscall"
)

func ShutdownProcess(s *GRPCServer) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	s.Stop()
	if DB != nil {
		DB.Close()
	}
}