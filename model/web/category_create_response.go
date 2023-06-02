package web

type CategoryCreateResponse struct {
	Id   int `validate:"required"`
	Name string
}
