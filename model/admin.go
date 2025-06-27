package model

import (
	"time"
)

type Admin struct {
	ID       uint
	Name     string `gorm:"size:30;not null;default:''"`
	Email    string `gorm:"size:60;not null;default:''"`
	Age      uint8  `gorm:"not null;default:0"`
	Password string `gorm:"size:100;not null;default:''"`
	ModelCreatedAtUpdatedAt
}

type AdminRole struct {
	ID      uint
	AdminId uint
	RoleId  uint
	ModelCreatedAtUpdatedAt
}

type AdminFilter struct {
	ID        *uint
	Name      *string
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
