package helper

import (
	"dexyarya123/belajar_golang_restful_api/model/domain"
	"dexyarya123/belajar_golang_restful_api/model/web"
)

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}
