package repository

import (
	"github.com/maxlcoder/homework-backend/model"
	"gorm.io/gorm"
)

type Repository[T any] interface {
	Create(entity *T, tx *gorm.DB) error
	CreateBatch(entity []*T, tx *gorm.DB) error
	FindById(id uint) (*T, error)
	Update(entity *T, tx *gorm.DB) error
	DeleteById(id uint, tx *gorm.DB) error
	DeleteBy(cond ConditionScope, tx *gorm.DB) error
	Page(cond ConditionScope, pagination model.Pagination) (int64, []T, error)
	FindBy(cond ConditionScope) (*T, error)
	CountBy(cond ConditionScope) (int64, error)
}

func First[T any, PT interface {
	*T
	model.Authenticatable
}](db *gorm.DB) (PT, error) {
	var t T
	ptr := any(&t).(PT)
	err := db.First(&ptr).Error
	if err != nil {
		return nil, err
	}
	return ptr, nil
}
