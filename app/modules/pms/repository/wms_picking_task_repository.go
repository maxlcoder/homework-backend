package repository

import (
	"github.com/maxlcoder/homework-backend/app/modules/wms/model"
	base_repository "github.com/maxlcoder/homework-backend/repository"
	"gorm.io/gorm"
)

type PickingTaskRepository interface {
	base_repository.Repository[model.PickingTask]
}

type PickingTaskRepositoryImpl struct {
	*base_repository.BaseRepository[model.PickingTask]
}

func NewPickingTaskRepository(db *gorm.DB) PickingTaskRepository {
	return &PickingTaskRepositoryImpl{
		base_repository.NewBaseRepository[model.PickingTask](db),
	}
}
