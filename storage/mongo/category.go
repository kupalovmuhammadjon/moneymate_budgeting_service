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

type categoryRepo struct {
	db  *mongo.Database
	log logger.ILogger
}

func NewCategoryRepo(db *mongo.Database, log logger.ILogger) *categoryRepo {
	return &categoryRepo{
		db:  db,
		log: log,
	}
}

func (c *categoryRepo) Create(ctx context.Context, request *pb.CreateCategory) (*pb.Category, error) {
	model := models.CreateCategory{
		UserID: request.GetUserId(),
		Name:   request.GetName(),
		Type:   request.GetType(),
	}

	res, err := c.db.Collection("categories").InsertOne(ctx, model)
	if err != nil {
		c.log.Error("Error while creating category in storage layer", logger.Error(err))
		return nil, err
	}

	id, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		c.log.Error("Error while getting ID of created category in storage layer")
		return nil, fmt.Errorf("error while getting ID of created category")
	}

	return &pb.Category{
		Id:        id.Hex(),
		UserId:    model.UserID,
		Name:      model.Name,
		Type:      model.Type,
		CreatedAt: time.Now().Format(time.RFC3339),
		UpdatedAt: time.Now().Format(time.RFC3339),
	}, nil
}

func (c *categoryRepo) GetById(ctx context.Context, req *pb.PrimaryKey) (*pb.Category, error) {
	var category models.Category
	objectID, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		c.log.Error("Invalid ID format")
		return nil, err
	}

	filter := bson.M{"_id": objectID}

	err = c.db.Collection("categories").FindOne(ctx, filter).Decode(&category)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		c.log.Error("Error while getting category by ID", logger.Error(err))
		return nil, err
	}

	return &pb.Category{
		Id:        category.ID,
		UserId:    category.UserID,
		Name:      category.Name,
		Type:      category.Type,
		CreatedAt: category.CreatedAt,
		UpdatedAt: category.UpdatedAt,
	}, nil
}

func (c *categoryRepo) GetAll(ctx context.Context, req *pb.CategoryFilter) (*pb.Categories, error) {
	filter := bson.M{}
	if req.UserId != "" {
		filter["user_id"] = req.UserId
	}
	if req.Name != "" {
		filter["name"] = req.Name
	}
	if req.Type != "" {
		filter["type"] = req.Type
	}

	offset := int64((req.Page - 1) * 10)
	limit := int64(req.Limit)
	options := &options.FindOptions{
		Skip:  &offset,
		Limit: &limit,
	}

	cursor, err := c.db.Collection("categories").Find(ctx, filter, options)
	if err != nil {
		c.log.Error("Error while getting all categories", logger.Error(err))
		return nil, err
	}
	defer cursor.Close(ctx)

	var categories []*pb.Category
	for cursor.Next(ctx) {
		var model models.Category
		err = cursor.Decode(&model)
		if err != nil {
			c.log.Error("Error while decoding category", logger.Error(err))
			return nil, err
		}

		categories = append(categories, &pb.Category{
			Id:        model.ID,
			UserId:    model.UserID,
			Name:      model.Name,
			Type:      model.Type,
			CreatedAt: model.CreatedAt,
			UpdatedAt: model.UpdatedAt,
		})
	}
	if err = cursor.Err(); err != nil {
		c.log.Error("Error while iterating over cursor", logger.Error(err))
		return nil, err
	}

	return &pb.Categories{Categories: categories}, nil
}

func (c *categoryRepo) Update(ctx context.Context, req *pb.Category) (*pb.Category, error) {
	objectID, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		c.log.Error("Invalid ID format")
		return nil, err
	}

	filter := bson.M{"_id": objectID}

	update := bson.M{
		"$set": bson.M{
			"user_id":    req.UserId,
			"name":       req.Name,
			"type":       req.Type,
			"updated_at": time.Now().Format(time.RFC3339),
		},
	}

	_, err = c.db.Collection("categories").UpdateOne(ctx, filter, update)
	if err != nil {
		c.log.Error("Error while updating category", logger.Error(err))
		return nil, err
	}

	return &pb.Category{
		Id:        req.Id,
		UserId:    req.UserId,
		Name:      req.Name,
		Type:      req.Type,
		CreatedAt: req.CreatedAt,
		UpdatedAt: time.Now().Format(time.RFC3339),
	}, nil
}

func (c *categoryRepo) Delete(ctx context.Context, req *pb.PrimaryKey) error {
	objectID, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		c.log.Error("Invalid ID format")
		return err
	}

	filter := bson.M{"_id": objectID}

	_, err = c.db.Collection("categories").DeleteOne(ctx, filter)
	if err != nil {
		c.log.Error("Error while deleting category", logger.Error(err))
		return err
	}

	return nil
}
