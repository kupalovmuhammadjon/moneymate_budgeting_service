package mongo

import (
	"budgeting_service/pkg/logger"
	"context"

	pb "budgeting_service/genproto/budgeting_service"

	"go.mongodb.org/mongo-driver/mongo"
)

type goalRepo struct {
	db  *mongo.Database
	log logger.ILogger
}

func NewGoalRepo(db *mongo.Database, log logger.ILogger) *goalRepo {
	return &goalRepo{
		db:  db,
		log: log,
	}
}

func (g *goalRepo) Create(context.Context, *pb.CreateGoal) (*pb.Goal, error)
func (g *goalRepo) GetById(context.Context, *pb.PrimaryKey) (*pb.Goal, error)
func (g *goalRepo) GetAll(context.Context, *pb.GoalFilter) (*pb.Goals, error)
func (g *goalRepo) Update(context.Context, *pb.Goal) (*pb.Goal, error)
func (g *goalRepo) Delete(context.Context, *pb.PrimaryKey) error
