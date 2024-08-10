package mongo

import (
	"budgeting_service/pkg/logger"
	"context"

	pb "budgeting_service/genproto/budgeting_service"

	"go.mongodb.org/mongo-driver/mongo"
)

type budgetRepo struct {
	db  *mongo.Database
	log logger.ILogger
}

func NewBudgetRepo(db *mongo.Database, log logger.ILogger) *budgetRepo {
	return &budgetRepo{
		db:  db,
		log: log,
	}
}

func (b *budgetRepo) Create(context.Context, *pb.CreateBudget) (*pb.Budget, error)
func (b *budgetRepo) GetById(context.Context, *pb.PrimaryKey) (*pb.Budget, error)
func (b *budgetRepo) GetAll(context.Context, *pb.BudgetFilter) (*pb.Budgets, error)
func (b *budgetRepo) Update(context.Context, *pb.Budget) (*pb.Budget, error)
func (b *budgetRepo) Delete(context.Context, *pb.PrimaryKey) error
