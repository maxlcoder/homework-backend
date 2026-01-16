package repository

import (
	"github.com/maxlcoder/homework-backend/app/modules/core/model"
	"github.com/maxlcoder/homework-backend/repository"
	"gorm.io/gorm"
)

type MenuRepository interface {
	repository.Repository[model.Menu]
	// 自定义方法
	GetPermissionsByMenuId(id uint) ([]model.Permission, error)
	GetPermissionsByMenuIds(ids []uint) ([]model.Permission, error)
	GetMenusByRoleId(roleId uint) ([]model.Menu, error)
}

type MenuRepositoryImpl struct {
	*repository.BaseRepository[model.Menu]
}

func NewMenuRepository(db *gorm.DB) MenuRepository {
	return &MenuRepositoryImpl{
		repository.NewBaseRepository[model.Menu](db),
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

func (r *MenuRepositoryImpl) GetMenusByRoleId(roleId uint) ([]model.Menu, error) {
	var menus []model.Menu
	subQuery := r.DB.Model(&model.RoleMenu{}).
		Select("1").
		Where("role_menus.menu_id = menus.id").
		Where("role_id = ?", roleId)

	err := r.DB.Where("EXISTS (?)", subQuery).Find(&menus).Error
	return menus, err
}
