package helper

import (
	"dexyarya123/belajar_golang_restful_api/model/domain"
	"dexyarya123/belajar_golang_restful_api/model/web"
)

func ToCategoryResponses(categories []domain.Category) []web.CategoryResponse {

	var categoryResponses []web.CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, ToCategoryResponse(category))
	}
	return categoryResponses
}
