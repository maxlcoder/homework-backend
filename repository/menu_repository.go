package repository

import (
	"github.com/maxlcoder/homework-backend/model"
	"gorm.io/gorm"
)

type MenuRepository interface {
	Repository[model.Menu]
	// 自定义方法
	GetPermissionsByMenuId(id uint) ([]model.Permission, error)
	GetPermissionsByMenuIds(ids []uint) ([]model.Permission, error)
}

type MenuRepositoryImpl struct {
	*BaseRepository[model.Menu]
}

func NewMenuRepository(db *gorm.DB) MenuRepository {
	return &MenuRepositoryImpl{
		NewBaseRepository[model.Menu](db),
	}
}

func (r *MenuRepositoryImpl) GetPermissionsByMenuId(id uint) ([]model.Permission, error) {
	subQuery := r.DB.Model(&model.MenuPermission{}).
		Select("permission_id").
		Where("menu_id = ?", id)
	var permissions []model.Permission
	r.DB.Where("id IN (?)", subQuery).Find(&permissions)
	return permissions, nil
}

func (r *MenuRepositoryImpl) GetPermissionsByMenuIds(ids []uint) ([]model.Permission, error) {
	subQuery := r.DB.Model(&model.MenuPermission{}).
		Select("permission_id").
		Where("menu_id IN (?)", ids)
	var permissions []model.Permission
	r.DB.Where("id IN (?)", subQuery).Find(&permissions)
	return permissions, nil
}
