package mongo

import (
	"budgeting_service/pkg/logger"
	"context"

	pb "budgeting_service/genproto/budgeting_service"

	"go.mongodb.org/mongo-driver/mongo"
)

type categoryRepo struct {
	db  *mongo.Database
	log logger.ILogger
}

func NewCategoryRepo(db *mongo.Database, log logger.ILogger) *categoryRepo {
	return &categoryRepo{
		db:  db,
		log: log,
	}
}

func (c *categoryRepo) Create(context.Context, *pb.CreateCategory) (*pb.Category, error)
func (c *categoryRepo) GetById(context.Context, *pb.PrimaryKey) (*pb.Category, error)
func (c *categoryRepo) GetAll(context.Context, *pb.CategoryFilter) (*pb.Categories, error)
func (c *categoryRepo) Update(context.Context, *pb.Category) (*pb.Category, error)
func (c *categoryRepo) Delete(context.Context, *pb.PrimaryKey) error
