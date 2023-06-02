package repository

import (
	"context"
	"database/sql"
	"dexyarya123/belajar_golang_restful_api/model/domain"
)

type CategoryRepository interface {
	Save(ctx context.Context, tx *sql.Tx, category domain.Category) (domain.Category, error)
	Update(ctx context.Context, tx *sql.Tx, category domain.Category) (domain.Category, error)
	Delete(ctx context.Context, tx *sql.Tx, categoryId int)
	FindById(ctx context.Context, tx *sql.Tx, idCategory int) (domain.Category, error)
	FindAll(ctx context.Context, tx *sql.Tx) ([]domain.Category, error)
}
