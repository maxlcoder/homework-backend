package service

import (
	"fmt"

	"github.com/casbin/casbin/v2"
	"github.com/maxlcoder/homework-backend/app/modules/core/model"
	base_model "github.com/maxlcoder/homework-backend/model"
	"gorm.io/gorm/clause"

	"github.com/maxlcoder/homework-backend/repository"
	"github.com/samber/lo"
	"gorm.io/gorm"
)

type RoleServiceInterface interface {
	Create(role *model.Role) (*model.Role, error)
	CreateWithMenus(role *model.Role, menus []model.Menu) (*model.Role, error)
	UpdateWithMenus(role *model.Role, menus []model.Menu) (*model.Role, error)
	GetById(id uint) (*model.Role, error)
	GetPageByFilter(modelFilter model.RoleFilter, pagination base_model.Pagination) (int64, []model.Role, error)
	Delete(role *model.Role) error
}

type RoleService struct {
	db          *gorm.DB
	enforcer    *casbin.Enforcer
	menuService MenuServiceInterface
}

func NewRoleService(db *gorm.DB, enforcer *casbin.Enforcer, menuService MenuServiceInterface) RoleServiceInterface {
	return &RoleService{
		db:          db,
		enforcer:    enforcer,
		menuService: menuService,
	}
}

