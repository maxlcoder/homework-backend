package repository

import "github.com/maxlcoder/homework-backend/model"

type Repository[T any] interface {
	Create(entity *T) error
	FindById(id uint) (*T, error)
	Update(entity *T) error
	DeleteById(id uint) error
	Page(cond QueryCondition[T], pagination model.Pagination) (int64, []T, error)
}
