package model

import "time"

type Role struct {
	BaseModel
	Name string `gorm:"size:60;not null;default:''"`
}

type RoleFilter struct {
	ID        *uint
	Name      *string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

type RoleMenu struct {
	BaseModel
	RoleID uint `gorm:"not null;default:0"`
	MenuID uint `gorm:"not null;default:0"`
}

type RolePermission struct {
	BaseModel
	RoleID       uint `gorm:"not null;default:0"`
	PermissionID uint `gorm:"not null;default:0"`
}
