package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type PaginationQuery struct {
	Page    int `form:"page" binding:"min=1"`
	PerPage int `form:"per_page" binding:"min=1,max=100"`
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

type Authenticatable interface {
	GetId() uint
	GetPassword() string
}

func Models() []interface{} {
	return []interface{}{
		&User{},
		&Admin{},
		&Casbin{},
	}
}

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(Models()...)
}
