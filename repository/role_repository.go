package repository

import (
	"errors"

	"github.com/maxlcoder/homework-backend/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type RoleRepository interface {
	Repository[model.Role] // 继承通用方法
	// 扩展业务方法
	FindByName(name string) (*model.Role, error)
	UpsertCreateRolePermissionBatch(rolePermissions []model.RolePermission, tx *gorm.DB) error
	DeleteRolePermissionsByRoleId(roleId uint, tx *gorm.DB) error
}

type RoleRepositoryImpl struct {
	*BaseRepository[model.Role] // 内嵌字段，继承方法提升
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &RoleRepositoryImpl{
		BaseRepository: NewBaseRepository[model.Role](db),
	}
}

// 扩展方法
func (r *RoleRepositoryImpl) FindByName(name string) (*model.Role, error) {
	if name == "" {
		return nil, errors.New("name is empty")
	}
	var role model.Role
	if err := r.db.Where("name = ?", name).First(&role).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *RoleRepositoryImpl) UpsertCreateRolePermissionBatch(rolePermissions []model.RolePermission, tx *gorm.DB) error {
	r.getDB(tx).Clauses(clause.OnConflict{
		DoNothing: true,
	}).Create(&rolePermissions)
	return nil
}

func (r *RoleRepositoryImpl) DeleteRolePermissionsByRoleId(roleId uint, tx *gorm.DB) error {
	r.getDB(tx).Where("role_id = ?", roleId).Delete(&model.RolePermission{})
	return nil
}
