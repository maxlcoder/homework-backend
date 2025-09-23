package model

import (
	"time"
)

type Admin struct {
	BaseModel
	Name     string  `gorm:"size:30;not null;default:''"`
	Email    string  `gorm:"size:60;not null;default:''"`
	Age      uint8   `gorm:"not null;default:0"`
	Password string  `gorm:"size:100;not null;default:''"`
	Roles    []*Role `gorm:"many2many:admin_roles;"`
}

type AdminRole struct {
	IDModel
	AdminId uint `gorm:"not null;default:0;uniqueIndex:uq_admin_role"`
	RoleId  uint `gorm:"not null;default:0;uniqueIndex:uq_admin_role"`
}

type AdminFilter struct {
	ID        *uint
	Name      *string `form:"name"`
	Email     *string
	Age       *uint8
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

func (a *Admin) GetId() uint {
	return a.ID
}

func (a *Admin) GetPassword() string {
	return a.Password
}
