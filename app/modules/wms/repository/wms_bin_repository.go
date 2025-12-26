package repository

import (
	"github.com/maxlcoder/homework-backend/app/modules/wms/model"
	"github.com/maxlcoder/homework-backend/repository"
	"gorm.io/gorm"
)

type BinRepository interface {
	repository.Repository[model.Bin]
}

type BinRepositoryImpl struct {
	*repository.BaseRepository[model.Bin]
}

func NewBinRepository(db *gorm.DB) BinRepository {
	return &BinRepositoryImpl{
		repository.NewBaseRepository[model.Bin](db),
	}
}
