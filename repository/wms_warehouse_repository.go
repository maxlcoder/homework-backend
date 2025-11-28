package repository

import (
	"github.com/maxlcoder/homework-backend/model"
	"gorm.io/gorm"
)

type WarehouseRepository interface {
	Repository[model.Warehouse]
}

type WarehouseRepositoryImpl struct {
	*BaseRepository[model.Warehouse]
}

func NewWarehouseRepository(db *gorm.DB) WarehouseRepository {
	return &WarehouseRepositoryImpl{
		NewBaseRepository[model.Warehouse](db),
	}
}
