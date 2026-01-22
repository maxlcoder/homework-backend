package service

import (
	"fmt"

	"github.com/maxlcoder/homework-backend/app/modules/core/admin/request"
	"github.com/maxlcoder/homework-backend/app/modules/core/model"
	base_model "github.com/maxlcoder/homework-backend/model"
	"github.com/maxlcoder/homework-backend/repository"
	"github.com/samber/lo"
	"gorm.io/gorm"
)

type AdminServiceInterface interface {
	Page(pageRequest request.AdminPageRequest) ([]model.Admin, int64, error)
	Create(admin *model.Admin, roles []model.Role) (*model.Admin, error)
	Update(admin *model.Admin, roles []model.Role) (*model.Admin, error)
	Delete(id uint) error
	FindById(id uint) (*model.Admin, error)
	GetMenusByRoleId(roleId uint) ([]*model.Menu, error)
	GetMenusWithChildrenByRoleId(roleId uint) ([]*model.Menu, error)
}

type AdminService struct {
	db *gorm.DB
}

func NewAdminService(db *gorm.DB) AdminServiceInterface {
	return &AdminService{
		db: db,
	}
}

func (u *AdminService) Page(pageRequest request.AdminPageRequest) ([]model.Admin, int64, error) {
	cond := repository.ConditionScope{
		Preloads: []string{
			"Roles",
		},
		Scopes: []func(*gorm.DB) *gorm.DB{
			func(db *gorm.DB) *gorm.DB {
				if pageRequest.Name != nil && len(*pageRequest.Name) > 0 {
					return repository.LikeScope("name", *pageRequest.Name)(db)
				} else {
					return db
				}
			},
		},
	}

	// 创建分页参数
	pagination := base_model.Pagination{
		Page:    pageRequest.Page,
		PerPage: pageRequest.PerPage,
	}

	// 查询数据
	count, admins, err := repository.NewBaseRepository[model.Admin](u.db).Page(cond, pagination)
	if err != nil {
		return nil, 0, fmt.Errorf("获取管理员列表失败: %w", err)
	}

	return admins, count, nil
}

func (u *AdminService) Create(admin *model.Admin, roles []model.Role) (*model.Admin, error) {
	// 判断是否存在已经适用的名称
	filter := model.AdminFilter{
		Name: &admin.Name,
	}

	cond := repository.ConditionScope{
		StructCond: filter,
	}

	find, _ := repository.NewBaseRepository[model.Admin](u.db).FindBy(cond)
	if find != nil {
		return nil, fmt.Errorf("当前账号名称已存在，请检查")
	}

	// 角色校验
	roleCond := repository.ConditionScope{
		Scopes: []func(*gorm.DB) *gorm.DB{
			func(db *gorm.DB) *gorm.DB {
				return db.Where("id IN ?", lo.Map(roles, func(item model.Role, index int) uint {
					return item.ID
				}))
			},
		},
	}
	roleCount, err := repository.NewBaseRepository[model.Admin](u.db).CountBy(roleCond)
	if err != nil {
		return nil, fmt.Errorf("角色参数校验失败，请检查")
	}
	if roleCount != int64(len(roles)) {
		return nil, fmt.Errorf("角色参数校验失败，请检查")
	}

	err = repository.NewBaseRepository[model.Admin](u.db).Create(admin, nil)
	if err != nil {
		return nil, fmt.Errorf("账号创建失败: %w", err)
	}

	return admin, nil
}