func (u *RoleService) Create(role *model.Role) (*model.Role, error) {
	// 判断是否存在已经适用的名称
	filter := model.RoleFilter{
		Name: &role.Name,
	}

	cond := repository.ConditionScope{
		StructCond: filter,
	}

	findUser, _ := repository.NewBaseRepository[model.Role](u.db).FindBy(cond)
	if findUser != nil {
		return nil, fmt.Errorf("当前角色已存在，请检查")
	}
	err := repository.NewBaseRepository[model.Role](u.db).Create(role, nil)
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

	cond := repository.ConditionScope{
		StructCond: filter,
	}

	find, _ := repository.NewBaseRepository[model.Role](u.db).FindBy(cond)
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
	menuCount, err := repository.NewBaseRepository[model.Menu](u.db).CountBy(menuCond)
	if err != nil {
		return nil, fmt.Errorf("菜单参数校验失败，请检查")
	}
	if menuCount != int64(len(menus)) {
		return nil, fmt.Errorf("菜单参数校验失败，请检查")
	}

	// 启动事务
	err = u.db.Transaction(func(tx *gorm.DB) error {
		err := repository.NewBaseRepository[model.Role](u.db).Create(role, tx)
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
			u.db.Clauses(clause.OnConflict{
				DoNothing: true,
			}).Create(&roleMenus)
		}

		// 角色授权
		// 菜单关联的权限
		permissions, _ := u.GetPermissionsByMenuIds(lo.Map(menus, func(item model.Menu, index int) uint {
			return item.ID
		}))
		// 1. role_permission 表增加记录
		rolePermissions := lo.Map(permissions, func(item model.Permission, index int) model.RolePermission {
			return model.RolePermission{
				RoleID:       role.ID,
				PermissionID: item.ID,
			}
		})
		u.db.Clauses(clause.OnConflict{
			DoNothing: true,
		}).Create(rolePermissions)
		// 2. casbin 授权
		casbinPermission := lo.Map(permissions, func(item model.Permission, index int) []string {
			return []string{"1", item.PATH, item.Method}
		})
		if len(casbinPermission) > 0 {
			// 补充 casbin 中 p 规则，给角色赋权
			_, err := u.enforcer.AddPermissionsForUser("role_"+role.Name, casbinPermission...)
			if err != nil {
				return fmt.Errorf("角色Casbin授权失败: %w", err)
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return role, nil
}

func (u *RoleService) GetPermissionsByMenuIds(ids []uint) ([]model.Permission, error) {
	subQuery := u.db.Model(&model.MenuPermission{}).
		Select("permission_id").
		Where("menu_id IN (?)", ids)
	var permissions []model.Permission
	u.db.Where("id IN (?)", subQuery).Find(&permissions)
	return permissions, nil
}

func (u *RoleService) UpdateWithMenus(role *model.Role, menus []model.Menu) (*model.Role, error) {
	// 判断角色是否存在
	filter := model.RoleFilter{
		ID: &role.ID,
	}
	cond := repository.ConditionScope{
		StructCond: filter,
	}
	find, _ := repository.NewBaseRepository[model.Role](u.db).FindBy(cond)
	if find == nil {
		return nil, fmt.Errorf("当前角色不存在，请检查")
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
	menuCount, err := repository.NewBaseRepository[model.Menu](u.db).CountBy(menuCond)
	if err != nil {
		return nil, fmt.Errorf("菜单参数校验失败，请检查")
	}
	if menuCount != int64(len(menus)) {
		return nil, fmt.Errorf("菜单参数校验失败，请检查")
	}

	// 启动事务
	err = u.db.Transaction(func(tx *gorm.DB) error {
		err := repository.NewBaseRepository[model.Role](u.db).Update(role, tx)
		if err != nil {
			return fmt.Errorf("角色创建失败: %w", err)
		}
		// 菜单处理
		// 删除关联菜单
		deleteCond := repository.ConditionScope{
			Scopes: []func(*gorm.DB) *gorm.DB{
				func(db *gorm.DB) *gorm.DB {
					return db.Where("role_id = ?", role.ID)
				},
			},
		}
		err = repository.NewBaseRepository[model.RoleMenu](u.db).DeleteBy(deleteCond, tx)
		if err != nil {
			return fmt.Errorf("角色菜单处理失败: %w", err)
		}
		// 补充新菜单
		var roleMenus []model.RoleMenu
		roleMenus = lo.Map(menus, func(item model.Menu, index int) model.RoleMenu {
			return model.RoleMenu{
				RoleID: role.ID,
				MenuID: item.ID,
			}
		})
		if len(roleMenus) > 0 {
			u.db.Clauses(clause.OnConflict{
				DoNothing: true,
			}).Create(&roleMenus)
		}

		// 角色授权
		// 菜单关联的权限
		permissions, _ := u.menuService.GetPermissionsByMenuIds(lo.Map(menus, func(item model.Menu, index int) uint {
			return item.ID
		}))
		// 1. 先删除role_permission 表记录，再增加记录
		// 删除记录
		u.db.Where("role_id = ?", role.ID).Delete(&model.RolePermission{})

		rolePermissions := lo.Map(permissions, func(item model.Permission, index int) model.RolePermission {
			return model.RolePermission{
				RoleID:       role.ID,
				PermissionID: item.ID,
			}
		})
		u.db.Clauses(clause.OnConflict{
			DoNothing: true,
		}).Create(&rolePermissions)
		// 2. 先删除 casbin 授权，再添加
		u.enforcer.RemoveFilteredPolicy(0, "role_"+role.Name, "1")
		casbinPermission := lo.Map(permissions, func(item model.Permission, index int) []string {
			return []string{"1", item.PATH, item.Method}
		})
		if len(casbinPermission) > 0 {
			// 补充 casbin 中 p 规则，给角色赋权
			_, err := u.enforcer.AddPermissionsForUser("role_"+role.Name, casbinPermission...)
			if err != nil {
				return fmt.Errorf("角色Casbin授权失败: %w", err)
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return role, nil
}

func (u *RoleService) Delete(role *model.Role) error {

	// 检查角色是否存在
	role, err := repository.NewBaseRepository[model.Role](u.db).FindById(role.ID)
	if err != nil {
		return err
	}

	// 启动事务
	err = u.db.Transaction(func(tx *gorm.DB) error {
		// 1. 删除 role 表记录
		err := repository.NewBaseRepository[model.Role](u.db).DeleteById(role.ID, tx)
		if err != nil {
			return fmt.Errorf("角色删除失败: %w", err)
		}
		// 2. 删除关联菜单 role_menu 表记录
		deleteCond := repository.ConditionScope{
			Scopes: []func(*gorm.DB) *gorm.DB{
				func(db *gorm.DB) *gorm.DB {
					return db.Where("role_id = ?", role.ID)
				},
			},
		}
		err = repository.NewBaseRepository[model.RoleMenu](u.db).DeleteBy(deleteCond, tx)
		if err != nil {
			return fmt.Errorf("角色菜单删除失败: %w", err)
		}
		// 3. 删除 role_permission 表记录
		u.db.Where("role_id = ?", role.ID).Delete(&model.RolePermission{})
		// 4. 删除 casbin 记录
		u.enforcer.RemoveFilteredPolicy(0, "role_"+role.Name, "1")
		return nil
	})
	return err
}

func (u *RoleService) List() {
	//TODO implement me
	panic("implement me")
}

func (u *RoleService) GetById(id uint) (*model.Role, error) {
	filter := model.RoleFilter{
		ID: &id,
	}

	cond := repository.ConditionScope{
		StructCond: filter,
	}
	user, err := repository.NewBaseRepository[model.Role](u.db).FindBy(cond)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *RoleService) GetByObject() {
	//TODO implement me
	panic("implement me")
}

func (u *RoleService) GetPageByFilter(modelFilter model.RoleFilter, pagination base_model.Pagination) (int64, []model.Role, error) {

	cond := repository.ConditionScope{
		Scopes: []func(*gorm.DB) *gorm.DB{func(db *gorm.DB) *gorm.DB {
			if modelFilter.Name != nil {
				db.Where("name like ?", "%"+*modelFilter.Name+"%")
			}
			return db
		}},
		Order: []string{"created_at desc"},
	}
	total, users, err := repository.NewBaseRepository[model.Role](u.db).Page(cond, pagination)
	if err != nil {
		return 0, nil, fmt.Errorf("用户分页查询失败: %w", err)
	}
	return total, users, nil
}
