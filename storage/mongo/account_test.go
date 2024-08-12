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

var (
	client     *mongo.Client
	db         *mongo.Database
	repo       *accountRepo
	testLogger logger.ILogger
)

func setup() {
	cfg  := configs.Load()
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s", cfg.MongoDBHost, cfg.MongoDBPort))
	var err error
	client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		panic(err)
	}
	db = client.Database("test_db")
	testLogger = logger.NewLogger("", "debug", "app.log")
	repo = NewAccountRepo(db, testLogger)
}

func close() {
	client.Disconnect(context.Background())
}

func TestCreateAccount(t *testing.T) {
	setup()
	defer close()

	reqCreate := &pb.CreateAccount{
		UserId:   "user1",
		Name:     "Test Account",
		Type:     "Checking",
		Balance:  100.0,
		Currency: "USD",
	}
	id, err := repo.Create(context.Background(), reqCreate)
	if err != nil {
		t.Fatal(err)
	}

	err = repo.Delete(context.Background(), &pb.PrimaryKey{Id: id})
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetById(t *testing.T) {
	setup()
	defer close()

	reqCreate := &pb.CreateAccount{
		UserId:   "user1",
		Name:     "Test Account",
		Type:     "Checking",
		Balance:  100.0,
		Currency: "USD",
	}
	id, err := repo.Create(context.Background(), reqCreate)
	if err != nil {
		t.Fatal(err)
	}

	reqGetById := &pb.PrimaryKey{Id: id}
	account, err := repo.GetById(context.Background(), reqGetById)
	if err != nil {
		t.Fatal(err)
	}
	if account.Name != reqCreate.Name {
		t.Errorf("Expected account name %s, got %s", reqCreate.Name, account.Name)
	}

	err = repo.Delete(context.Background(), reqGetById)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUpdate(t *testing.T) {
	setup()
	defer close()

	reqCreate := &pb.CreateAccount{
		UserId:   "user1",
		Name:     "Test Account",
		Type:     "Checking",
		Balance:  100.0,
		Currency: "USD",
	}
	id, err := repo.Create(context.Background(), reqCreate)
	if err != nil {
		t.Fatal(err)
	}

	reqGetById := &pb.PrimaryKey{Id: id}
	account, err := repo.GetById(context.Background(), reqGetById)
	if err != nil {
		t.Fatal(err)
	}

	account.Name = "Updated Account"
	err = repo.Update(context.Background(), account)
	if err != nil {
		t.Fatal(err)
	}

	updatedAccount, err := repo.GetById(context.Background(), reqGetById)
	if err != nil {
		t.Fatal(err)
	}
	if updatedAccount.Name != "Updated Account" {
		t.Errorf("Expected updated account name %s, got %s", "Updated Account", updatedAccount.Name)
	}

	err = repo.Delete(context.Background(), reqGetById)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDelete(t *testing.T) {
	setup()
	defer close()

	reqCreate := &pb.CreateAccount{
		UserId:   "user1",
		Name:     "Test Account",
		Type:     "Checking",
		Balance:  100.0,
		Currency: "USD",
	}
	id, err := repo.Create(context.Background(), reqCreate)
	if err != nil {
		t.Fatal(err)
	}

	reqGetById := &pb.PrimaryKey{Id: id}
	err = repo.Delete(context.Background(), reqGetById)
	if err != nil {
		t.Fatal(err)
	}

	deletedAccount, _ := repo.GetById(context.Background(), reqGetById)
	if deletedAccount != nil {
		t.Fatal("Error while deleting didnt deleted")
	}
}