func (u *AdminService) Update(admin *model.Admin, roles []model.Role) (*model.Admin, error) {
	// 判断是否存在已经适用的名称（排除自身）
	filter := model.AdminFilter{
		Name: &admin.Name,
	}

	cond := repository.ConditionScope{
		StructCond: filter,
		Scopes: []func(*gorm.DB) *gorm.DB{
			func(db *gorm.DB) *gorm.DB {
				db.Not(&model.Admin{
					BaseModel: base_model.BaseModel{
						ID: admin.ID,
					},
				})
				return db
			},
		},
	}

	find, _ := repository.NewBaseRepository[model.Admin](u.db).FindBy(cond)
	if find != nil {
		return nil, fmt.Errorf("当前账号名称已被占用，请检查")
	}

	// 角色校验
	roleCond := repository.ConditionScope{
		Scopes: []func(*gorm.DB) *gorm.DB{
			func(db *gorm.DB) *gorm.DB {
				return db.Where("id IN ?", lo.Map(roles, func(item model.Role, index int) uint {
					return item.ID
				}))
			},
		},
	}
	roleCount, err := repository.NewBaseRepository[model.Admin](u.db).CountBy(roleCond)
	if err != nil {
		return nil, fmt.Errorf("角色参数校验失败，请检查")
	}
	if roleCount != int64(len(roles)) {
		return nil, fmt.Errorf("角色参数校验失败，请检查")
	}

	err = u.db.Transaction(func(tx *gorm.DB) error {
		// 删除之前的 admin_roles 关联
		roleCond := repository.ConditionScope{
			Scopes: []func(*gorm.DB) *gorm.DB{
				func(db *gorm.DB) *gorm.DB {
					return db.Where("admin_id = ?", admin.ID)
				},
			},
		}
		repository.NewBaseRepository[model.Admin](u.db).DeleteBy(roleCond, tx)
		err = repository.NewBaseRepository[model.Admin](u.db).Update(admin, tx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("账号更新失败: %w", err)
	}

	return admin, nil
}

func (u *AdminService) Delete(id uint) error {
	err := u.db.Transaction(func(tx *gorm.DB) error {
		// 删除 admins
		repository.NewBaseRepository[model.Admin](u.db).DeleteById(id, tx)
		// 删除 admin_roles
		cond := repository.ConditionScope{
			StructCond: model.Admin{BaseModel: base_model.BaseModel{ID: id}},
		}
		repository.NewBaseRepository[model.Admin](u.db).DeleteBy(cond, tx)
		return nil
	})
	if err != nil {
		return fmt.Errorf("账号删除失败: %w", err)
	}
	return nil
}

func (u *AdminService) FindById(id uint) (*model.Admin, error) {
	filter := model.AdminFilter{
		ID: &id,
	}
	cond := repository.ConditionScope{
		StructCond: filter,
		Preloads: []string{
			"Roles",
		},
	}
	user, err := repository.NewBaseRepository[model.Admin](u.db).FindBy(cond)
	if err != nil {
		return nil, fmt.Errorf("账号查询失败: %w", err)
	}
	if user == nil {
		return nil, fmt.Errorf("账号不存在")
	}
	return user, nil
}

func (u *AdminService) GetMenusByRoleId(roleId uint) ([]*model.Menu, error) {
	var menus []model.Menu
	subQuery := u.db.Model(&model.RoleMenu{}).
		Select("1").
		Where("role_menus.menu_id = menus.id").
		Where("role_id = ?", roleId)

	err := u.db.Where("EXISTS (?)", subQuery).Find(&menus).Error
	if err != nil {
		return nil, err
	}
	return lo.Map(menus, func(menu model.Menu, index int) *model.Menu {
		return &menus[index]
	}), nil

}

func (u *AdminService) GetMenusWithChildrenByRoleId(roleId uint) ([]*model.Menu, error) {
	menus, _ := u.GetMenusByRoleId(roleId)
	if menus == nil {
		return make([]*model.Menu, 0), nil
	}
	return buildMenuTree(menus, 0), nil
}

func buildMenuTree(menus []*model.Menu, parentId uint) []*model.Menu {
	var tree []*model.Menu
	for _, menu := range menus {
		if menu.ParentID == parentId {
			// 递归，当前 menu 的菜单
			menu.Children = buildMenuTree(menus, menu.ID)
			tree = append(tree, menu)
		}
	}
	return tree
}
