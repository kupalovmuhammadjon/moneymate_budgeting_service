package storage

import (
	"budgeting_service/configs"
	"budgeting_service/pkg/logger"
	"context"
)


type IStorage interface{
	Close()
}

type storage struct{

}

func NewIStorage(ctx context.Context, cfg *configs.Config, log logger.ILogger) (IStorage, error){
	return &storage{

	}, nil
}

func (s *storage) Close(){

}
