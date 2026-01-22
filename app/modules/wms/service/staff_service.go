package service

import (
	"errors"

	"github.com/maxlcoder/homework-backend/app/modules/wms/model"
	base_model "github.com/maxlcoder/homework-backend/model"
	"gorm.io/gorm"

	"github.com/maxlcoder/homework-backend/app/modules/wms/admin/request"
	base_request "github.com/maxlcoder/homework-backend/app/request"
	"github.com/maxlcoder/homework-backend/repository"
)

// StaffServiceInterface 仓库人员服务接口
type StaffServiceInterface interface {
	Page(pageRequest base_request.PageRequest) ([]model.Staff, int64, error)
	Create(model *model.Staff) (*model.Staff, error)
	Update(model *model.Staff) (*model.Staff, error)
	Delete(id uint) error
	GetStaff(id uint) (*model.Staff, error)
	UpdateStaff(id uint, request request.StaffUpdateRequest) (*model.Staff, error)
	ListStaffs(filter request.StaffFilterRequest) ([]model.Staff, int64, error)
	UpdateStaffState(id uint, state model.StaffState) (*model.Staff, error)
}

// StaffService 仓库人员服务实现
type StaffService struct {
	db *gorm.DB
}

// NewStaffService 创建仓库人员服务实例
func NewStaffService(db *gorm.DB) StaffServiceInterface {
	return &StaffService{
		db: db,
	}
}

// Page 分页获取仓库人员列表
func (u *StaffService) Page(pageRequest base_request.PageRequest) ([]model.Staff, int64, error) {
	// 创建分页参数
	pagination := base_model.Pagination{
		Page:    pageRequest.Page,
		PerPage: pageRequest.PerPage,
	}

	// 查询数据
	count, staffs, err := repository.NewBaseRepository[model.Staff](u.db).Page(repository.ConditionScope{}, pagination)
	if err != nil {
		return nil, 0, err
	}

	return staffs, count, nil
}

// Create 创建仓库人员
func (u *StaffService) Create(staff *model.Staff) (*model.Staff, error) {
	// 检查姓名是否已存在
	cond := repository.ConditionScope{
		MapCond: map[string]interface{}{
			"name": staff.Name,
		},
	}
	find, err := repository.NewBaseRepository[model.Staff](u.db).FindBy(cond)
	if err == nil && find != nil {
		return nil, errors.New("该姓名的仓库人员已存在")
	}

	err = repository.NewBaseRepository[model.Staff](u.db).Create(staff, nil)
	if err != nil {
		return nil, err
	}

	return staff, nil
}

// Update 更新仓库人员信息
func (u *StaffService) Update(staff *model.Staff) (*model.Staff, error) {
	// 保存更新
	err := repository.NewBaseRepository[model.Staff](u.db).Update(staff, nil)
	if err != nil {
		return nil, err
	}

	return staff, nil
}

// Delete 删除仓库人员
func (u *StaffService) Delete(id uint) error {
	return repository.NewBaseRepository[model.Staff](u.db).DeleteById(id, nil)
}

// GetStaff 根据 ID 获取仓库人员
func (u *StaffService) GetStaff(id uint) (*model.Staff, error) {
	return repository.NewBaseRepository[model.Staff](u.db).FindById(id)
}

// UpdateStaff 更新仓库人员信息
func (u *StaffService) UpdateStaff(id uint, request request.StaffUpdateRequest) (*model.Staff, error) {
	// 获取现有仓库人员
	staff, err := repository.NewBaseRepository[model.Staff](u.db).FindById(id)
	if err != nil {
		return nil, err
	}

	// 如果姓名变更，检查新姓名是否已存在
	if request.Name != "" && request.Name != staff.Name {
		existingStaff, err := repository.NewBaseRepository[model.Staff](u.db).FindByName(request.Name)
		if err == nil && existingStaff != nil && existingStaff.ID != id {
			return nil, errors.New("该姓名的仓库人员已存在")
		}
		staff.Name = request.Name
	}

	// 更新状态
	if request.State != 0 {
		staff.State = model.StaffState(request.State)
	}

	// 保存更新
	err = repository.NewBaseRepository[model.Staff](u.db).Update(staff, nil)
	if err != nil {
		return nil, err
	}

	return staff, nil
}

// ListStaffs 获取仓库人员列表
func (u *StaffService) ListStaffs(filter request.StaffFilterRequest) ([]model.Staff, int64, error) {
	// 创建查询条件
	cond := repository.ConditionScope{
		MapCond: make(map[string]interface{}),
		Scopes:  []func(*gorm.DB) *gorm.DB{},
	}
	if filter.Name != "" {
		cond.Scopes = append(cond.Scopes, func(db *gorm.DB) *gorm.DB {
			return db.Where("name LIKE ?", "%"+filter.Name+"%")
		})
	}
	if filter.State != 0 {
		cond.MapCond["state"] = filter.State
	}

	// 创建分页参数
	pagination := base_model.Pagination{
		Page:    filter.Page,
		PerPage: filter.PerPage,
	}

	// 查询数据
	count, staffs, err := repository.NewBaseRepository[model.Staff](u.db).Page(cond, pagination)
	if err != nil {
		return nil, 0, err
	}

	return staffs, count, nil
}

// UpdateStaffState 更新仓库人员状态
func (u *StaffService) UpdateStaffState(id uint, state model.StaffState) (*model.Staff, error) {
	// 验证状态值
	if !state.IsValid() {
		return nil, errors.New("无效的状态值")
	}

	// 获取现有仓库人员
	staff, err := repository.NewBaseRepository[model.Staff](u.db).FindById(id)
	if err != nil {
		return nil, err
	}

	// 更新状态
	staff.State = state

	// 保存更新
	err = repository.NewBaseRepository[model.Staff](u.db).Update(staff, nil)
	if err != nil {
		return nil, err
	}

	return staff, nil
}
