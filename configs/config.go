package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {

	BudgetingServiceGrpcHost string
	BudgetingServiceGrpcPort string

	MongoDBHost     string
	MongoDBPort     string
	MongoDBUser     string
	MongoDBName     string
	MongoDBPassword string

	ServiceName string
	LoggerLevel string
	LogPath     string
}

func Load() *Config {
	if err := godotenv.Load(".env"); err != nil {
		log.Printf(".env file not found: %s", err)
	}
	config := Config{}

	config.BudgetingServiceGrpcHost = cast.ToString(coalesce("BUDGETING_SERVICE_GRPC_HOST", "localhost"))
	config.BudgetingServiceGrpcPort = cast.ToString(coalesce("BUDGETING_SERVICE_GRPC_PORT", ":3333"))

	config.MongoDBHost = cast.ToString(coalesce("MONGODB_HOST", "localhost"))
	config.MongoDBPort = cast.ToString(coalesce("MONGODB_PORT", "5432"))
	config.MongoDBUser = cast.ToString(coalesce("MONGODB_USER", "mongo"))
	config.MongoDBName = cast.ToString(coalesce("MONGODB_NAME", "market_product_service"))
	config.MongoDBPassword = cast.ToString(coalesce("MONGODB_PASSWORD", ""))

	config.ServiceName = cast.ToString(coalesce("SERVICE_NAME", "auth_service"))
	config.LoggerLevel = cast.ToString(coalesce("LOGGER_LEVEL", "debug"))
	config.LogPath = cast.ToString(coalesce("LOG_PATH", "app.log"))

	return &config
}

func coalesce(key string, defaultValue interface{}) interface{} {
	if res, exists := os.LookupEnv(key); exists {
		return res
	}
	return defaultValue
}
