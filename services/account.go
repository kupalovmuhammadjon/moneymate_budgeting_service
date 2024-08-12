package services

import (
	pb "budgeting_service/genproto/budgeting_service"
	"budgeting_service/pkg/logger"
	"budgeting_service/storage"
	"context"
)

type accountService struct {
	storage storage.IStorage
	log     logger.ILogger
	pb.UnimplementedAccountServiceServer
}

func NewAccountService(storage storage.IStorage, log logger.ILogger) *accountService {
	return &accountService{
		storage: storage,
		log:     log,
	}
}

func (s *accountService) Create(ctx context.Context, req *pb.CreateAccount) (*pb.Account, error) {
	s.log.Info("Create account request received", logger.Any("user_id", req.UserId))

	accountID, err := s.storage.Accounts().Create(ctx, req)
	if err != nil {
		s.log.Error("Failed to create account", logger.Error(err))
		return nil, err
	}

	account, err := s.storage.Accounts().GetById(ctx, &pb.PrimaryKey{Id: accountID})
	if err != nil {
		s.log.Error("Failed to get account by ID after creation", logger.Error(err))
		return nil, err
	}

	return account, nil
}

func (s *accountService) GetById(ctx context.Context, req *pb.PrimaryKey) (*pb.Account, error) {
	s.log.Info("Get account by ID request received", logger.Any("id", req.Id))

	account, err := s.storage.Accounts().GetById(ctx, req)
	if err != nil {
		s.log.Error("Failed to get account by ID", logger.Error(err))
		return nil, err
	}

	return account, nil
}

func (s *accountService) GetAll(ctx context.Context, req *pb.AccountFilter) (*pb.Accounts, error) {
	s.log.Info("Get all accounts request received", logger.Any("filter", req))

	accounts, err := s.storage.Accounts().GetAll(ctx, req)
	if err != nil {
		s.log.Error("Failed to get all accounts", logger.Error(err))
		return nil, err
	}

	return accounts, nil
}

func (s *accountService) Update(ctx context.Context, req *pb.Account) (*pb.Account, error) {
	s.log.Info("Update account request received", logger.Any("account", req))

	err := s.storage.Accounts().Update(ctx, req)
	if err != nil {
		s.log.Error("Failed to update account", logger.Error(err))
		return nil, err
	}

	updatedAccount, err := s.storage.Accounts().GetById(ctx, &pb.PrimaryKey{Id: req.Id})
	if err != nil {
		s.log.Error("Failed to get updated account by ID", logger.Error(err))
		return nil, err
	}

	return updatedAccount, nil
}

func (s *accountService) Delete(ctx context.Context, req *pb.PrimaryKey) (*pb.Void, error) {
	s.log.Info("Delete account request received", logger.Any("id", req.Id))

	err := s.storage.Accounts().Delete(ctx, req)
	if err != nil {
		s.log.Error("Failed to delete account", logger.Error(err))
		return nil, err
	}

	return &pb.Void{}, nil
}
