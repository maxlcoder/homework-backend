package service

import (
	"fmt"

	"github.com/maxlcoder/homework-backend/app/modules/wms/model"
	"github.com/maxlcoder/homework-backend/app/request"
	base_model "github.com/maxlcoder/homework-backend/model"
	"github.com/maxlcoder/homework-backend/repository"
	"gorm.io/gorm"
)

type PickingCarServiceInterface interface {
	Page(pageRequest request.PageRequest) ([]model.PickingCar, int64, error)
	Create(model *model.PickingCar) (*model.PickingCar, error)
	Update(model *model.PickingCar) (*model.PickingCar, error)
	Delete(id uint) error
	FindById(id uint) (*model.PickingCar, error)
}

type PickingCarService struct {
	db *gorm.DB
}

func NewPickingCarService(db *gorm.DB) PickingCarServiceInterface {
	return &PickingCarService{
		db: db,
	}
}

func (u *PickingCarService) Create(pickingCar *model.PickingCar) (*model.PickingCar, error) {
	// 判断是否存在已经适用的名称
	filer := model.PickingCarFilter{
		Code: &pickingCar.Code,
	}
	cond := repository.ConditionScope{
		StructCond: filer,
	}
	find, _ := repository.NewBaseRepository[model.PickingCar](u.db).FindBy(cond)
	if find != nil {
		return nil, fmt.Errorf("当前拣货车编号不可用，请检查")
	}
	err := repository.NewBaseRepository[model.PickingCar](u.db).Create(pickingCar, nil)
	if err != nil {
		return nil, fmt.Errorf("拣货车创建失败: %w", err)
	}
	return pickingCar, nil
}

func (u *PickingCarService) Update(pickingCar *model.PickingCar) (*model.PickingCar, error) {
	// 检查 code
	filter := model.PickingCarFilter{
		Code: &pickingCar.Code,
	}
	cond := repository.ConditionScope{
		StructCond: filter,
		Scopes: []func(*gorm.DB) *gorm.DB{
			func(db *gorm.DB) *gorm.DB {
				return db.Where("id <> ?", pickingCar.ID)
			},
		},
	}
	find, _ := repository.NewBaseRepository[model.PickingCar](u.db).FindBy(cond)
	if find != nil && find.ID != pickingCar.ID {
		return nil, fmt.Errorf("当前拣货车编号不可用，请检查")
	}
	err := repository.NewBaseRepository[model.PickingCar](u.db).Update(pickingCar, nil)
	if err != nil {
		return nil, fmt.Errorf("拣货车更新失败: %w", err)
	}
	return pickingCar, nil
}

func (u *PickingCarService) Delete(id uint) error {
	err := repository.NewBaseRepository[model.PickingCar](u.db).DeleteById(id, nil)
	if err != nil {
		return fmt.Errorf("拣货车删除失败: %w", err)
	}
	return nil
}

// Page 获取拣货车分页列表
func (u *PickingCarService) Page(pageRequest request.PageRequest) ([]model.PickingCar, int64, error) {
	cond := repository.ConditionScope{}

	// 创建分页参数
	pagination := base_model.Pagination{
		Page:    pageRequest.Page,
		PerPage: pageRequest.PerPage,
	}

	// 查询数据
	count, pickingCars, err := repository.NewBaseRepository[model.PickingCar](u.db).Page(cond, pagination)
	if err != nil {
		return nil, 0, fmt.Errorf("获取拣货车列表失败: %w", err)
	}

	return pickingCars, count, nil
}

// FindById 根据 id 查询拣货车
func (u *PickingCarService) FindById(id uint) (*model.PickingCar, error) {
	pickingCar, err := repository.NewBaseRepository[model.PickingCar](u.db).FindById(id)
	if err != nil {
		return nil, fmt.Errorf("获取拣货车失败: %w", err)
	}
	return pickingCar, nil
}
