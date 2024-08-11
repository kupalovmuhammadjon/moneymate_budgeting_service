package mongo

import (
	"budgeting_service/pkg/logger"
	"context"

	pb "budgeting_service/genproto/budgeting_service"

	"go.mongodb.org/mongo-driver/mongo"
)

type transactionRepo struct {
	db  *mongo.Database
	log logger.ILogger
}

func NewTransactionRepo(db *mongo.Database, log logger.ILogger) *transactionRepo {
	return &transactionRepo{
		db:  db,
		log: log,
	}
}

func (t *transactionRepo) Create(context.Context, *pb.CreateTransaction) (*pb.Transaction, error){

	return nil, nil
}
func (t *transactionRepo) GetById(context.Context, *pb.PrimaryKey) (*pb.Transaction, error){

	return nil, nil
}
func (t *transactionRepo) GetAll(context.Context, *pb.TransactionFilter) (*pb.Transactions, error){

	return nil, nil
}
func (t *transactionRepo) Update(context.Context, *pb.Transaction) (*pb.Transaction, error){

	return nil, nil
}
func (t *transactionRepo) Delete(context.Context, *pb.PrimaryKey) error{

	return nil
}
func (t *transactionRepo) GenerateSpendingReport(context.Context, *pb.PrimaryKey) (*pb.Spendings, error){

	return nil, nil
}
func (t *transactionRepo) GenerateIncomeReport(context.Context, *pb.PrimaryKey) (*pb.Incomes, error){

	return nil, nil
}
func (t *transactionRepo) GenerateBudgetPerformanceReport(context.Context, *pb.PrimaryKey) (*pb.BugetPerformance, error){

	return nil, nil
}
func (t *transactionRepo) GenerateGoalProgressReport(context.Context, *pb.PrimaryKey) (*pb.GoalProgress, error){

	return nil, nil
}