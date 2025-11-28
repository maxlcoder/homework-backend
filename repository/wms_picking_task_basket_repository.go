package repository

import (
	"github.com/maxlcoder/homework-backend/model"
	"gorm.io/gorm"
)

type PickingTaskBasketRepository interface {
	Repository[model.PickingTaskBasket]
}

type PickingTaskBasketRepositoryImpl struct {
	*BaseRepository[model.PickingTaskBasket]
}

func NewPickingTaskBasketRepository(db *gorm.DB) PickingTaskBasketRepository {
	return &PickingTaskBasketRepositoryImpl{
		NewBaseRepository[model.PickingTaskBasket](db),
	}
}
