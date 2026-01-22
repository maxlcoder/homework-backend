package service

import (
	"fmt"

	"github.com/maxlcoder/homework-backend/app/modules/core/model"
	base_model "github.com/maxlcoder/homework-backend/model"
	"github.com/maxlcoder/homework-backend/repository"
	"gorm.io/gorm"
)

type UserServiceInterface interface {
	Create(*model.User) (*model.User, error)
	GetById(id uint) (*model.User, error)
	GetPageByFilter(modelFilter model.UserFilter, pagination base_model.Pagination) (int64, []model.User, error)
}

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) UserServiceInterface {
	return &UserService{
		db: db,
	}
}

func (u *UserService) Create(user *model.User) (*model.User, error) {
	// 判断是否存在已经适用的名称
	userFiler := model.UserFilter{
		Name: &user.Name,
	}
	cond := repository.ConditionScope{
		StructCond: userFiler,
	}
	findUser, _ := repository.NewBaseRepository[model.User](u.db).FindBy(cond)
	if findUser != nil {
		return nil, fmt.Errorf("当前用户名不可用，请检查")
	}
	err := repository.NewBaseRepository[model.User](u.db).Create(user, nil)
	if err != nil {
		return nil, fmt.Errorf("用户创建失败: %w", err)
	}
	return user, nil
}

func (u *UserService) Update() {
	//TODO implement me
	panic("implement me")
}

func (u *UserService) Delete() {
	//TODO implement me
	panic("implement me")
}

func (u *UserService) List() {
	//TODO implement me
	panic("implement me")
}

func (u *UserService) GetById(id uint) (*model.User, error) {
	userFiler := model.UserFilter{
		ID: &id,
	}
	cond := repository.ConditionScope{
		StructCond: userFiler,
	}
	user, err := repository.NewBaseRepository[model.User](u.db).FindBy(cond)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u UserService) GetByObject() {
	//TODO implement me
	panic("implement me")
}

func (u UserService) GetPageByFilter(modelFilter model.UserFilter, pagination base_model.Pagination) (int64, []model.User, error) {
	cond := repository.ConditionScope{
		StructCond: modelFilter,
	}
	total, users, err := repository.NewBaseRepository[model.User](u.db).Page(cond, pagination)
	if err != nil {
		return 0, nil, fmt.Errorf("用户分页查询失败: %w", err)
	}
	return total, users, nil
}
