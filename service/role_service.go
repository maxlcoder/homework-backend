package service

import (
	"fmt"

	"github.com/maxlcoder/homework-backend/model"
	"github.com/maxlcoder/homework-backend/repository"
)

type RoleServiceInterface interface {
	Create(role *model.Role) (*model.Role, error)
	GetById(id uint) (*model.Role, error)
	GetPageByFilter(modelFilter model.RoleFilter, paginationQuery model.PaginationQuery) (int64, []model.Role, error)
}

type RoleService struct {
	RoleRepository repository.RoleRepository
}

func (u *RoleService) Create(role *model.Role) (*model.Role, error) {
	// 判断是否存在已经适用的名称
	filter := model.RoleFilter{
		Name: &role.Name,
	}
	findUser, _ := u.RoleRepository.FindBy(filter)
	if findUser != nil {
		return nil, fmt.Errorf("当前用户名不可用，请检查")
	}
	err := u.RoleRepository.Create(role)
	if err != nil {
		return nil, fmt.Errorf("用户创建失败: %w", err)
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
	user, err := u.RoleRepository.FindBy(filter)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *RoleService) GetByObject() {
	//TODO implement me
	panic("implement me")
}

func (u *RoleService) GetPageByFilter(modelFilter model.RoleFilter, paginationQuery model.PaginationQuery) (int64, []model.Role, error) {
	total, users, err := u.RoleRepository.Paginate(modelFilter, paginationQuery)
	if err != nil {
		return 0, nil, fmt.Errorf("用户分页查询失败: %w", err)
	}
	return total, users, nil
}
