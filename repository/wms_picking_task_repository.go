package repository

import (
	"github.com/maxlcoder/homework-backend/model"
	"gorm.io/gorm"
)

type PickingTaskRepository interface {
	Repository[model.PickingTask]
}

type PickingTaskRepositoryImpl struct {
	*BaseRepository[model.PickingTask]
}

func NewPickingTaskRepository(db *gorm.DB) PickingTaskRepository {
	return &PickingTaskRepositoryImpl{
		NewBaseRepository[model.PickingTask](db),
	}
}
