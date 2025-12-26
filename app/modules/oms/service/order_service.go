package service

import (
	"fmt"

	"github.com/maxlcoder/homework-backend/model"
	"github.com/maxlcoder/homework-backend/repository"
	"gorm.io/gorm"
)

type OrderServiceInterface interface {
	Create(model *model.PickingCar) (*model.PickingCar, error)
}

type OrderService struct {
	db                   *gorm.DB
	PickingCarRepository repository.PickingCarRepository
}

func NewService(db *gorm.DB, pickingCarRepository repository.PickingCarRepository) *PickingCarService {
	return &PickingCarService{
		db:                   db,
		PickingCarRepository: pickingCarRepository,
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
	find, _ := u.PickingCarRepository.FindBy(cond)
	if find != nil {
		return nil, fmt.Errorf("当前拣货车编号不可用，请检查")
	}
	err := u.PickingCarRepository.Create(pickingCar, nil)
	if err != nil {
		return nil, fmt.Errorf("拣货车创建失败: %w", err)
	}
	return pickingCar, nil
}
