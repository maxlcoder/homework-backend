package model

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Model struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type ModelCreatedAtUpdatedAt struct {
	CreatedAt time.Time
	UpdatedAt time.Time
}

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
		&Menu{},
		&MenuPermission{},
		&Permission{},
		&Role{},
		&AdminRole{},
	}
}

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(Models()...)
}
