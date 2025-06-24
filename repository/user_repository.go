package repository

import (
	"errors"

	"github.com/maxlcoder/homework-backend/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *model.User) error
	FindBy(userFilter model.UserFilter) (*model.User, error)
	Paginate(userFilter model.UserFilter, paginationQuery model.PaginationQuery) (int64, []model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) Create(user *model.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) FindBy(userFilter model.UserFilter) (*model.User, error) {
	query := r.db.Model(&model.User{})
	if userFilter.ID != nil {
		query.Where("id = ?", userFilter.ID)
	}
	if userFilter.Name != nil {
		query.Where("name = ?", userFilter.Name)
	}
	if userFilter.Email != nil {
		query.Where("email = ?", userFilter.Email)
	}
	user := model.User{}
	result := query.First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}
	return &user, nil
}

func (r *userRepository) FindByEmail(email string) (*model.User, error) {
	user := &model.User{}
	user.Email = email
	return user, r.db.First(user, "email = ?", email).Error
}

func (r *userRepository) Paginate(userFilter model.UserFilter, paginationQuery model.PaginationQuery) (int64, []model.User, error) {
	var users []model.User
	var total int64
	query := r.db.Model(&model.User{})
	if userFilter.ID != nil {
		query.Where("id = ?", userFilter.ID)
	}
	if userFilter.Name != nil {
		query.Where("name = ?", userFilter.Name)
	}
	if userFilter.Email != nil {
		query.Where("email = ?", userFilter.Email)
	}
	total, err := Paginate(query, paginationQuery, &users)
	return total, users, err
}
