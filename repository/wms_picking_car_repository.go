package repository

import (
	"github.com/maxlcoder/homework-backend/model"
	"gorm.io/gorm"
)

type PickingCarRepository interface {
	Repository[model.PickingCar]
}

type PickingCarRepositoryImpl struct {
	*BaseRepository[model.PickingCar]
}

func NewPickingCarRepository(db *gorm.DB) PickingCarRepository {
	return &PickingCarRepositoryImpl{
		NewBaseRepository[model.PickingCar](db),
	}
}
