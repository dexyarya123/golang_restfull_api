package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type CategoryController interface {
	CreateCategory(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	UpdateCategory(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	DeleteCategory(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	FindAllCategory(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	FindByIdCategory(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}
