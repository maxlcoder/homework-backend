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
func (r *BaseRepository[T]) getDB(tx *gorm.DB) *gorm.DB {
	if tx != nil {
		return tx
	}
	return r.db
}

func (r *BaseRepository[T]) Create(entity *T, tx *gorm.DB) error {
	return r.getDB(tx).Create(entity).Error
}

func (r *BaseRepository[T]) CreateBatch(entities []*T, tx *gorm.DB) error {
	return r.getDB(tx).CreateInBatches(entities, 100).Error
}

func (r *BaseRepository[T]) FindById(id uint) (*T, error) {
	var entity T
	if err := r.db.First(&entity, id).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *BaseRepository[T]) Update(entity *T, tx *gorm.DB) error {
	return r.getDB(tx).Updates(entity).Error
}

func (r *BaseRepository[T]) DeleteById(id uint, tx *gorm.DB) error {
	var entity T
	return r.getDB(tx).Delete(&entity, id).Error
}

func (r *BaseRepository[T]) DeleteBy(cond ConditionScope, tx *gorm.DB) error {
	var entity T
	query := cond.Apply(r.getDB(tx))
	return query.Delete(&entity).Error
}

// 查询条件 where , 分页条件 paginationQuery
func (r *BaseRepository[T]) Page(cond ConditionScope, pagination model.Pagination) (int64, []T, error) {
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

func (r *BaseRepository[T]) FindBy(cond ConditionScope) (*T, error) {
	var entity T
	query := cond.Apply(r.db.Model(&entity))
	if err := query.First(&entity).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *BaseRepository[T]) CountBy(cond ConditionScope) (int64, error) {
	var entity T
	var count int64
	query := cond.Apply(r.db.Model(&entity))
	if err := query.Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
