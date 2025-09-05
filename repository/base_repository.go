package repository

import (
	"github.com/maxlcoder/homework-backend/model"
	"gorm.io/gorm"
)

// BaseRepository 通用仓库
type BaseRepository[T any] struct {
	db *gorm.DB
}

func NewBaseRepository[T any](db *gorm.DB) *BaseRepository[T] {
	return &BaseRepository[T]{
		db: db,
	}
}

// 实现基础方法
func (r *BaseRepository[T]) Create(entity *T) error {
	return r.db.Create(entity).Error
}

func (r *BaseRepository[T]) FindById(id uint) (*T, error) {
	var entity T
	if err := r.db.First(&entity, id).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *BaseRepository[T]) Update(entity *T) error {
	return r.db.Save(entity).Error
}

func (r *BaseRepository[T]) DeleteById(id uint) error {
	var entity T
	return r.db.Delete(&entity, id).Error
}

// 查询条件 where , 分页条件 paginationQuery
func (r *BaseRepository[T]) Page(cond QueryCondition[T], pagination model.Pagination) (int64, []T, error) {
	var entity T
	var entities []T
	var total int64 // gorm 默认总数使用 int64
	query := cond.Apply(r.db.Model(&entity))
	// 统计总数
	if err := query.Count(&total).Error; err != nil {
		return 0, nil, err
	}
	// 分页查询
	if err := query.Offset((pagination.Page - 1) * pagination.PerPage).
		Limit(pagination.PerPage).
		Find(&entities).Error; err != nil {
		return 0, nil, err
	}
	return total, entities, nil
}

func (r *BaseRepository[T]) FindBy(cond QueryCondition[T]) (*T, error) {
	var entity T
	query := cond.Apply(r.db.Model(&entity))
	if err := query.First(&entity).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}
