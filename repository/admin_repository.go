package repository

import (
	"errors"

	"github.com/maxlcoder/homework-backend/model"
	"gorm.io/gorm"
)

type AdminRepository interface {
	Create(admin *model.Admin) error
	FindBy(userFilter model.UserFilter) (*model.Admin, error)
	Paginate(userFilter model.UserFilter, paginationQuery model.PaginationQuery) (int64, []model.Admin, error)
}

type adminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) AdminRepository {
	return &adminRepository{db}
}

func (r *adminRepository) Create(user *model.Admin) error {
	return r.db.Create(user).Error
}

func (r *adminRepository) FindBy(userFilter model.UserFilter) (*model.Admin, error) {
	query := r.db.Model(&model.Admin{})
	if userFilter.ID != nil {
		query.Where("id = ?", userFilter.ID)
	}
	if userFilter.Name != nil {
		query.Where("name = ?", userFilter.Name)
	}
	if userFilter.Email != nil {
		query.Where("email = ?", userFilter.Email)
	}
	user := model.Admin{}
	result := query.First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}
	return &user, nil
}

func (r *adminRepository) FindByEmail(email string) (*model.Admin, error) {
	user := &model.Admin{}
	user.Email = email
	return user, r.db.First(user, "email = ?", email).Error
}

func (r *adminRepository) Paginate(userFilter model.UserFilter, paginationQuery model.PaginationQuery) (int64, []model.Admin, error) {
	var users []model.Admin
	var total int64
	query := r.db.Model(&model.Admin{})
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
