package model

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uint
	CreatedAt time.Time `gorm:"not null;autoCreateTime;comment:创建时间"`
	UpdatedAt time.Time `gorm:"not null;autoUpdateTime;comment:更新时间"`
}

type IDModel struct {
	ID uint
}

type BaseTenantModel struct {
	BaseModel
	TenantID uint `gorm:"not null;default:0;comment:租户 ID"`
}

type ModelCreatedAtUpdatedAt struct {
	CreatedAt time.Time `gorm:"not null;autoCreateTime;comment:创建时间"`
	UpdatedAt time.Time `gorm:"not null;autoUpdateTime;comment:创建时间"`
}

type Pagination struct {
	Page    int `form:"page" default:"1" binding:"omitempty,min=1" label:"页码"`
	PerPage int `form:"per_page" default:"10" binding:"omitempty,min=1,max=100" label:"每页大小"`
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
		&RoleMenu{},
		&RolePermission{},
	}
}

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(Models()...)
}
