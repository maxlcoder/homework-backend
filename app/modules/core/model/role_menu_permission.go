package model

import (
	"time"

	base_model "github.com/maxlcoder/homework-backend/model"
)

type Role struct {
	base_model.BaseModel
	Name     string `gorm:"size:60;not null;default:''"`
	TenantId uint   `gorm:"not null;default:0;comment:租户 ID"`
}

type RoleFilter struct {
	ID        *uint
	Name      *string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

type RoleMenu struct {
	base_model.BaseModel
	RoleID uint `gorm:"not null;default:0"`
	MenuID uint `gorm:"not null;default:0"`
}

type RolePermission struct {
	base_model.BaseModel
	RoleID       uint `gorm:"not null;default:0"`
	PermissionID uint `gorm:"not null;default:0"`
}
