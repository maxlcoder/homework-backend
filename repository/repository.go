package repository

import (
	"github.com/maxlcoder/homework-backend/model"
	"gorm.io/gorm"
)

func Paginate[T any](db *gorm.DB, paginationQuery model.PaginationQuery, result *[]T) (int64, error) {

	var total int64 // gorm 默认总数使用 int64
	// 统计总数
	if err := db.Count(&total).Error; err != nil {
		return 0, err
	}
	// 分页查询
	if err := db.Limit(paginationQuery.PerPage).Offset((paginationQuery.Page - 1) * paginationQuery.PerPage).Find(result).Error; err != nil {
		return 0, err
	}
	return total, nil
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
