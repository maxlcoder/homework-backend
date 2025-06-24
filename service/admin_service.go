package service

import (
	"fmt"
	"log"

	"github.com/maxlcoder/homework-backend/model"
	"github.com/maxlcoder/homework-backend/repository"
)

type AdminServiceInterface interface {
	Create(admin *model.Admin) (*model.Admin, error)
	GetById(id uint) (*model.Admin, error)
	GetPageByFilter(modelFilter model.UserFilter, paginationQuery model.PaginationQuery) (int64, []model.Admin, error)
}

type AdminService struct {
	AdminRepository repository.AdminRepository
}

func (u *AdminService) Create(admin *model.Admin) (*model.Admin, error) {
	// 判断是否存在已经适用的名称
	userFiler := model.UserFilter{
		Name: &admin.Name,
	}
	log.Default().Println(admin.Name)
	findUser, _ := u.AdminRepository.FindBy(userFiler)
	if findUser != nil {
		return nil, fmt.Errorf("当前用户名不可用，请检查")
	}
	err := u.AdminRepository.Create(admin)
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
	user, err := u.AdminRepository.FindBy(userFiler)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u AdminService) GetByObject() {
	//TODO implement me
	panic("implement me")
}

func (u AdminService) GetPageByFilter(modelFilter model.UserFilter, paginationQuery model.PaginationQuery) (int64, []model.Admin, error) {
	total, users, err := u.AdminRepository.Paginate(modelFilter, paginationQuery)
	if err != nil {
		return 0, nil, fmt.Errorf("用户分页查询失败: %w", err)
	}
	return total, users, nil
}
