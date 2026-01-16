package model

import (
	"time"

	base_model "github.com/maxlcoder/homework-backend/model"
)

type Admin struct {
	base_model.BaseModel
	Name     string  `gorm:"size:30;not null;default:''"`
	Email    string  `gorm:"size:60;not null;default:''"`
	Age      uint8   `gorm:"not null;default:0"`
	Password string  `gorm:"size:100;not null;default:''"`
	RoleId   uint    `gorm:"comment:当前角色 ID"`
	Roles    []*Role `gorm:"many2many:admin_roles;"`
}

type AdminRole struct {
	base_model.IDModel
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
