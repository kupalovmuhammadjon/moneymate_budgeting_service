package mongo

import (
	"budgeting_service/pkg/logger"
	"context"

	pb "budgeting_service/genproto/budgeting_service"

	"go.mongodb.org/mongo-driver/mongo"
)

type accountRepo struct {
	db  *mongo.Database
	log logger.ILogger
}

func NewAccountRepo(db *mongo.Database, log logger.ILogger) *accountRepo {
	return &accountRepo{
		db:  db,
		log: log,
	}
}

func (a *accountRepo) Create(context.Context, *pb.CreateAccount) (*pb.Account, error)
func (a *accountRepo) GetById(context.Context, *pb.PrimaryKey) (*pb.Account, error)
func (a *accountRepo) GetAll(context.Context, *pb.AccountFilter) (*pb.Accounts, error)
func (a *accountRepo) Update(context.Context, *pb.Account) (*pb.Account, error)
func (a *accountRepo) Delete(context.Context, *pb.PrimaryKey) error
