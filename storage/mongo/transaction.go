package mongo

import (
	"context"
	"fmt"
	"time"

	pb "budgeting_service/genproto/budgeting_service"
	"budgeting_service/models"
	"budgeting_service/pkg/logger"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (t *transactionRepo) Create(ctx context.Context, req *pb.CreateTransaction) (*pb.Transaction, error) {
	model := models.CreateTransaction{
		UserID:      req.GetUserId(),
		AccountID:   req.GetAccountId(),
		CategoryID:  req.GetCategoryId(),
		Amount:      req.GetAmount(),
		Type:        req.GetType(),
		Description: req.GetDescription(),
		Date:        req.GetDate(),
	}

	res, err := t.db.Collection("transactions").InsertOne(ctx, model)
	if err != nil {
		t.log.Error("Error while creating transaction in storage layer", logger.Error(err))
		return nil, err
	}

	id, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		t.log.Error("Error while getting ID of created transaction in storage layer")
		return nil, fmt.Errorf("error while getting ID of created transaction")
	}

	return &pb.Transaction{
		Id:          id.Hex(),
		UserId:      model.UserID,
		AccountId:   model.AccountID,
		CategoryId:  model.CategoryID,
		Amount:      model.Amount,
		Type:        model.Type,
		Description: model.Description,
		Date:        model.Date,
		CreatedAt:   time.Now().Format(time.RFC3339),
		UpdatedAt:   time.Now().Format(time.RFC3339),
	}, nil
}

func (t *transactionRepo) GetById(ctx context.Context, req *pb.PrimaryKey) (*pb.Transaction, error) {
	var transaction models.Transaction
	objectID, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		t.log.Error("Invalid ID format")
		return nil, err
	}

	filter := bson.M{"_id": objectID}

	err = t.db.Collection("transactions").FindOne(ctx, filter).Decode(&transaction)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		t.log.Error("Error while getting transaction by ID", logger.Error(err))
		return nil, err
	}

	return &pb.Transaction{
		Id:          transaction.ID,
		UserId:      transaction.UserID,
		AccountId:   transaction.AccountID,
		CategoryId:  transaction.CategoryID,
		Amount:      transaction.Amount,
		Type:        transaction.Type,
		Description: transaction.Description,
		Date:        transaction.Date,
		CreatedAt:   transaction.CreatedAt,
		UpdatedAt:   transaction.UpdatedAt,
	}, nil
}

func (t *transactionRepo) GetAll(ctx context.Context, req *pb.TransactionFilter) (*pb.Transactions, error) {
	filter := bson.M{}
	if req.UserId != "" {
		filter["user_id"] = req.UserId
	}
	if req.AccountId != "" {
		filter["account_id"] = req.AccountId
	}
	if req.CategoryId != "" {
		filter["category_id"] = req.CategoryId
	}
	if req.Amount != 0 {
		filter["amount"] = req.Amount
	}
	if req.Type != "" {
		filter["type"] = req.Type
	}
	if req.Date != "" {
		filter["date"] = req.Date
	}

	offset := int64((req.Page - 1) * 10)
	limit := int64(req.Limit)
	options := &options.FindOptions{
		Skip:  &offset,
		Limit: &limit,
	}

	cursor, err := t.db.Collection("transactions").Find(ctx, filter, options)
	if err != nil {
		t.log.Error("Error while getting all transactions", logger.Error(err))
		return nil, err
	}
	defer cursor.Close(ctx)

	var transactions []*pb.Transaction
	for cursor.Next(ctx) {
		var model models.Transaction
		err = cursor.Decode(&model)
		if err != nil {
			t.log.Error("Error while decoding transaction", logger.Error(err))
			return nil, err
		}

		transactions = append(transactions, &pb.Transaction{
			Id:          model.ID,
			UserId:      model.UserID,
			AccountId:   model.AccountID,
			CategoryId:  model.CategoryID,
			Amount:      model.Amount,
			Type:        model.Type,
			Description: model.Description,
			Date:        model.Date,
			CreatedAt:   model.CreatedAt,
			UpdatedAt:   model.UpdatedAt,
		})
	}
	if err = cursor.Err(); err != nil {
		t.log.Error("Error while iterating over cursor", logger.Error(err))
		return nil, err
	}

	return &pb.Transactions{Transactions: transactions}, nil
}

func (t *transactionRepo) Update(ctx context.Context, req *pb.Transaction) (*pb.Transaction, error) {
	objectID, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		t.log.Error("Invalid ID format")
		return nil, err
	}

	filter := bson.M{"_id": objectID}

	update := bson.M{
		"$set": bson.M{
			"user_id":     req.UserId,
			"account_id":  req.AccountId,
			"category_id": req.CategoryId,
			"amount":      req.Amount,
			"type":        req.Type,
			"description": req.Description,
			"date":        req.Date,
			"updated_at":  time.Now().Format(time.RFC3339),
		},
	}

	_, err = t.db.Collection("transactions").UpdateOne(ctx, filter, update)
	if err != nil {
		t.log.Error("Error while updating transaction", logger.Error(err))
		return nil, err
	}

	return &pb.Transaction{
		Id:          req.Id,
		UserId:      req.UserId,
		AccountId:   req.AccountId,
		CategoryId:  req.CategoryId,
		Amount:      req.Amount,
		Type:        req.Type,
		Description: req.Description,
		Date:        req.Date,
		CreatedAt:   req.CreatedAt,
		UpdatedAt:   time.Now().Format(time.RFC3339),
	}, nil
}

func (t *transactionRepo) Delete(ctx context.Context, req *pb.PrimaryKey) error {
	objectID, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		t.log.Error("Invalid ID format")
		return err
	}

	filter := bson.M{"_id": objectID}

	_, err = t.db.Collection("transactions").DeleteOne(ctx, filter)
	if err != nil {
		t.log.Error("Error while deleting transaction", logger.Error(err))
		return err
	}

	return nil
}

func (t *transactionRepo) GenerateSpendingReport(context.Context, *pb.PrimaryKey) (*pb.Spendings, error) {

	return nil, nil
}
func (t *transactionRepo) GenerateIncomeReport(context.Context, *pb.PrimaryKey) (*pb.Incomes, error) {

	return nil, nil
}
func (t *transactionRepo) GenerateBudgetPerformanceReport(context.Context, *pb.PrimaryKey) (*pb.BugetPerformance, error) {

	return nil, nil
}
func (t *transactionRepo) GenerateGoalProgressReport(context.Context, *pb.PrimaryKey) (*pb.GoalProgress, error) {

	return nil, nil
}
