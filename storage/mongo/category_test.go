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


func setUpCategory() *categoryRepo {
	cfg := configs.Load()
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s", cfg.MongoDBHost, cfg.MongoDBPort))
	var err error
	client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		panic(err)
	}
	db = client.Database("test_db")
	testLogger = logger.NewLogger("", "debug", "app.log")
	return NewCategoryRepo(db, testLogger)
}

func tearDownCategory(t *testing.T) {
	if err := client.Disconnect(context.Background()); err != nil {
		t.Errorf("Failed to disconnect from MongoDB: %v", err)
	}
}

func TestCreateCategory(t *testing.T) {
	categoryRepo := setUpCategory()
	defer tearDownCategory(t)

	reqCreate := &pb.CreateCategory{
		UserId: "user123",
		Name:   "Entertainment",
		Type:   "Expense",
	}
	category, err := categoryRepo.Create(context.Background(), reqCreate)
	if err != nil {
		t.Fatal(err)
	}

	if category.Id == "" {
		t.Fatal("Expected non-empty category ID")
	}

	err = categoryRepo.Delete(context.Background(), &pb.PrimaryKey{Id: category.Id})
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetByIdCategory(t *testing.T) {
	categoryRepo := setUpCategory()
	defer tearDownCategory(t)

	reqCreate := &pb.CreateCategory{
		UserId: "user123",
		Name:   "Entertainment",
		Type:   "Expense",
	}
	category, err := categoryRepo.Create(context.Background(), reqCreate)
	if err != nil {
		t.Fatal("Failed to create category:", err)
	}

	if category == nil {
		t.Fatal("Created category is nil")
	}

	reqGetById := &pb.PrimaryKey{Id: category.Id}
	gotCategory, err := categoryRepo.GetById(context.Background(), reqGetById)
	if err != nil {
		t.Fatal("Failed to get category by ID:", err)
	}

	if gotCategory == nil {
		t.Fatal("Got category is nil")
	}



	err = categoryRepo.Delete(context.Background(), reqGetById)
	if err != nil {
		t.Fatal("Failed to delete category:", err)
	}
}


func TestUpdateCategory(t *testing.T) {
	categoryRepo := setUpCategory()
	defer tearDownCategory(t)

	reqCreate := &pb.CreateCategory{
		UserId: "user123",
		Name:   "Entertainment",
		Type:   "Expense",
	}
	category, err := categoryRepo.Create(context.Background(), reqCreate)
	if err != nil {
		t.Fatal(err)
	}

	reqGetById := &pb.PrimaryKey{Id: category.Id}
	categoryToUpdate, err := categoryRepo.GetById(context.Background(), reqGetById)
	if err != nil {
		t.Fatal(err)
	}

	categoryToUpdate.Name = "Updated Category"
	updatedCategory, err := categoryRepo.Update(context.Background(), categoryToUpdate)
	if err != nil {
		t.Fatal(err)
	}

	if updatedCategory.Name != "Updated Category" {
		t.Errorf("Expected updated category name %s, got %s", "Updated Category", updatedCategory.Name)
	}

	err = categoryRepo.Delete(context.Background(), reqGetById)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDeleteCategory(t *testing.T) {
	categoryRepo := setUpCategory()
	defer tearDownCategory(t)

	reqCreate := &pb.CreateCategory{
		UserId: "user123",
		Name:   "Entertainment",
		Type:   "Expense",
	}
	category, err := categoryRepo.Create(context.Background(), reqCreate)
	if err != nil {
		t.Fatal(err)
	}

	reqGetById := &pb.PrimaryKey{Id: category.Id}
	err = categoryRepo.Delete(context.Background(), reqGetById)
	if err != nil {
		t.Fatal(err)
	}

	deletedCategory, err := categoryRepo.GetById(context.Background(), reqGetById)
	if err != nil && err != mongo.ErrNoDocuments {
		t.Fatal("Expected error for deleted category")
	}
	if deletedCategory != nil {
		t.Fatal("Error: Category not deleted")
	}
}
