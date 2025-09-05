package repository

import (
	"github.com/maxlcoder/homework-backend/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	Repository[model.User]
}

type UserRepositoryImpl struct {
	*BaseRepository[model.User]
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{
		NewBaseRepository[model.User](db),
	}
}
