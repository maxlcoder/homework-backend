package repository

import (
	"github.com/maxlcoder/homework-backend/model"
	"gorm.io/gorm"
)

type BinRepository interface {
	Repository[model.Bin]
}

type BinRepositoryImpl struct {
	*BaseRepository[model.Bin]
}

func NewBinRepository(db *gorm.DB) BinRepository {
	return &BinRepositoryImpl{
		NewBaseRepository[model.Bin](db),
	}
}
