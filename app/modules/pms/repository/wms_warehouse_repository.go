package repository

import (
	"github.com/maxlcoder/homework-backend/app/modules/wms/model"
	base_repository "github.com/maxlcoder/homework-backend/repository"
	"gorm.io/gorm"
)

type WarehouseRepository interface {
	base_repository.Repository[model.Warehouse]
}

type WarehouseRepositoryImpl struct {
	*base_repository.BaseRepository[model.Warehouse]
}

func NewWarehouseRepository(db *gorm.DB) WarehouseRepository {
	return &WarehouseRepositoryImpl{
		base_repository.NewBaseRepository[model.Warehouse](db),
	}
}
