package controller

import (
	"dexyarya123/belajar_golang_restful_api/helper"
	"dexyarya123/belajar_golang_restful_api/model/web"
	"dexyarya123/belajar_golang_restful_api/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImplementation struct {
	CategoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) CategoryController {
	return &CategoryControllerImplementation{CategoryService: categoryService}
}

func (controller *CategoryControllerImplementation) CreateCategory(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	categoryCreateRequest := web.CategoryCreateRequest{}
	helper.ReadFromJsonBody(r, &categoryCreateRequest)
	categoryResponse := controller.CategoryService.CreateCategory(r.Context(), categoryCreateRequest)
	webResponse := web.CategoryWeb{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}
	helper.WriteFromJsonBody(w, webResponse)

}
func (controller *CategoryControllerImplementation) UpdateCategory(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	categoryUpdateRequest := web.CategoryUpdateRequest{}
	helper.ReadFromJsonBody(r, &categoryUpdateRequest)
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)
	categoryUpdateRequest.Id = id
	categoryResponse := controller.CategoryService.UpdateCategory(r.Context(), categoryUpdateRequest)
	webResponse := web.CategoryWeb{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}
	helper.WriteFromJsonBody(w, webResponse)

}
func (controller *CategoryControllerImplementation) DeleteCategory(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)
	controller.CategoryService.DeleteCategory(r.Context(), id)
	w.Header().Add("Content-Type", "application/json")

}
func (controller *CategoryControllerImplementation) FindAllCategory(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	categoryResponses, err := controller.CategoryService.FindAllCategory(r.Context())
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	webResponse := web.CategoryWeb{
		Code:   200,
		Status: "OK",
		Data:   categoryResponses,
	}

	helper.WriteFromJsonBody(w, webResponse)
	fmt.Println("ok")
}
func (controller *CategoryControllerImplementation) FindByIdCategory(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)
	categoryResponse := controller.CategoryService.FindByIdCategory(r.Context(), id)
	webResponse := web.CategoryWeb{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}
	helper.WriteFromJsonBody(w, webResponse)
}
