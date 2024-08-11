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

func (a *accountRepo) Create(ctx context.Context, request *pb.CreateAccount) (string, error) {
	var (
		err   error
		model models.CreateAccount
		res   *mongo.InsertOneResult
	)
	model = models.CreateAccount{
		UserID:   request.GetUserId(),
		Name:     request.GetName(),
		Type:     request.GetType(),
		Balance:  request.GetBalance(),
		Currency: request.GetCurrency(),
	}

	res, err = a.db.Collection("accounts").InsertOne(ctx, model)
	if err != nil {
		a.log.Error("Error while creating account in storage layer")
		return "", err
	}

	id, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		a.log.Error("Error while getting id of created account in storage layer")
		return "", fmt.Errorf("error while getting id of created account")
	}

	return id.Hex(), nil
}


func (a *accountRepo) GetById(ctx context.Context, request *pb.PrimaryKey) (*pb.Account, error) {
	var (
		err    error
		model  models.Account
		filter bson.M
	)

	objectID, err := primitive.ObjectIDFromHex(request.Id)
	if err != nil {
		a.log.Error("Invalid ID format")
		return nil, err
	}

	filter = bson.M{"_id": objectID}
	err = a.db.Collection("accounts").FindOne(ctx, filter).Decode(&model)
	if err != nil {
		a.log.Error("Error while getting account by id in storage layer")
		return nil, err
	}

	return &pb.Account{
		Id:        model.ID, 
		UserId:    model.UserID,
		Name:      model.Name,
		Type:      model.Type,
		Balance:   model.Balance,
		Currency:  model.Currency,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}, nil
}


func (a *accountRepo) GetAll(ctx context.Context, request *pb.AccountFilter) (*pb.Accounts, error) {
	var (
		err    error
		res    = &pb.Accounts{}
		model  = models.Account{}
		filter = bson.M{}
		cursor *mongo.Cursor
	)

	if request.UserId != "" {
		filter["user_id"] = request.UserId
	}
	if request.Name != "" {
		filter["name"] = request.Name
	}
	if request.Type != "" {
		filter["type"] = request.Type
	}
	if request.BalanceFrom != 0 && request.BalanceTo != 0 {
		filter["balance"] = bson.M{"$gte": request.BalanceFrom}
		filter["balance"] = bson.M{"$lte": request.BalanceTo}
	}
	if request.Currency != "" {
		filter["currency"] = request.Currency
	}

	cursor, err = a.db.Collection("accounts").Find(ctx, filter)
	if err != nil {
		a.log.Error("Error while getting all accounts in storage layer")
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		err = cursor.Decode(&model)
		if err != nil {
			a.log.Error("Error while getting all accounts in storage layer")
			return nil, err
		}
		res.Accounts = append(res.Accounts, &pb.Account{
			Id:        model.ID,
			UserId:    model.UserID,
			Name:      model.Name,
			Type:      model.Type,
			Balance:   model.Balance,
			Currency:  model.Currency,
			CreatedAt: model.CreatedAt,
			UpdatedAt: model.UpdatedAt,
		})
	}
	if err = cursor.Err(); err != nil {
		a.log.Error("Error while getting all accounts in storage layer")
		return nil, err
	}

	return res, nil
}

func (a *accountRepo) Update(ctx context.Context, request *pb.Account) error {
	var (
		err    error
		filter = bson.M{}
		update = bson.M{}
	)
	update["$set"] = bson.M{
		"user_id":    request.UserId,
		"name":       request.Name,
		"type":       request.Type,
		"balance":    request.Balance,
		"currency":   request.Currency,
		"updated_at": time.Now().Format(time.RFC3339),
	}

	objectID, err := primitive.ObjectIDFromHex(request.Id)
	if err != nil {
		a.log.Error("Invalid ID format")
		return err
	}

	filter["_id"] = objectID

	_, err = a.db.Collection("accounts").UpdateOne(ctx, filter, update)
	if err != nil {
			a.log.Error("Error while updating account in storage layer")
			return err
	}

	return nil
}

func (a *accountRepo) Delete(ctx context.Context, request *pb.PrimaryKey) error {
	var (
		err    error
		filter = bson.M{}
	)

	objectID, err := primitive.ObjectIDFromHex(request.Id)
	if err != nil {
		a.log.Error("Invalid ID format")
		return err
	}

	filter["_id"] = objectID

	_, err = a.db.Collection("accounts").DeleteOne(ctx, filter)
	if err != nil {
			a.log.Error("Error while deleting account in storage layer")
			return err
	}

	return nil
}
