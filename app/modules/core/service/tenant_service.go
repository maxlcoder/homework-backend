package service

import (
	"fmt"

	"github.com/maxlcoder/homework-backend/app/modules/core/admin/request"
	model2 "github.com/maxlcoder/homework-backend/app/modules/core/model"
	repository2 "github.com/maxlcoder/homework-backend/app/modules/core/repository"
	"github.com/maxlcoder/homework-backend/model"
	"github.com/maxlcoder/homework-backend/repository"
	"gorm.io/gorm"
)

type TenantServiceInterface interface {
	Page(pageRequest request.TenantPageRequest) ([]model2.Tenant, int64, error)
	Create(model *model2.Tenant) (*model2.Tenant, error)
	Update(model *model2.Tenant) (*model2.Tenant, error)
	Delete(id uint) error
	FindById(id uint) (*model2.Tenant, error)
}

type TenantService struct {
	db               *gorm.DB
	TenantRepository repository2.TenantRepository
}

func NewTenantService(db *gorm.DB, tenantRepository repository2.TenantRepository) TenantServiceInterface {
	return &TenantService{
		db:               db,
		TenantRepository: tenantRepository,
	}
}

func (u *TenantService) Page(pageRequest request.TenantPageRequest) ([]model2.Tenant, int64, error) {
	cond := repository.ConditionScope{}

	if pageRequest.Name != nil && len(*pageRequest.Name) > 0 {
		cond.StructCond = request.TenantPageRequest{
			Name: pageRequest.Name,
		}
	}

	// 创建分页参数
	pagination := model.Pagination{
		Page:    pageRequest.Page,
		PerPage: pageRequest.PerPage,
	}

	// 查询数据
	count, tenants, err := u.TenantRepository.Page(cond, pagination)
	if err != nil {
		return nil, 0, fmt.Errorf("获取租户列表失败: %w", err)
	}

	return tenants, count, nil
}

func (u *TenantService) Create(tenant *model2.Tenant) (*model2.Tenant, error) {
	// 判断是否存在已经适用的名称
	filer := model2.Tenant{
		Name: tenant.Name,
	}
	cond := repository.ConditionScope{
		StructCond: filer,
	}
	find, _ := u.TenantRepository.FindBy(cond)
	if find != nil {
		return nil, fmt.Errorf("当前租户名称不可用，请检查")
	}
	err := u.TenantRepository.Create(tenant, nil)
	if err != nil {
		return nil, fmt.Errorf("租户创建失败: %w", err)
	}
	return tenant, nil
}

func (u *TenantService) Update(tenant *model2.Tenant) (*model2.Tenant, error) {
	// 判断是否存在已经适用的名称（排除自身）
	filer := model2.Tenant{
		Name: tenant.Name,
	}
	cond := repository.ConditionScope{
		StructCond: filer,
	}
	find, _ := u.TenantRepository.FindBy(cond)
	if find != nil && find.ID != tenant.ID {
		return nil, fmt.Errorf("当前租户名称不可用，请检查")
	}

	err := u.TenantRepository.Update(tenant, u.db)
	if err != nil {
		return nil, fmt.Errorf("租户更新失败: %w", err)
	}
	return tenant, nil
}

func (u *TenantService) Delete(id uint) error {
	err := u.TenantRepository.DeleteById(id, nil)
	if err != nil {
		return fmt.Errorf("租户删除失败: %w", err)
	}
	return nil
}

func (u *TenantService) FindById(id uint) (*model2.Tenant, error) {
	tenant, err := u.TenantRepository.FindById(id)
	if err != nil {
		return nil, fmt.Errorf("租户查询失败: %w", err)
	}
	if tenant == nil {
		return nil, fmt.Errorf("租户不存在")
	}
	return tenant, nil
}
