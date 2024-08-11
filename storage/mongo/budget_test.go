package mongo

import (
	"context"
	"fmt"
	"testing"

	"budgeting_service/configs"
	pb "budgeting_service/genproto/budgeting_service"
	"budgeting_service/pkg/logger"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


func setUpBudget() *budgetRepo {
	cfg := configs.Load()
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s", cfg.MongoDBHost, cfg.MongoDBPort))
	var err error
	client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		panic(err)
	}
	db = client.Database("test_db")
	testLogger = logger.NewLogger("", "debug", "app.log")
	return NewBudgetRepo(db, testLogger)
}

func tearDownBudget(t *testing.T) {
	if err := client.Disconnect(context.Background()); err != nil {
		t.Errorf("Failed to disconnect from MongoDB: %v", err)
	}
}

func TestCreateBudget(t *testing.T) {
	budgetRepo := setUpBudget()
	defer tearDownBudget(t)

	reqCreate := &pb.CreateBudget{
		UserId:     "user1",
		CategoryId: "category1",
		Amount:     100.0,
		Period:     "Monthly",
		StartDate:  "2024-01-01",
		EndDate:    "2024-12-31",
	}
	budget, err := budgetRepo.Create(context.Background(), reqCreate)
	if err != nil {
		t.Fatal(err)
	}

	err = budgetRepo.Delete(context.Background(), &pb.PrimaryKey{Id: budget.Id})
	if err != nil {
		t.Fatal(err)
	}
}

func TestBudgetGetById(t *testing.T) {
	budgetRepo := setUpBudget()
	defer tearDownBudget(t)

	reqCreate := &pb.CreateBudget{
		UserId:     "user1",
		CategoryId: "category1",
		Amount:     100.0,
		Period:     "Monthly",
		StartDate:  "2024-01-01",
		EndDate:    "2024-12-31",
	}
	budget, err := budgetRepo.Create(context.Background(), reqCreate)
	if err != nil {
		t.Fatal(err)
	}

	reqGetById := &pb.PrimaryKey{Id: budget.Id}
	gotBudget, err := budgetRepo.GetById(context.Background(), reqGetById)
	if err != nil {
		t.Fatal(err)
	}
	if gotBudget.Amount != reqCreate.Amount {
		t.Errorf("Expected budget amount %f, got %f", reqCreate.Amount, gotBudget.Amount)
	}

	err = budgetRepo.Delete(context.Background(), reqGetById)
	if err != nil {
		t.Fatal(err)
	}
}

func TestBudgetUpdate(t *testing.T) {
	budgetRepo := setUpBudget()
	defer tearDownBudget(t)

	reqCreate := &pb.CreateBudget{
		UserId:     "user1",
		CategoryId: "category1",
		Amount:     100.0,
		Period:     "Monthly",
		StartDate:  "2024-01-01",
		EndDate:    "2024-12-31",
	}
	budget, err := budgetRepo.Create(context.Background(), reqCreate)
	if err != nil {
		t.Fatal(err)
	}

	reqGetById := &pb.PrimaryKey{Id: budget.Id}
	budgetToUpdate, err := budgetRepo.GetById(context.Background(), reqGetById)
	if err != nil {
		t.Fatal(err)
	}

	budgetToUpdate.Amount = 150.0
	updatedBudget, err := budgetRepo.Update(context.Background(), budgetToUpdate)
	if err != nil {
		t.Fatal(err)
	}

	if updatedBudget.Amount != 150.0 {
		t.Errorf("Expected updated budget amount %f, got %f", 150.0, updatedBudget.Amount)
	}

	err = budgetRepo.Delete(context.Background(), reqGetById)
	if err != nil {
		t.Fatal(err)
	}
}

func TestBudgetDelete(t *testing.T) {
	budgetRepo := setUpBudget()
	defer tearDownBudget(t)

	reqCreate := &pb.CreateBudget{
		UserId:     "user1",
		CategoryId: "category1",
		Amount:     100.0,
		Period:     "Monthly",
		StartDate:  "2024-01-01",
		EndDate:    "2024-12-31",
	}
	budget, err := budgetRepo.Create(context.Background(), reqCreate)
	if err != nil {
		t.Fatal(err)
	}

	reqGetById := &pb.PrimaryKey{Id: budget.Id}
	err = budgetRepo.Delete(context.Background(), reqGetById)
	if err != nil {
		t.Fatal(err)
	}

	deletedBudget, _ := budgetRepo.GetById(context.Background(), reqGetById)
	if deletedBudget != nil {
		t.Fatal("Error: Budget not deleted")
	}
}
