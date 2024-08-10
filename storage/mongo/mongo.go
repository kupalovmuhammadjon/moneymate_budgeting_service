package mongo

import (
	"budgeting_service/configs"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func ConnectMongoDB(ctx context.Context, cfg *configs.Config) (*mongo.Database, error) {

	url := fmt.Sprintf(`mongo://%s:%s`, cfg.MongoDBHost, cfg.MongoDBPort)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, readpref.Nearest())

	if err != nil {
		return nil, err
	}

	db := client.Database(cfg.MongoDBName)
	
	return db, nil
}
