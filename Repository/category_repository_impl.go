package repository

import (
	"context"
	"database/sql"
	"dexyarya123/belajar_golang_restful_api/helper"
	"dexyarya123/belajar_golang_restful_api/model/domain"
	"errors"
	"fmt"
)

type CateogoryRepositoryImplementation struct {
}

func NewCategoryRepository() CategoryRepository {
	return &CateogoryRepositoryImplementation{}
}

func (repository *CateogoryRepositoryImplementation) Save(ctx context.Context, tx *sql.Tx, category domain.Category) (domain.Category, error) {
	querySql := "INSERT INTO category(name) VALUES(?) "
	result, err := tx.ExecContext(ctx, querySql, category.Name)
	helper.PanicIfError(err)
	id, err := result.LastInsertId()
	helper.PanicIfError(err)
	category.Id = int(id)
	return category, nil
}
func (repository *CateogoryRepositoryImplementation) Delete(ctx context.Context, tx *sql.Tx, id int) {
	querySql := "DELETE category WHERE id= ?"
	_, err := tx.ExecContext(ctx, querySql, id)
	helper.PanicIfError(err)
}
func (repository *CateogoryRepositoryImplementation) Update(ctx context.Context, tx *sql.Tx, category domain.Category) (domain.Category, error) {
	querySql := "UPDATE category SET name = ? WHERE id= ?"
	_, err := tx.ExecContext(ctx, querySql, category.Name, category.Id)
	helper.PanicIfError(err)
	return category, nil
}

func (repository *CateogoryRepositoryImplementation) FindById(ctx context.Context, tx *sql.Tx, id int) (domain.Category, error) {
	querySql := "SELECT id,name FROM category WHERE id = ?"
	rows, err := tx.QueryContext(ctx, querySql, id)
	helper.PanicIfError(err)
	category := domain.Category{}
	if rows.Next() {
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		return category, nil
	} else {
		return category, errors.New("category Not Found")
	}

}

func (repository *CateogoryRepositoryImplementation) FindAll(ctx context.Context, tx *sql.Tx) ([]domain.Category, error) {

	querySql := "SELECT * FROM category"
	rows, err := tx.QueryContext(ctx, querySql)
	fmt.Println(err)
	helper.PanicIfError(err)
	var categories []domain.Category
	for rows.Next() {
		category := domain.Category{}
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		categories = append(categories, category)

	}
	return categories, nil
}
