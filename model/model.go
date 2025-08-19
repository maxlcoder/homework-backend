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

type BaseTenantModel struct {
	BaseModel
	TenantID uint `gorm:"not null;default:0;comment:租户 ID"`
}

type ModelCreatedAtUpdatedAt struct {
	CreatedAt time.Time `gorm:"not null;autoCreateTime;comment:创建时间"`
	UpdatedAt time.Time `gorm:"not null;autoUpdateTime;comment:创建时间"`
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
		&RoleMenu{},
		&RolePermission{},
	}
}

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(Models()...)
}
