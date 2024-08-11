package grpc

import (
	"budgeting_service/pkg/logger"
	"budgeting_service/services"
	pb "budgeting_service/genproto/budgeting_service"
	"budgeting_service/storage"

	"google.golang.org/grpc"
)

func SetUpServer(iServiceManager services.IServiceManager, storage storage.IStorage, log logger.ILogger) *grpc.Server {
	grpcServer := grpc.NewServer()

	pb.RegisterAccountServiceServer(grpcServer, iServiceManager.AccountService())
	pb.RegisterBudgetServiceServer(grpcServer, iServiceManager.BudgetService())
	pb.RegisterCategoryServiceServer(grpcServer, iServiceManager.CategoryService())
	pb.RegisterGoalServiceServer(grpcServer, iServiceManager.GoalService())
	pb.RegisterTransactionServiceServer(grpcServer, iServiceManager.TransactionService())
	
	return grpcServer
}
