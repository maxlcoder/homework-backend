package repository

import (
	"github.com/maxlcoder/homework-backend/app/modules/wms/model"
	base_repository "github.com/maxlcoder/homework-backend/repository"
	"gorm.io/gorm"
)

type PickingCarRepository interface {
	base_repository.Repository[model.PickingCar]
}

type PickingCarRepositoryImpl struct {
	*base_repository.BaseRepository[model.PickingCar]
}

func NewPickingCarRepository(db *gorm.DB) PickingCarRepository {
	return &PickingCarRepositoryImpl{
		base_repository.NewBaseRepository[model.PickingCar](db),
	}
}
