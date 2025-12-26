package repository

import (
	"github.com/maxlcoder/homework-backend/app/modules/wms/model"
	base_repository "github.com/maxlcoder/homework-backend/repository"
	"gorm.io/gorm"
)

type PickingBasketRepository interface {
	base_repository.Repository[model.PickingBasket]
}

type PickingBasketRepositoryImpl struct {
	*base_repository.BaseRepository[model.PickingBasket]
}

func NewPickingBasketRepository(db *gorm.DB) PickingBasketRepository {
	return &PickingBasketRepositoryImpl{
		base_repository.NewBaseRepository[model.PickingBasket](db),
	}
}
