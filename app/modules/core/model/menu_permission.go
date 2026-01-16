package model

import (
	"time"

	base_model "github.com/maxlcoder/homework-backend/model"
)

type Menu struct {
	base_model.BaseModel
	Name        string        `gorm:"size:60;not null;default:''"`
	ParentID    uint          `gorm:"not null;default:0"`
	Sort        int           `gorm:"not null;default:0"`
	IsDisabled  bool          `gorm:"default:0"`
	Number      string        `gorm:"not null;default:'';uniqueIndex"`
	Children    []*Menu       `gorm:"-"`
	Permissions []*Permission `gorm:"-"`
}

type MenuFilter struct {
	ID        *uint
	Name      *string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

type MenuPermission struct {
	MenuID       uint `gorm:"not null;default:0;uniqueIndex:uq_menu_permission"`
	PermissionID uint `gorm:"not null;default:0;uniqueIndex:uq_menu_permission"`
}
