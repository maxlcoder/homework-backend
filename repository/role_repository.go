package repository

import (
	"errors"

	"github.com/maxlcoder/homework-backend/model"
	"gorm.io/gorm"
)

type RoleRepository interface {
	Create(role *model.Role) error
	FindBy(roleFilter model.RoleFilter) (*model.Role, error)
	Paginate(roleFilter model.RoleFilter, paginationQuery model.PaginationQuery) (int64, []model.Role, error)
}

type roleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &roleRepository{db}
}

func (r *roleRepository) Create(user *model.Role) error {
	return r.db.Create(user).Error
}

func (r *roleRepository) FindBy(roleFilter model.RoleFilter) (*model.Role, error) {
	query := r.db.Model(&model.Role{})
	if roleFilter.ID != nil {
		query.Where("id = ?", roleFilter.ID)
	}
	if roleFilter.Name != nil {
		query.Where("name = ?", roleFilter.Name)
	}
	role := model.Role{}
	result := query.First(&role)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}
	return &role, nil
}
func (r *roleRepository) Paginate(roleFilter model.RoleFilter, paginationQuery model.PaginationQuery) (int64, []model.Role, error) {
	var roles []model.Role
	var total int64
	query := r.db.Model(&model.Role{})
	if roleFilter.ID != nil {
		query.Where("id = ?", roleFilter.ID)
	}
	if roleFilter.Name != nil {
		query.Where("name = ?", roleFilter.Name)
	}
	total, err := Paginate(query, paginationQuery, &roles)
	return total, roles, err
}
