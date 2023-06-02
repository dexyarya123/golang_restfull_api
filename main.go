package main

import (
	_ "github.com/go-sql-driver/mysql"

	repository "dexyarya123/belajar_golang_restful_api/Repository"
	"dexyarya123/belajar_golang_restful_api/app"
	"dexyarya123/belajar_golang_restful_api/controller"
	"dexyarya123/belajar_golang_restful_api/helper"
	"dexyarya123/belajar_golang_restful_api/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func main() {
	db := app.GetConnectionDb()
	validate := validator.New()
	CategoryRepository := repository.NewCategoryRepository()
	CategoryService := service.NewCategoryService(CategoryRepository, db, validate)
	CategoryController := controller.NewCategoryController(CategoryService)

	router := httprouter.New()
	router.GET("/api/categories", CategoryController.FindAllCategory)
	router.GET("/api/categoies/:categoryId", CategoryController.FindByIdCategory)
	router.POST("/api/categories/create", CategoryController.CreateCategory)
	router.PUT("/api/categories/:categoryId", CategoryController.UpdateCategory)
	router.DELETE("/api/categories/:categoryId", CategoryController.DeleteCategory)
	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
