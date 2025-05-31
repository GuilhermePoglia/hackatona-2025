package services

import (
	"context"
	"database/sql"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

type BaseService struct {
	DB *sql.DB
}

func NewBaseService(db *sql.DB) *BaseService {
	boil.SetDB(db)

	return &BaseService{
		DB: db,
	}
}

func (s *BaseService) GetContext() context.Context {
	return context.Background()
}
