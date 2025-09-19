package repository

import (
	"github.com/maxlcoder/homework-backend/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type RoleMenuRepository interface {
	Repository[model.RoleMenu] // 继承通用方法
	UpsertCreate(roleMenu *model.RoleMenu) error
	UpsertCreateBatch(roleMenus []model.RoleMenu, tx *gorm.DB) error
}

type RoleMenuRepositoryImpl struct {
	*BaseRepository[model.RoleMenu] // 内嵌字段，继承方法提升
}

func NewRoleMenuRepository(db *gorm.DB) RoleMenuRepository {
	return &RoleMenuRepositoryImpl{
		BaseRepository: NewBaseRepository[model.RoleMenu](db),
	}
}

func (r *RoleMenuRepositoryImpl) UpsertCreate(roleMenu *model.RoleMenu) error {
	r.db.Clauses(clause.OnConflict{
		UpdateAll: false,
	}).Create(&model.RoleMenu{
		RoleID: roleMenu.RoleID,
		MenuID: roleMenu.MenuID,
	})
	return nil
}

func (r *RoleMenuRepositoryImpl) UpsertCreateBatch(roleMenus []model.RoleMenu, tx *gorm.DB) error {
	r.getDB(tx).Clauses(clause.OnConflict{
		DoNothing: true,
	}).Create(&roleMenus)
	return nil
}
