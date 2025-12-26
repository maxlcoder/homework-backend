package repository

import (
	"fmt"

	"github.com/maxlcoder/homework-backend/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type AdminRoleRepository interface {
	Repository[model.AdminRole] // 继承通用方法
	UpsertCreate(adminRole *model.AdminRole) error
	UpsertCreateBatch(adminRoles []model.AdminRole, tx *gorm.DB) error
}

type AdminRoleRepositoryImpl struct {
	*BaseRepository[model.AdminRole] // 内嵌字段，继承方法提升
}

func NewAdminRoleRepository(db *gorm.DB) AdminRoleRepository {
	return &AdminRoleRepositoryImpl{
		BaseRepository: NewBaseRepository[model.AdminRole](db),
	}
}

func (r *AdminRoleRepositoryImpl) UpsertCreate(adminRole *model.AdminRole) error {
	r.DB.Clauses(clause.OnConflict{
		UpdateAll: false,
	}).Create(&model.AdminRole{
		AdminId: adminRole.AdminId,
		RoleId:  adminRole.RoleId,
	})
	return nil
}

func (r *AdminRoleRepositoryImpl) UpsertCreateBatch(adminRoles []model.AdminRole, tx *gorm.DB) error {
	fmt.Print(adminRoles)
	r.getDB(tx).Clauses(clause.OnConflict{
		UpdateAll: false,
	}).Create(&adminRoles)
	return nil
}
