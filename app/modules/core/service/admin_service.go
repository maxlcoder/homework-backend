package service

import (
	"fmt"

	"github.com/maxlcoder/homework-backend/model"
	"github.com/maxlcoder/homework-backend/repository"
	"github.com/samber/lo"
	"gorm.io/gorm"
)

type AdminServiceInterface interface {
	Create(admin *model.Admin) (*model.Admin, error)
	CreateWithRoles(admin *model.Admin, roles []model.Role) (*model.Admin, error)
	UpdateWithRoles(admin *model.Admin, roles []model.Role) (*model.Admin, error)
	GetById(id uint) (*model.Admin, error)
	GetPageByFilter(modelFilter model.AdminFilter, pagination model.Pagination) (int64, []model.Admin, error)
	Delete(admin *model.Admin) error
}

type AdminService struct {
	db                  *gorm.DB
	AdminRepository     repository.AdminRepository
	RoleRepository      repository.RoleRepository
	AdminRoleRepository repository.AdminRoleRepository
}

func NewAdminService(db *gorm.DB, adminRepository repository.AdminRepository, roleRepository repository.RoleRepository, adminRoleRepository repository.AdminRoleRepository) AdminServiceInterface {
	return &AdminService{
		db:                  db,
		AdminRepository:     adminRepository,
		RoleRepository:      roleRepository,
		AdminRoleRepository: adminRoleRepository,
	}
}

func (u *AdminService) Create(admin *model.Admin) (*model.Admin, error) {
	// 判断是否存在已经适用的名称
	userFiler := model.UserFilter{
		Name: &admin.Name,
	}
	cond := repository.ConditionScope{
		StructCond: userFiler,
	}
	findUser, _ := u.AdminRepository.FindBy(cond)
	if findUser != nil {
		return nil, fmt.Errorf("当前用户名不可用，请检查")
	}
	err := u.AdminRepository.Create(admin, nil)
	if err != nil {
		return nil, fmt.Errorf("用户创建失败: %w", err)
	}
	return admin, nil
}

func (u *AdminService) CreateWithRoles(admin *model.Admin, roles []model.Role) (*model.Admin, error) {

	// 判断是否存在已经适用的名称
	filter := model.AdminFilter{
		Name: &admin.Name,
	}

	cond := repository.ConditionScope{
		StructCond: filter,
	}

	find, _ := u.AdminRepository.FindBy(cond)
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
	roleCount, err := u.RoleRepository.CountBy(roleCond)
	if err != nil {
		return nil, fmt.Errorf("角色参数校验失败，请检查")
	}
	if roleCount != int64(len(roles)) {
		return nil, fmt.Errorf("角色参数校验失败，请检查")
	}
	// admin 已经配置了关联，会自动 处理 roles 和 admin_roles 表
	err = u.AdminRepository.Create(admin, nil)
	if err != nil {
		return nil, fmt.Errorf("账号创建失败: %w", err)
	}
	return admin, nil
}

func (u *AdminService) UpdateWithRoles(admin *model.Admin, roles []model.Role) (*model.Admin, error) {

	// 判断是否存在已经适用的名称
	filter := model.AdminFilter{
		Name: &admin.Name,
	}

	cond := repository.ConditionScope{
		StructCond: filter,
		Scopes: []func(*gorm.DB) *gorm.DB{
			func(db *gorm.DB) *gorm.DB {
				db.Not(&model.Admin{
					BaseModel: model.BaseModel{
						ID: admin.ID,
					},
				})
				return db
			},
		},
	}

	find, _ := u.AdminRepository.FindBy(cond)
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
	roleCount, err := u.RoleRepository.CountBy(roleCond)
	if err != nil {
		return nil, fmt.Errorf("角色参数校验失败，请检查")
	}
	if roleCount != int64(len(roles)) {
		return nil, fmt.Errorf("角色参数校验失败，请检查")
	}
	// admin 已经配置了关联，会自动 处理 roles 和 admin_roles 表
	err = u.db.Transaction(func(tx *gorm.DB) error {
		// 删除之前的 admin_roles 关联
		roleCond := repository.ConditionScope{
			Scopes: []func(*gorm.DB) *gorm.DB{
				func(db *gorm.DB) *gorm.DB {
					return db.Where("admin_id = ?", admin.ID)
				},
			},
		}
		u.AdminRoleRepository.DeleteBy(roleCond, tx)
		err = u.AdminRepository.Update(admin, tx)
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

func (u *AdminService) Update() {
	//TODO implement me
	panic("implement me")
}

func (u *AdminService) Delete(admin *model.Admin) error {
	err := u.db.Transaction(func(tx *gorm.DB) error {
		// 删除 admins
		u.AdminRepository.DeleteById(admin.ID, tx)
		// 删除 admin_roles
		cond := repository.ConditionScope{
			StructCond: admin,
		}
		u.AdminRoleRepository.DeleteBy(cond, tx)
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
func (u *AdminService) List() {
	//TODO implement me
	panic("implement me")
}

func (u *AdminService) GetById(id uint) (*model.Admin, error) {
	filter := model.AdminFilter{
		ID: &id,
	}
	cond := repository.ConditionScope{
		StructCond: filter,
		Preloads: []string{
			"Roles",
		},
	}
	user, err := u.AdminRepository.FindBy(cond)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *AdminService) GetByObject() {
	//TODO implement me
	panic("implement me")
}

func (u *AdminService) GetPageByFilter(filter model.AdminFilter, pagination model.Pagination) (int64, []model.Admin, error) {

	cond := repository.ConditionScope{
		Preloads: []string{
			"Roles",
		},
		Scopes: []func(*gorm.DB) *gorm.DB{
			func(db *gorm.DB) *gorm.DB {
				if filter.Name != nil {
					return repository.LikeScope("name", *filter.Name)(db)
				} else {
					return db
				}

			},
		},
	}
	total, admins, err := u.AdminRepository.Page(cond, pagination)
	if err != nil {
		return 0, nil, fmt.Errorf("用户分页查询失败: %w", err)
	}
	return total, admins, nil
}
