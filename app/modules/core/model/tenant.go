package model

import (
	"time"

	model2 "github.com/maxlcoder/homework-backend/model"
)

type Tenant struct {
	model2.BaseModel
	Name string `gorm:"size:60;not null;default:'';unique"`
}

type TenantFilter struct {
	ID        *uint
	Name      *string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

type TenantUser struct {
	model2.BaseModel
	TenantId uint `gorm:"not null;default:0;comment:租户 ID"`
	UserId   uint `gorm:"not null;default:0;comment:用户 ID"`
}

type TenantAdmin struct {
	model2.BaseModel
	TenantId uint `gorm:"not null;default:0;comment:租户 ID"`
	AdminId  uint `gorm:"not null;default:0;comment:管理员 ID"`
}
