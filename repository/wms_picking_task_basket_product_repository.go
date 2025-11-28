package repository

import (
	"github.com/maxlcoder/homework-backend/model"
	"gorm.io/gorm"
)

type PickingTaskBasketProductRepository interface {
	Repository[model.PickingTaskBasketProduct]
}

type PickingTaskBasketProductRepositoryImpl struct {
	*BaseRepository[model.PickingTaskBasketProduct]
}

func NewPickingTaskBasketProductRepository(db *gorm.DB) PickingTaskBasketProductRepository {
	return &PickingTaskBasketProductRepositoryImpl{
		NewBaseRepository[model.PickingTaskBasketProduct](db),
	}
}
