package service

import (
	"context"
	"database/sql"
	repository "dexyarya123/belajar_golang_restful_api/Repository"
	"dexyarya123/belajar_golang_restful_api/helper"
	"dexyarya123/belajar_golang_restful_api/model/domain"
	"dexyarya123/belajar_golang_restful_api/model/web"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type CategoryServiceImplementation struct {
	repo     repository.CategoryRepository
	DB       *sql.DB
	Validate *validator.Validate
}

func NewCategoryService(repo repository.CategoryRepository, DB *sql.DB, Validate *validator.Validate) CategoryService {
	return &CategoryServiceImplementation{}
}

func (service *CategoryServiceImplementation) CreateCategory(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	categori := domain.Category{
		Name: request.Name,
	}

	categori, err = service.repo.Save(ctx, tx, categori)
	helper.PanicIfError(err)
	return helper.ToCategoryResponse(categori)
}

func (service *CategoryServiceImplementation) UpdateCategory(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	categori := domain.Category{
		Id:   request.Id,
		Name: request.Name,
	}

	categori, err = service.repo.FindById(ctx, tx, request.Id)
	if err != nil {
		helper.PanicIfError(err)
	}
	categori.Name = request.Name
	categori, err = service.repo.Update(ctx, tx, categori)
	helper.PanicIfError(err)
	return helper.ToCategoryResponse(categori)
}

func (service *CategoryServiceImplementation) DeleteCategory(ctx context.Context, categoryId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	category, err := service.repo.FindById(ctx, tx, categoryId)
	if err != nil {
		helper.PanicIfError(err)
	}
	service.repo.Delete(ctx, tx, category.Id)

}

func (service *CategoryServiceImplementation) FindByIdCategory(ctx context.Context, id int) web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	category, err := service.repo.FindById(ctx, tx, id)
	helper.PanicIfError(err)
	return helper.ToCategoryResponse(category)

}
func (service *CategoryServiceImplementation) FindAllCategory(ctx context.Context) ([]web.CategoryResponse, error) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	categories, err := service.repo.FindAll(ctx, tx)
	helper.PanicIfError(err)
	fmt.Println("ini kedua")
	return helper.ToCategoryResponses(categories), nil

}
