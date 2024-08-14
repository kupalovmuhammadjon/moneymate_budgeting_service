package main

import (
	"budgeting_service/configs"
	"budgeting_service/grpc"
	"budgeting_service/pkg/logger"
	"budgeting_service/pkg/messege_brokers/kafka"
	"budgeting_service/services"
	"budgeting_service/storage"
	"context"

	"fmt"
	"net"

	"go.uber.org/zap"
)

func main() {
	cfg := configs.Load()

	log := logger.NewLogger(cfg.ServiceName, cfg.LoggerLevel, cfg.LogPath)
	defer logger.Cleanup(log)

	storage, err := storage.NewIStorage(context.Background(), cfg, log)
	if err != nil {
		log.Panic("error while creating storage ", logger.Error(err))
		return
	}
	defer storage.Close()

	iServiceManager := services.NewIServiceManager(storage, log)

	iKafka, err := kafka.NewIKafka()
	if err != nil {
		log.Fatal("Failed to create Kafka producer and consumer", zap.Error(err))
		return
	}
	defer iKafka.Close()

	go iKafka.ConsumeMessages([]string{"transaction_created", "budget_updated", "goal_progress_updated", 
	"notification_created"}, iServiceManager)

	listener, err := net.Listen("tcp",
		cfg.BudgetingServiceGrpcPort,
	)
	if err != nil {
		log.Panic("error while creating listener for budgeting service", logger.Error(err))
		return
	}
	defer listener.Close()

	server := grpc.SetUpServer(iServiceManager, storage, log)

	fmt.Printf("Budgeting service is listening on port %s...\n",
		cfg.BudgetingServiceGrpcHost+cfg.BudgetingServiceGrpcPort)
	if err := server.Serve(listener); err != nil {
		log.Fatal("Error with listening budgeting server", logger.Error(err))
	}
}
