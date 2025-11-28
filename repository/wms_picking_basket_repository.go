package repository

import (
	"github.com/maxlcoder/homework-backend/model"
	"gorm.io/gorm"
)

type PickingBasketRepository interface {
	Repository[model.PickingBasket]
}

type PickingBasketRepositoryImpl struct {
	*BaseRepository[model.PickingBasket]
}

func NewPickingBasketRepository(db *gorm.DB) PickingBasketRepository {
	return &PickingBasketRepositoryImpl{
		NewBaseRepository[model.PickingBasket](db),
	}
}
