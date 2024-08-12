package services

import (
	"context"

	pb "budgeting_service/genproto/budgeting_service"
	"budgeting_service/pkg/logger"
	"budgeting_service/storage"
)

type transactionService struct {
	storage storage.IStorage
	log     logger.ILogger
	pb.UnimplementedTransactionServiceServer
}

func NewTransactionService(storage storage.IStorage, log logger.ILogger) *transactionService {
	return &transactionService{
		storage: storage,
		log:     log,
	}
}

func (s *transactionService) Create(ctx context.Context, req *pb.CreateTransaction) (*pb.Transaction, error) {

	result, err := s.storage.Transactions().Create(ctx, req)
	if err != nil {
		s.log.Error("Error while creating transaction in service layer", logger.Error(err))
		return nil, err
	}

	return result, nil
}

func (s *transactionService) GetById(ctx context.Context, req *pb.PrimaryKey) (*pb.Transaction, error) {
	result, err := s.storage.Transactions().GetById(ctx, req)
	if err != nil {
		s.log.Error("Error while getting transaction by ID in service layer", logger.Error(err))
		return nil, err
	}

	return result, nil
}

func (s *transactionService) GetAll(ctx context.Context, req *pb.TransactionFilter) (*pb.Transactions, error) {
	result, err := s.storage.Transactions().GetAll(ctx, req)
	if err != nil {
		s.log.Error("Error while getting all transactions in service layer", logger.Error(err))
		return nil, err
	}

	return result, nil
}

func (s *transactionService) Update(ctx context.Context, req *pb.Transaction) (*pb.Transaction, error) {

	result, err := s.storage.Transactions().Update(ctx, req)
	if err != nil {
		s.log.Error("Error while updating transaction in service layer", logger.Error(err))
		return nil, err
	}

	return result, nil
}

func (s *transactionService) Delete(ctx context.Context, req *pb.PrimaryKey) (*pb.Void, error) {
	err := s.storage.Transactions().Delete(ctx, req)
	if err != nil {
		s.log.Error("Error while deleting transaction in service layer", logger.Error(err))
		return nil, err
	}

	return &pb.Void{}, nil
}
