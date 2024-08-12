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

	type goalRepo struct {
		db  *mongo.Database
		log logger.ILogger
	}

	func NewGoalRepo(db *mongo.Database, log logger.ILogger) *goalRepo {
		return &goalRepo{
			db:  db,
			log: log,
		}
	}

	func (g *goalRepo) Create(ctx context.Context, req *pb.CreateGoal) (*pb.Goal, error) {
		model := models.CreateGoal{
			UserID:        req.GetUserId(),
			Name:          req.GetName(),
			TargetAmount:  req.GetTargetAmount(),
			CurrentAmount: req.GetCurrentAmount(),
			Deadline:      req.GetDeadline(),
			Status:        req.GetStatus(),
		}

		res, err := g.db.Collection("goals").InsertOne(ctx, model)
		if err != nil {
			g.log.Error("Error while creating goal in storage layer", logger.Error(err))
			return nil, err
		}

		id, ok := res.InsertedID.(primitive.ObjectID)
		if !ok {
			g.log.Error("Error while getting ID of created goal in storage layer")
			return nil, fmt.Errorf("error while getting ID of created goal")
		}

		return &pb.Goal{
			Id:            id.Hex(),
			UserId:        model.UserID,
			Name:          model.Name,
			TargetAmount:  model.TargetAmount,
			CurrentAmount: model.CurrentAmount,
			Deadline:      model.Deadline,
			Status:        model.Status,
			CreatedAt:     time.Now().Format(time.RFC3339),
			UpdatedAt:     time.Now().Format(time.RFC3339),
		}, nil
	}

	func (g *goalRepo) GetById(ctx context.Context, req *pb.PrimaryKey) (*pb.Goal, error) {
		var goal models.Goal
		objectID, err := primitive.ObjectIDFromHex(req.GetId())
		if err != nil {
			g.log.Error("Invalid ID format")
			return nil, err
		}

		filter := bson.M{"_id": objectID}

		err = g.db.Collection("goals").FindOne(ctx, filter).Decode(&goal)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				return nil, nil
			}
			g.log.Error("Error while getting goal by ID", logger.Error(err))
			return nil, err
		}

		return &pb.Goal{
			Id:            goal.ID,
			UserId:        goal.UserID,
			Name:          goal.Name,
			TargetAmount:  goal.TargetAmount,
			CurrentAmount: goal.CurrentAmount,
			Deadline:      goal.Deadline,
			Status:        goal.Status,
			CreatedAt:     goal.CreatedAt,
			UpdatedAt:     goal.UpdatedAt,
		}, nil
	}

	func (g *goalRepo) GetAll(ctx context.Context, req *pb.GoalFilter) (*pb.Goals, error) {
		filter := bson.M{}
		if req.UserId != "" {
			filter["user_id"] = req.UserId
		}
		if req.Name != "" {
			filter["name"] = req.Name
		}
		if req.TargetAmount != 0 {
			filter["target_amount"] = req.TargetAmount
		}
		if req.CurrentAmount != 0 {
			filter["current_amount"] = req.CurrentAmount
		}
		if req.Deadline != "" {
			filter["deadline"] = req.Deadline
		}
		if req.Status != "" {
			filter["status"] = req.Status
		}

		offset := int64((req.Page - 1) * 10)
		limit := int64(req.Limit)
		options := &options.FindOptions{
			Skip:  &offset,
			Limit: &limit,
		}

		cursor, err := g.db.Collection("goals").Find(ctx, filter, options)
		if err != nil {
			g.log.Error("Error while getting all goals", logger.Error(err))
			return nil, err
		}
		defer cursor.Close(ctx)

		var goals []*pb.Goal
		for cursor.Next(ctx) {
			var model models.Goal
			err = cursor.Decode(&model)
			if err != nil {
				g.log.Error("Error while decoding goal", logger.Error(err))
				return nil, err
			}

			goals = append(goals, &pb.Goal{
				Id:            model.ID,
				UserId:        model.UserID,
				Name:          model.Name,
				TargetAmount:  model.TargetAmount,
				CurrentAmount: model.CurrentAmount,
				Deadline:      model.Deadline,
				Status:        model.Status,
				CreatedAt:     model.CreatedAt,
				UpdatedAt:     model.UpdatedAt,
			})
		}
		if err = cursor.Err(); err != nil {
			g.log.Error("Error while iterating over cursor", logger.Error(err))
			return nil, err
		}

		return &pb.Goals{Goals: goals}, nil
	}

	func (g *goalRepo) Update(ctx context.Context, req *pb.Goal) (*pb.Goal, error) {
		objectID, err := primitive.ObjectIDFromHex(req.GetId())
		if err != nil {
			g.log.Error("Invalid ID format")
			return nil, err
		}

		filter := bson.M{"_id": objectID}

		update := bson.M{
			"$set": bson.M{
				"user_id":       req.UserId,
				"name":          req.Name,
				"target_amount": req.TargetAmount,
				"current_amount": req.CurrentAmount,
				"deadline":      req.Deadline,
				"status":        req.Status,
				"updated_at":    time.Now().Format(time.RFC3339),
			},
		}

		_, err = g.db.Collection("goals").UpdateOne(ctx, filter, update)
		if err != nil {
			g.log.Error("Error while updating goal", logger.Error(err))
			return nil, err
		}

		return &pb.Goal{
			Id:            req.Id,
			UserId:        req.UserId,
			Name:          req.Name,
			TargetAmount:  req.TargetAmount,
			CurrentAmount: req.CurrentAmount,
			Deadline:      req.Deadline,
			Status:        req.Status,
			CreatedAt:     req.CreatedAt,
			UpdatedAt:     time.Now().Format(time.RFC3339),
		}, nil
	}

	func (g *goalRepo) Delete(ctx context.Context, req *pb.PrimaryKey) error {
		objectID, err := primitive.ObjectIDFromHex(req.GetId())
		if err != nil {
			g.log.Error("Invalid ID format")
			return err
		}

		filter := bson.M{"_id": objectID}

		_, err = g.db.Collection("goals").DeleteOne(ctx, filter)
		if err != nil {
			g.log.Error("Error while deleting goal", logger.Error(err))
			return err
		}

		return nil
	}
