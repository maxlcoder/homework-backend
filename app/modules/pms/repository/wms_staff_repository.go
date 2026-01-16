package repository

import (
	"github.com/maxlcoder/homework-backend/app/modules/wms/model"
	repository "github.com/maxlcoder/homework-backend/repository"
	"gorm.io/gorm"
)

// StaffRepository 员工仓库接口
type StaffRepository interface {
	repository.Repository[model.Staff]
	FindByName(name string) (*model.Staff, error)
	FindByState(state string) ([]model.Staff, error)
}

// StaffRepositoryImpl 员工仓库实现
type StaffRepositoryImpl struct {
	*repository.BaseRepository[model.Staff]
}

// NewStaffRepository 创建员工仓库实例
func NewStaffRepository(db *gorm.DB) StaffRepository {
	return &StaffRepositoryImpl{
		repository.NewBaseRepository[model.Staff](db),
	}
}

// FindByName 根据姓名查询员工
func (r *StaffRepositoryImpl) FindByName(name string) (*model.Staff, error) {
	var staff model.Staff
	result := r.DB.Where("name = ?", name).First(&staff)
	if result.Error != nil {
		return nil, result.Error
	}
	return &staff, nil
}

// FindByState 根据状态查询员工
func (r *StaffRepositoryImpl) FindByState(state string) ([]model.Staff, error) {
	var staffs []model.Staff
	result := r.DB.Where("state = ?", state).Find(&staffs)
	if result.Error != nil {
		return nil, result.Error
	}
	return staffs, nil
}
