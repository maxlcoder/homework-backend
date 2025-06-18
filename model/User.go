package model

import (
	"time"

	"golang.org/x/crypto/bcrypt"
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

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
