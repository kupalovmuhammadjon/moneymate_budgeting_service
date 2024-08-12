package services

import (
	"context"
	pb "budgeting_service/genproto/budgeting_service"
	"budgeting_service/pkg/logger"
	"budgeting_service/storage"
)

type categoryService struct {
	storage storage.IStorage
	log     logger.ILogger
	pb.UnimplementedCategoryServiceServer
}

func NewCategoryService(storage storage.IStorage, log logger.ILogger) *categoryService {
	return &categoryService{
		storage: storage,
		log:     log,
	}
}

func (s *categoryService) Create(ctx context.Context, req *pb.CreateCategory) (*pb.Category, error) {
	s.log.Info("Create category request received", logger.Any("user_id", req.UserId))

	category, err := s.storage.Categories().Create(ctx, req)
	if err != nil {
		s.log.Error("Failed to create category", logger.Error(err))
		return nil, err
	}

	return category, nil
}

func (s *categoryService) GetById(ctx context.Context, req *pb.PrimaryKey) (*pb.Category, error) {
	s.log.Info("Get category by ID request received", logger.Any("id", req.Id))

	category, err := s.storage.Categories().GetById(ctx, req)
	if err != nil {
		s.log.Error("Failed to get category by ID", logger.Error(err))
		return nil, err
	}

	return category, nil
}

func (s *categoryService) GetAll(ctx context.Context, req *pb.CategoryFilter) (*pb.Categories, error) {
	s.log.Info("Get all categories request received", logger.Any("filter", req))

	categories, err := s.storage.Categories().GetAll(ctx, req)
	if err != nil {
		s.log.Error("Failed to get all categories", logger.Error(err))
		return nil, err
	}

	return categories, nil
}

func (s *categoryService) Update(ctx context.Context, req *pb.Category) (*pb.Category, error) {
	s.log.Info("Update category request received", logger.Any("category", req))

	_, err := s.storage.Categories().Update(ctx, req)
	if err != nil {
		s.log.Error("Failed to update category", logger.Error(err))
		return nil, err
	}

	updatedCategory, err := s.storage.Categories().GetById(ctx, &pb.PrimaryKey{Id: req.Id})
	if err != nil {
		s.log.Error("Failed to get updated category by ID", logger.Error(err))
		return nil, err
	}

	return updatedCategory, nil
}

func (s *categoryService) Delete(ctx context.Context, req *pb.PrimaryKey) (*pb.Void, error) {
	s.log.Info("Delete category request received", logger.Any("id", req.Id))

	err := s.storage.Categories().Delete(ctx, req)
	if err != nil {
		s.log.Error("Failed to delete category", logger.Error(err))
		return nil, err
	}

	return &pb.Void{}, nil
}
