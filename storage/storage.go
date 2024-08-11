package storage

import (
	"budgeting_service/configs"
	pb "budgeting_service/genproto/budgeting_service"
	"budgeting_service/pkg/logger"
	repo "budgeting_service/storage/mongo"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type IStorage interface {
	Close()
	Accounts() IAccountRepo
	Budgets() IBudgetRepo
	Categories() ICategoryRepo
	Goals() IGoalRepo
	Transactions() ITransactionRepo
}

type IAccountRepo interface {
	Create(context.Context, *pb.CreateAccount) (string, error)
	GetById(context.Context, *pb.PrimaryKey) (*pb.Account, error)
	GetAll(context.Context, *pb.AccountFilter) (*pb.Accounts, error)
	Update(context.Context, *pb.Account) error
	Delete(context.Context, *pb.PrimaryKey) error
}

type IBudgetRepo interface {
	Create(context.Context, *pb.CreateBudget) (*pb.Budget, error)
	GetById(context.Context, *pb.PrimaryKey) (*pb.Budget, error)
	GetAll(context.Context, *pb.BudgetFilter) (*pb.Budgets, error)
	Update(context.Context, *pb.Budget) (*pb.Budget, error)
	Delete(context.Context, *pb.PrimaryKey) error
}

type ICategoryRepo interface {
	Create(context.Context, *pb.CreateCategory) (*pb.Category, error)
	GetById(context.Context, *pb.PrimaryKey) (*pb.Category, error)
	GetAll(context.Context, *pb.CategoryFilter) (*pb.Categories, error)
	Update(context.Context, *pb.Category) (*pb.Category, error)
	Delete(context.Context, *pb.PrimaryKey) error
}

type IGoalRepo interface {
	Create(context.Context, *pb.CreateGoal) (*pb.Goal, error)
	GetById(context.Context, *pb.PrimaryKey) (*pb.Goal, error)
	GetAll(context.Context, *pb.GoalFilter) (*pb.Goals, error)
	Update(context.Context, *pb.Goal) (*pb.Goal, error)
	Delete(context.Context, *pb.PrimaryKey) error
}

type ITransactionRepo interface {
	Create(context.Context, *pb.CreateTransaction) (*pb.Transaction, error)
	GetById(context.Context, *pb.PrimaryKey) (*pb.Transaction, error)
	GetAll(context.Context, *pb.TransactionFilter) (*pb.Transactions, error)
	Update(context.Context, *pb.Transaction) (*pb.Transaction, error)
	Delete(context.Context, *pb.PrimaryKey) error
	GenerateSpendingReport(context.Context, *pb.PrimaryKey) (*pb.Spendings, error)
	GenerateIncomeReport(context.Context, *pb.PrimaryKey) (*pb.Incomes, error)
	GenerateBudgetPerformanceReport(context.Context, *pb.PrimaryKey) (*pb.BugetPerformance, error)
	GenerateGoalProgressReport(context.Context, *pb.PrimaryKey) (*pb.GoalProgress, error)
}

type storage struct {
	db  *mongo.Database
	log logger.ILogger
}

func NewIStorage(ctx context.Context, cfg *configs.Config, log logger.ILogger) (IStorage, error) {
	db, err := repo.ConnectMongoDB(ctx, cfg)
	if err != nil {
		return nil, err
	}

	return &storage{
		db:  db,
		log: log,
	}, nil
}

func (s *storage) Close() {
	s.db.Client().Disconnect(context.Background())
}

func (s *storage) Accounts() IAccountRepo {
	return repo.NewAccountRepo(s.db, s.log)
}

func (s *storage) Budgets() IBudgetRepo {
	return repo.NewBudgetRepo(s.db, s.log)
}

func (s *storage) Categories() ICategoryRepo {
	return repo.NewCategoryRepo(s.db, s.log)
}

func (s *storage) Goals() IGoalRepo {
	return repo.NewGoalRepo(s.db, s.log)
}

func (s *storage) Transactions() ITransactionRepo {
	return repo.NewTransactionRepo(s.db, s.log)
}
