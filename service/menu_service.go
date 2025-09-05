package service

import (
	"fmt"

	"github.com/maxlcoder/homework-backend/model"
	"github.com/maxlcoder/homework-backend/repository"
)

type MenuServiceInterface interface {
	Create(menu *model.Menu) (*model.Menu, error)
	GetById(id uint) (*model.Menu, error)
	GetPageByFilter(modelFilter model.MenuFilter, pagination model.Pagination) (int64, []model.Menu, error)
}

type MenuService struct {
	MenuRepository repository.MenuRepository
}

func (u *MenuService) Create(menu *model.Menu) (*model.Menu, error) {
	// 判断是否存在已经适用的名称
	filter := model.MenuFilter{
		Name: &menu.Name,
	}
	cond := repository.StructCondition[model.MenuFilter]{
		filter,
	}
	findUser, _ := u.MenuRepository.FindBy(cond)
	if findUser != nil {
		return nil, fmt.Errorf("当前用户名不可用，请检查")
	}
	err := u.MenuRepository.Create(menu)
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
	cond := repository.StructCondition[model.MenuFilter]{
		filter,
	}
	user, err := u.MenuRepository.FindBy(cond)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *MenuService) GetByObject() {
	//TODO implement me
	panic("implement me")
}

func (u *MenuService) GetPageByFilter(modelFilter model.MenuFilter, pagination model.Pagination) (int64, []model.Menu, error) {

	cond := repository.StructCondition[model.MenuFilter]{
		modelFilter,
	}

	total, menus, err := u.MenuRepository.Page(cond, pagination)
	if err != nil {
		return 0, nil, fmt.Errorf("用户分页查询失败: %w", err)
	}
	return total, menus, nil
}
