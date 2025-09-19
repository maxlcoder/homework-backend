package service

import (
	"fmt"

	"github.com/maxlcoder/homework-backend/model"
	"github.com/maxlcoder/homework-backend/repository"
	"github.com/samber/lo"
	"gorm.io/gorm"
)

type RoleServiceInterface interface {
	Create(role *model.Role) (*model.Role, error)
	CreateWithMenus(role *model.Role, menus []model.Menu) (*model.Role, error)
	GetById(id uint) (*model.Role, error)
	GetPageByFilter(modelFilter model.RoleFilter, pagination model.Pagination) (int64, []model.Role, error)
}

type RoleService struct {
	db                 *gorm.DB
	RoleRepository     repository.RoleRepository
	MenuRepository     repository.MenuRepository
	RoleMenuRepository repository.RoleMenuRepository
}

func NewRoleService(db *gorm.DB, roleRepository repository.RoleRepository, menuRepository repository.MenuRepository, roleMenuRepository repository.RoleMenuRepository) *RoleService {
	return &RoleService{
		db:                 db,
		RoleRepository:     roleRepository,
		MenuRepository:     menuRepository,
		RoleMenuRepository: roleMenuRepository,
	}
}

func (u *RoleService) Create(role *model.Role) (*model.Role, error) {
	// 判断是否存在已经适用的名称
	filter := model.RoleFilter{
		Name: &role.Name,
	}

	cond := repository.StructCondition[model.RoleFilter]{
		Cond: filter,
	}

	findUser, _ := u.RoleRepository.FindBy(cond)
	if findUser != nil {
		return nil, fmt.Errorf("当前角色已存在，请检查")
	}
	err := u.RoleRepository.Create(role, nil)
	if err != nil {
		return nil, fmt.Errorf("用户创建失败: %w", err)
	}
	return role, nil
}

func (u *RoleService) CreateWithMenus(role *model.Role, menus []model.Menu) (*model.Role, error) {

	// 判断是否存在已经适用的名称
	filter := model.RoleFilter{
		Name: &role.Name,
	}

	cond := repository.StructCondition[model.RoleFilter]{
		Cond: filter,
	}

	find, _ := u.RoleRepository.FindBy(cond)
	if find != nil {
		return nil, fmt.Errorf("当前角色已存在，请检查")
	}

	// 菜单校验
	menuCond := repository.ConditionScope{
		Scopes: []func(*gorm.DB) *gorm.DB{
			func(db *gorm.DB) *gorm.DB {
				return db.Where("id IN ?", lo.Map(menus, func(item model.Menu, index int) uint {
					return item.ID
				}))
			},
		},
	}
	menuCount, err := u.MenuRepository.CountBy(menuCond)
	if err != nil {
		return nil, fmt.Errorf("菜单参数校验失败，请检查")
	}
	if menuCount != int64(len(menus)) {
		return nil, fmt.Errorf("菜单参数校验失败，请检查")
	}

	err = u.db.Transaction(func(tx *gorm.DB) error {
		// 启动事务
		err := u.RoleRepository.Create(role, tx)
		if err != nil {
			return fmt.Errorf("角色创建失败: %w", err)
		}

		// 菜单处理

		var roleMenus []model.RoleMenu
		roleMenus = lo.Map(menus, func(item model.Menu, index int) model.RoleMenu {
			return model.RoleMenu{
				RoleID: role.ID,
				MenuID: item.ID,
			}
		})
		if len(roleMenus) > 0 {
			err := u.RoleMenuRepository.UpsertCreateBatch(roleMenus, tx)
			if err != nil {
				return fmt.Errorf("角色菜单关联创建失败: %w", err)
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return role, nil
}

func (u *RoleService) Update() {
	//TODO implement me
	panic("implement me")
}

func (u *RoleService) Delete() {
	//TODO implement me
	panic("implement me")
}

func (u *RoleService) List() {
	//TODO implement me
	panic("implement me")
}

func (u *RoleService) GetById(id uint) (*model.Role, error) {
	filter := model.RoleFilter{
		ID: &id,
	}

	cond := repository.StructCondition[model.RoleFilter]{
		Cond: filter,
	}
	user, err := u.RoleRepository.FindBy(cond)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *RoleService) GetByObject() {
	//TODO implement me
	panic("implement me")
}

func (u *RoleService) GetPageByFilter(modelFilter model.RoleFilter, pagination model.Pagination) (int64, []model.Role, error) {

	cond := repository.ConditionScope{
		Scopes: []func(*gorm.DB) *gorm.DB{func(db *gorm.DB) *gorm.DB {
			if modelFilter.Name != nil {
				db.Where("name like ?", "%"+*modelFilter.Name+"%")
			}
			return db
		}},
		Order: []string{"created_at desc"},
	}
	total, users, err := u.RoleRepository.Page(cond, pagination)
	if err != nil {
		return 0, nil, fmt.Errorf("用户分页查询失败: %w", err)
	}
	return total, users, nil
}
