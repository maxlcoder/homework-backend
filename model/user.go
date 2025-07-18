package model

import (
	"time"
)

type User struct {
	ID        uint
	Name      string    `gorm:"size:30,unique;not null;default:''"`
	Email     string    `gorm:"size:60,unique;not null;default:''"`
	Age       uint8     `gorm:"not null;default:0"`
	Password  string    `gorm:"size:100;not null;default:''"`
	CreatedAt time.Time `gorm:"not null;autoCreateTime"`
	UpdatedAt time.Time `gorm:"not null;autoUpdateTime"`
}

type UserFilter struct {
	ID        *uint
	Name      *string
	Email     *string
	Age       *uint8
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

func (u *User) GetId() uint {
	return u.ID
}

func (u *User) GetPassword() string {
	return u.Password
}
