package service

import (
	"fmt"

	"github.com/maxlcoder/homework-backend/model"
	"github.com/maxlcoder/homework-backend/repository"
)

type UserServiceInterface interface {
	Create(*model.User) (*model.User, error)
	GetById(id uint) (*model.User, error)
	GetPageByFilter(modelFilter model.UserFilter, pagination model.Pagination) (int64, []model.User, error)
}

type UserService struct {
	UserRepository repository.UserRepository
}

func (u *UserService) Create(user *model.User) (*model.User, error) {
	// 判断是否存在已经适用的名称
	userFiler := model.UserFilter{
		Name: &user.Name,
	}
	cond := repository.StructCondition[model.UserFilter]{
		userFiler,
	}
	findUser, _ := u.UserRepository.FindBy(cond)
	if findUser != nil {
		return nil, fmt.Errorf("当前用户名不可用，请检查")
	}
	err := u.UserRepository.Create(user)
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
	cond := repository.StructCondition[model.UserFilter]{
		userFiler,
	}
	user, err := u.UserRepository.FindBy(cond)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u UserService) GetByObject() {
	//TODO implement me
	panic("implement me")
}

func (u UserService) GetPageByFilter(modelFilter model.UserFilter, pagination model.Pagination) (int64, []model.User, error) {
	cond := repository.StructCondition[model.UserFilter]{
		modelFilter,
	}
	total, users, err := u.UserRepository.Page(cond, pagination)
	if err != nil {
		return 0, nil, fmt.Errorf("用户分页查询失败: %w", err)
	}
	return total, users, nil
}
