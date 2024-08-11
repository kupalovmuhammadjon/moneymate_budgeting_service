package services

import (
	"context"
	pb "budgeting_service/genproto/budgeting_service"
	"budgeting_service/pkg/logger"
	"budgeting_service/storage"
)

type budgetService struct {
	storage storage.IStorage
	log     logger.ILogger
	pb.UnimplementedBudgetServiceServer
}

func NewBudgetService(storage storage.IStorage, log logger.ILogger) *budgetService {
	return &budgetService{
		storage: storage,
		log:     log,
	}
}

func (s *budgetService) Create(ctx context.Context, req *pb.CreateBudget) (*pb.Budget, error) {
	s.log.Info("Create budget request received", logger.Any("user_id", req.UserId))

	budget, err := s.storage.Budgets().Create(ctx, req)
	if err != nil {
		s.log.Error("Failed to create budget", logger.Error(err))
		return nil, err
	}


	return budget, nil
}

func (s *budgetService) GetById(ctx context.Context, req *pb.PrimaryKey) (*pb.Budget, error) {
	s.log.Info("Get budget by ID request received", logger.Any("id", req.Id))

	budget, err := s.storage.Budgets().GetById(ctx, req)
	if err != nil {
		s.log.Error("Failed to get budget by ID", logger.Error(err))
		return nil, err
	}

	return budget, nil
}

func (s *budgetService) GetAll(ctx context.Context, req *pb.BudgetFilter) (*pb.Budgets, error) {
	s.log.Info("Get all budgets request received", logger.Any("filter", req))

	budgets, err := s.storage.Budgets().GetAll(ctx, req)
	if err != nil {
		s.log.Error("Failed to get all budgets", logger.Error(err))
		return nil, err
	}

	return budgets, nil
}

func (s *budgetService) Update(ctx context.Context, req *pb.Budget) (*pb.Budget, error) {
	s.log.Info("Update budget request received", logger.Any("budget", req))

	_, err := s.storage.Budgets().Update(ctx, req)
	if err != nil {
		s.log.Error("Failed to update budget", logger.Error(err))
		return nil, err
	}

	updatedBudget, err := s.storage.Budgets().GetById(ctx, &pb.PrimaryKey{Id: req.Id})
	if err != nil {
		s.log.Error("Failed to get updated budget by ID", logger.Error(err))
		return nil, err
	}

	return updatedBudget, nil
}

func (s *budgetService) Delete(ctx context.Context, req *pb.PrimaryKey) (*pb.Void, error) {
	s.log.Info("Delete budget request received", logger.Any("id", req.Id))

	err := s.storage.Budgets().Delete(ctx, req)
	if err != nil {
		s.log.Error("Failed to delete budget", logger.Error(err))
		return nil, err
	}

	return &pb.Void{}, nil
}
