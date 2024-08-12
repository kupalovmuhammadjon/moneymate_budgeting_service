package mongo

import (
	"budgeting_service/models"
	"budgeting_service/pkg/logger"
	"context"
	"fmt"
	"time"

	pb "budgeting_service/genproto/budgeting_service"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (b *budgetRepo) Create(ctx context.Context, request *pb.CreateBudget) (*pb.Budget, error) {
	model := models.CreateBudget{
		UserID:     request.GetUserId(),
		CategoryID: request.GetCategoryId(),
		Amount:     request.GetAmount(),
		Period:     request.GetPeriod(),
		StartDate:  request.GetStartDate(),
		EndDate:    request.GetEndDate(),
	}

	res, err := b.db.Collection("budgets").InsertOne(ctx, model)
	if err != nil {
		b.log.Error("Error while creating budget in storage layer")
		return nil, err
	}

	id, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		b.log.Error("Error while getting id of created budget in storage layer")
		return nil, fmt.Errorf("error while getting id of created budget")
	}

	return &pb.Budget{
		Id:         id.Hex(),
		UserId:     model.UserID,
		CategoryId: model.CategoryID,
		Amount:     model.Amount,
		Period:     model.Period,
		StartDate:  model.StartDate,
		EndDate:    model.EndDate,
		CreatedAt:  time.Now().Format(time.RFC3339),
		UpdatedAt:  time.Now().Format(time.RFC3339),
	}, nil
}

func (b *budgetRepo) GetById(ctx context.Context, request *pb.PrimaryKey) (*pb.Budget, error) {
	objectID, err := primitive.ObjectIDFromHex(request.GetId())
	if err != nil {
		b.log.Error("Invalid ID format")
		return nil, err
	}

	filter := bson.M{"_id": objectID}
	var model models.Budget
	err = b.db.Collection("budgets").FindOne(ctx, filter).Decode(&model)
	if err != nil {
		b.log.Error("Error while getting budget by id in storage layer")
		return nil, err
	}

	return &pb.Budget{
		Id:         model.ID,
		UserId:     model.UserID,
		CategoryId: model.CategoryID,
		Amount:     model.Amount,
		Period:     model.Period,
		StartDate:  model.StartDate,
		EndDate:    model.EndDate,
		CreatedAt:  model.CreatedAt,
		UpdatedAt:  model.UpdatedAt,
	}, nil
}

func (b *budgetRepo) GetAll(ctx context.Context, request *pb.BudgetFilter) (*pb.Budgets, error) {
	filter := bson.M{}
	if request.UserId != "" {
		filter["user_id"] = request.UserId
	}
	if request.CategoryId != "" {
		filter["category_id"] = request.CategoryId
	}
	if request.Amount != 0 {
		filter["amount"] = request.Amount
	}
	if request.Period != "" {
		filter["period"] = request.Period
	}
	if request.StartDate != "" && request.EndDate != "" {
		filter["$and"] = bson.A{
			bson.M{"start_date": bson.M{"$gte": request.StartDate}},
			bson.M{"end_date": bson.M{"$lte": request.EndDate}},
		}
	}
	
	offset := int64((request.Page - 1) * 10)
	limit := int64(request.Limit)
	options := &options.FindOptions{
		Skip:  &offset,
		Limit: &limit,
	}

	cursor, err := b.db.Collection("budgets").Find(ctx, filter, options)
	if err != nil {
		b.log.Error("Error while getting all budgets in storage layer", logger.Error(err))
		return nil, err
	}
	defer cursor.Close(ctx)

	var budgets []*pb.Budget
	for cursor.Next(ctx) {
		var model models.Budget
		err = cursor.Decode(&model)
		if err != nil {
			b.log.Error("Error while decoding budget in storage layer", logger.Error(err))
			return nil, err
		}

		budgets = append(budgets, &pb.Budget{
			Id:         model.ID,
			UserId:     model.UserID,
			CategoryId: model.CategoryID,
			Amount:     model.Amount,
			Period:     model.Period,
			StartDate:  model.StartDate,
			EndDate:    model.EndDate,
			CreatedAt:  model.CreatedAt,
			UpdatedAt:  model.UpdatedAt,
		})
	}
	if err = cursor.Err(); err != nil {
		b.log.Error("Error while iterating over cursor in storage layer", logger.Error(err))
		return nil, err
	}

	return &pb.Budgets{Budgets: budgets}, nil
}

func (b *budgetRepo) Update(ctx context.Context, request *pb.Budget) (*pb.Budget, error) {
	objectID, err := primitive.ObjectIDFromHex(request.GetId())
	if err != nil {
		b.log.Error("Invalid ID format")
		return nil, err
	}

	update := bson.M{
		"$set": bson.M{
			"category_id": request.GetCategoryId(),
			"amount":      request.GetAmount(),
			"period":      request.GetPeriod(),
			"start_date":  request.GetStartDate(),
			"end_date":    request.GetEndDate(),
			"updated_at":  time.Now().Format(time.RFC3339),
		},
	}

	filter := bson.M{"_id": objectID}
	_, err = b.db.Collection("budgets").UpdateOne(ctx, filter, update)
	if err != nil {
		b.log.Error("Error while updating budget in storage layer")
		return nil, err
	}

	return &pb.Budget{
		Id:         request.GetId(),
		UserId:     request.GetUserId(),
		CategoryId: request.GetCategoryId(),
		Amount:     request.GetAmount(),
		Period:     request.GetPeriod(),
		StartDate:  request.GetStartDate(),
		EndDate:    request.GetEndDate(),
		CreatedAt:  request.GetCreatedAt(),
		UpdatedAt:  time.Now().Format(time.RFC3339),
	}, nil
}

func (b *budgetRepo) Delete(ctx context.Context, request *pb.PrimaryKey) error {
	objectID, err := primitive.ObjectIDFromHex(request.GetId())
	if err != nil {
		b.log.Error("Invalid ID format")
		return err
	}

	filter := bson.M{"_id": objectID}
	_, err = b.db.Collection("budgets").DeleteOne(ctx, filter)
	if err != nil {
		b.log.Error("Error while deleting budget in storage layer")
		return err
	}

	return nil
}
