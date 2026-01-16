package repository

import (
	"github.com/maxlcoder/homework-backend/app/modules/core/model"
	"github.com/maxlcoder/homework-backend/repository"
	"gorm.io/gorm"
)

type AdminRepository interface {
	repository.Repository[model.Admin]
}

type AdminRepositoryImpl struct {
	*repository.BaseRepository[model.Admin]
}

func NewAdminRepository(db *gorm.DB) AdminRepository {
	return &AdminRepositoryImpl{
		repository.NewBaseRepository[model.Admin](db),
	}
}
