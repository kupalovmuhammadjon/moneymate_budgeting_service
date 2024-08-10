package main

import (
	"budgeting_service/configs"
	"budgeting_service/grpc"
	"budgeting_service/pkg/logger"
	"budgeting_service/storage"
	"context"

	"fmt"
	"net"
)

func main() {
	cfg := configs.Load()

	log := logger.NewLogger(cfg.ServiceName, cfg.LoggerLevel, cfg.LogPath)
	defer logger.Cleanup(log)

	storage, err := storage.NewIStorage(context.Background(), cfg, log)
	if err != nil {
		log.Panic("error while creating storage in main", logger.Error(err))
		return
	}
	defer storage.Close()

	listener, err := net.Listen("tcp",
		cfg.UserServiceGrpcHost+cfg.UserServiceGrpcPort,
	)
	if err != nil {
		log.Panic("error while creating listener for user service", logger.Error(err))
		return
	}
	defer listener.Close()

	server := grpc.SetUpServer(storage, log)

	fmt.Printf("User service is listening on port %s...\n",
		cfg.UserServiceGrpcHost+cfg.UserServiceGrpcPort)
	if err := server.Serve(listener); err != nil {
		log.Fatal("Error with listening user server", logger.Error(err))
	}
}
