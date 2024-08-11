package services

import (
	pb "budgeting_service/genproto/budgeting_service"
	"budgeting_service/pkg/logger"
	"budgeting_service/storage"
)

type IServiceManager interface {
	AccountService() pb.AccountServiceServer
	BudgetService() pb.BudgetServiceServer
	CategoryService() pb.CategoryServiceServer
	GoalService() pb.GoalServiceServer
	TransactionService() pb.TransactionServiceServer
}

type serviceManager struct {
	accountService     pb.AccountServiceServer
	budgetService      pb.BudgetServiceServer
	categoryService    pb.CategoryServiceServer
	goalService        pb.GoalServiceServer
	transactionService pb.TransactionServiceServer
}

func NewIServiceManager(storage storage.IStorage, log logger.ILogger) IServiceManager {

	return &serviceManager{}
}

func (s *serviceManager) AccountService() pb.AccountServiceServer{
	return s.accountService
}

func (s *serviceManager) BudgetService() pb.BudgetServiceServer{
	return s.budgetService
}

func (s *serviceManager) CategoryService() pb.CategoryServiceServer{
	return s.categoryService
}

func (s *serviceManager) GoalService() pb.GoalServiceServer{
	return s.goalService
}

func (s *serviceManager) TransactionService() pb.TransactionServiceServer{
	return s.transactionService
}
