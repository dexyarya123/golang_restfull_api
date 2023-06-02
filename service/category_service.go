package service

import (
	"context"
	"dexyarya123/belajar_golang_restful_api/model/web"
)

type CategoryService interface {
	CreateCategory(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse
	UpdateCategory(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse
	DeleteCategory(ctx context.Context, categoryId int)
	FindByIdCategory(ctx context.Context, categoryId int) web.CategoryResponse
	FindAllCategory(ctx context.Context) ([]web.CategoryResponse, error)
}
