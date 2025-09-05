package repository

import (
	"github.com/maxlcoder/homework-backend/model"
	"gorm.io/gorm"
)

type MenuRepository interface {
	Repository[model.Menu]
}

type MenuRepositoryImpl struct {
	*BaseRepository[model.Menu]
}

func NewMenuRepository(db *gorm.DB) MenuRepository {
	return &MenuRepositoryImpl{
		NewBaseRepository[model.Menu](db),
	}
}
