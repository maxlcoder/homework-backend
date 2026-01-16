package repository

import (
	"errors"

	"github.com/maxlcoder/homework-backend/app/modules/core/model"
	"github.com/maxlcoder/homework-backend/repository"
	"gorm.io/gorm"
)

type TenantRepository interface {
	repository.Repository[model.Tenant] // 继承通用方法
	// 扩展业务方法
	FindByName(name string) (*model.Tenant, error)
}

type TenantRepositoryImpl struct {
	*repository.BaseRepository[model.Tenant] // 内嵌字段，继承方法提升
}

func NewTenantRepository(db *gorm.DB) TenantRepository {
	return &TenantRepositoryImpl{
		BaseRepository: repository.NewBaseRepository[model.Tenant](db),
	}
}

// 扩展方法
func (r *TenantRepositoryImpl) FindByName(name string) (*model.Tenant, error) {
	if name == "" {
		return nil, errors.New("name is empty")
	}
	var tenant model.Tenant
	if err := r.DB.Where("name = ?", name).First(&tenant).Error; err != nil {
		return nil, err
	}
	return &tenant, nil
}
