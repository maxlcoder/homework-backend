package service

import (
	"fmt"

	"github.com/maxlcoder/homework-backend/model"
	"github.com/maxlcoder/homework-backend/repository"
)

type UserServiceInterface interface {
	Create(*model.User) (*model.User, error)
	GetById(id uint) (*model.User, error)
	GetPageByFilter(modelFilter model.UserFilter, paginationQuery model.PaginationQuery) (int64, []model.User, error)
}

type UserService struct {
	UserRepository repository.UserRepository
}

func (u *UserService) Create(user *model.User) (*model.User, error) {
	// 判断是否存在已经适用的名称
	userFiler := model.UserFilter{
		Name: &user.Name,
	}
	findUser, _ := u.UserRepository.FindBy(userFiler)
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
	user, err := u.UserRepository.FindBy(userFiler)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u UserService) GetByObject() {
	//TODO implement me
	panic("implement me")
}

func (u UserService) GetPageByFilter(modelFilter model.UserFilter, paginationQuery model.PaginationQuery) (int64, []model.User, error) {
	total, users, err := u.UserRepository.Paginate(modelFilter, paginationQuery)
	if err != nil {
		return 0, nil, fmt.Errorf("用户分页查询失败: %w", err)
	}
	return total, users, nil
}
