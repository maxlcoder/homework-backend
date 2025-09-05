package repository

import (
	"github.com/maxlcoder/homework-backend/model"
	"gorm.io/gorm"
)

type AdminRepository interface {
	Repository[model.Admin]
}

type AdminRepositoryImpl struct {
	*BaseRepository[model.Admin]
}

func NewAdminRepository(db *gorm.DB) AdminRepository {
	return &AdminRepositoryImpl{
		NewBaseRepository[model.Admin](db),
	}
}
