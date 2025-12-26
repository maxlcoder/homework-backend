package repository

import (
	"github.com/maxlcoder/homework-backend/app/modules/wms/model"
	"github.com/maxlcoder/homework-backend/repository"
	"gorm.io/gorm"
)

type PickingTaskBasketProductRepository interface {
	repository.Repository[model.PickingTaskBasketProduct]
}

type PickingTaskBasketProductRepositoryImpl struct {
	*repository.BaseRepository[model.PickingTaskBasketProduct]
}

func NewPickingTaskBasketProductRepository(db *gorm.DB) PickingTaskBasketProductRepository {
	return &PickingTaskBasketProductRepositoryImpl{
		repository.NewBaseRepository[model.PickingTaskBasketProduct](db),
	}
}
