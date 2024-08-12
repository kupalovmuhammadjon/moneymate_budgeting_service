package services

import (
	"context"
	pb "budgeting_service/genproto/budgeting_service"
	"budgeting_service/pkg/logger"
	"budgeting_service/storage"
)

type goalService struct {
	storage storage.IStorage
	log     logger.ILogger
	pb.UnimplementedGoalServiceServer
}

func NewGoalService(storage storage.IStorage, log logger.ILogger) *goalService {
	return &goalService{
		storage: storage,
		log:     log,
	}
}

func (s *goalService) Create(ctx context.Context, req *pb.CreateGoal) (*pb.Goal, error) {
	s.log.Info("Create goal request received", logger.Any("user_id", req.UserId))

	goal, err := s.storage.Goals().Create(ctx, req)
	if err != nil {
		s.log.Error("Failed to create goal", logger.Error(err))
		return nil, err
	}

	return goal, nil
}

func (s *goalService) GetById(ctx context.Context, req *pb.PrimaryKey) (*pb.Goal, error) {
	s.log.Info("Get goal by ID request received", logger.Any("id", req.Id))

	goal, err := s.storage.Goals().GetById(ctx, req)
	if err != nil {
		s.log.Error("Failed to get goal by ID", logger.Error(err))
		return nil, err
	}

	return goal, nil
}

func (s *goalService) GetAll(ctx context.Context, req *pb.GoalFilter) (*pb.Goals, error) {
	s.log.Info("Get all goals request received", logger.Any("filter", req))

	goals, err := s.storage.Goals().GetAll(ctx, req)
	if err != nil {
		s.log.Error("Failed to get all goals", logger.Error(err))
		return nil, err
	}

	return goals, nil
}

func (s *goalService) Update(ctx context.Context, req *pb.Goal) (*pb.Goal, error) {
	s.log.Info("Update goal request received", logger.Any("goal", req))

	_, err := s.storage.Goals().Update(ctx, req)
	if err != nil {
		s.log.Error("Failed to update goal", logger.Error(err))
		return nil, err
	}

	updatedGoal, err := s.storage.Goals().GetById(ctx, &pb.PrimaryKey{Id: req.Id})
	if err != nil {
		s.log.Error("Failed to get updated goal by ID", logger.Error(err))
		return nil, err
	}

	return updatedGoal, nil
}

func (s *goalService) Delete(ctx context.Context, req *pb.PrimaryKey) (*pb.Void, error) {
	s.log.Info("Delete goal request received", logger.Any("id", req.Id))

	err := s.storage.Goals().Delete(ctx, req)
	if err != nil {
		s.log.Error("Failed to delete goal", logger.Error(err))
		return nil, err
	}

	return &pb.Void{}, nil
}
