package repository

import (
	"context"
	"database/sql"

	"go-database/entity"
)

type commentRepositoryImpl struct {
	DB *sql.DB
}

func (r *commentRepositoryImpl) Insert(ctx context.Context, req entity.Comment) error {

	return nil
}
