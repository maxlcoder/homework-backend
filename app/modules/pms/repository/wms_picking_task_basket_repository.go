package repository

import (
	"github.com/maxlcoder/homework-backend/app/modules/wms/model"
	base_repository "github.com/maxlcoder/homework-backend/repository"
	"gorm.io/gorm"
)

type PickingTaskBasketRepository interface {
	base_repository.Repository[model.PickingTaskBasket]
}

type PickingTaskBasketRepositoryImpl struct {
	*base_repository.BaseRepository[model.PickingTaskBasket]
}

func NewPickingTaskBasketRepository(db *gorm.DB) PickingTaskBasketRepository {
	return &PickingTaskBasketRepositoryImpl{
		base_repository.NewBaseRepository[model.PickingTaskBasket](db),
	}
}
