package service

import (
	"fmt"

	"github.com/maxlcoder/homework-backend/model"
	"github.com/maxlcoder/homework-backend/repository"
	"gorm.io/gorm"
)

type AdminServiceInterface interface {
	Create(admin *model.Admin) (*model.Admin, error)
	GetById(id uint) (*model.Admin, error)
	GetPageByFilter(modelFilter model.AdminFilter, pagination model.Pagination) (int64, []model.Admin, error)
}

type AdminService struct {
	db              *gorm.DB
	AdminRepository repository.AdminRepository
}

func NewAdminService(db *gorm.DB, adminRepository repository.AdminRepository) *AdminService {
	return &AdminService{
		db:              db,
		AdminRepository: adminRepository,
	}
}

func (u *AdminService) Create(admin *model.Admin) (*model.Admin, error) {
	// 判断是否存在已经适用的名称
	userFiler := model.UserFilter{
		Name: &admin.Name,
	}
	cond := repository.StructCondition[model.UserFilter]{
		Cond: userFiler,
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

func (u *AdminService) Update() {
	//TODO implement me
	panic("implement me")
}

func (u *AdminService) Delete() {
	//TODO implement me
	panic("implement me")
}

func (u *AdminService) List() {
	//TODO implement me
	panic("implement me")
}

func (u *AdminService) GetById(id uint) (*model.Admin, error) {
	userFiler := model.UserFilter{
		ID: &id,
	}
	cond := repository.StructCondition[model.UserFilter]{
		Cond: userFiler,
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
