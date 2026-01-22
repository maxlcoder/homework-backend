package service

import (
	"fmt"

	"github.com/maxlcoder/homework-backend/app/modules/core/model"
	base_model "github.com/maxlcoder/homework-backend/model"
	"github.com/maxlcoder/homework-backend/repository"
	"gorm.io/gorm"
)

type MenuServiceInterface interface {
	Create(menu *model.Menu) (*model.Menu, error)
	GetById(id uint) (*model.Menu, error)
	GetPageByFilter(modelFilter model.MenuFilter, pagination base_model.Pagination) (int64, []model.Menu, error)
	//GetPermissionsByMenuId(id uint) ([]model.Permission, error)
	GetPermissionsByMenuIds(ids []uint) ([]model.Permission, error)
	//GetMenusByRoleId(roleId uint) ([]model.Menu, error)
}

type MenuService struct {
	db *gorm.DB
}

func NewMenuService(db *gorm.DB) MenuServiceInterface {
	return &MenuService{
		db: db,
	}
}

func (u *MenuService) Create(menu *model.Menu) (*model.Menu, error) {
	// 判断是否存在已经适用的名称
	filter := model.MenuFilter{
		Name: &menu.Name,
	}
	cond := repository.ConditionScope{
		StructCond: filter,
	}
	findUser, _ := repository.NewBaseRepository[model.Menu](u.db).FindBy(cond)
	if findUser != nil {
		return nil, fmt.Errorf("当前用户名不可用，请检查")
	}
	err := repository.NewBaseRepository[model.Menu](u.db).Create(menu, nil)
	if err != nil {
		return nil, fmt.Errorf("用户创建失败: %w", err)
	}
	return menu, nil
}

func (u *MenuService) Update() {
	//TODO implement me
	panic("implement me")
}

func (u *MenuService) Delete() {
	//TODO implement me
	panic("implement me")
}

func (u *MenuService) List() {
	//TODO implement me
	panic("implement me")
}

func (u *MenuService) GetById(id uint) (*model.Menu, error) {
	filter := model.MenuFilter{
		ID: &id,
	}
	cond := repository.ConditionScope{
		StructCond: filter,
	}
	user, err := repository.NewBaseRepository[model.Menu](u.db).FindBy(cond)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *MenuService) GetByObject() {
	//TODO implement me
	panic("implement me")
}

func (u *MenuService) GetPageByFilter(modelFilter model.MenuFilter, pagination base_model.Pagination) (int64, []model.Menu, error) {

	cond := repository.ConditionScope{
		StructCond: modelFilter,
	}

	total, menus, err := repository.NewBaseRepository[model.Menu](u.db).Page(cond, pagination)
	if err != nil {
		return 0, nil, fmt.Errorf("用户分页查询失败: %w", err)
	}
	return total, menus, nil
}

func (u *MenuService) GetPermissionsByMenuIds(ids []uint) ([]model.Permission, error) {
	subQuery := u.db.Model(&model.MenuPermission{}).
		Select("permission_id").
		Where("menu_id IN (?)", ids)
	var permissions []model.Permission
	u.db.Where("id IN (?)", subQuery).Find(&permissions)
	return permissions, nil
}
