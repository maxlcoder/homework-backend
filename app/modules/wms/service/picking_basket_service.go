package service

import (
	"fmt"

	"github.com/maxlcoder/homework-backend/app/modules/wms/admin/request"
	"github.com/maxlcoder/homework-backend/app/modules/wms/model"
	"github.com/maxlcoder/homework-backend/app/modules/wms/repository"
	base_model "github.com/maxlcoder/homework-backend/model"
	base_repository "github.com/maxlcoder/homework-backend/repository"
	"gorm.io/gorm"
)

type PickingBasketServiceInterface interface {
	Page(pageRequest request.PickingBasketPageRequest) ([]model.PickingBasket, int64, error)
	Create(model *model.PickingBasket) (*model.PickingBasket, error)
	Update(model *model.PickingBasket) (*model.PickingBasket, error)
	Delete(id uint) error
	FindById(id uint) (*model.PickingBasket, error)
}

type PickingBasketService struct {
	db                      *gorm.DB
	PickingBasketRepository repository.PickingBasketRepository
}

func NewPickingBasketService(db *gorm.DB, pickingBasketRepository repository.PickingBasketRepository) PickingBasketServiceInterface {
	return &PickingBasketService{
		db:                      db,
		PickingBasketRepository: pickingBasketRepository,
	}
}

func (u *PickingBasketService) Page(pageRequest request.PickingBasketPageRequest) ([]model.PickingBasket, int64, error) {
	cond := base_repository.ConditionScope{}

	if pageRequest.Code != "" {
		cond.StructCond = model.PickingBasketFilter{
			Code: &pageRequest.Code,
		}
	}

	// 创建分页参数
	pagination := base_model.Pagination{
		Page:    pageRequest.Page,
		PerPage: pageRequest.PerPage,
	}

	// 查询数据
	count, pickingCars, err := u.PickingBasketRepository.Page(cond, pagination)
	if err != nil {
		return nil, 0, fmt.Errorf("获取拣货车列表失败: %w", err)
	}

	return pickingCars, count, nil
}

func (u *PickingBasketService) Create(pickingBasket *model.PickingBasket) (*model.PickingBasket, error) {
	// 判断是否存在已经适用的名称
	filer := model.PickingBasketFilter{
		Code: &pickingBasket.Code,
	}
	cond := base_repository.ConditionScope{
		StructCond: filer,
	}
	find, _ := u.PickingBasketRepository.FindBy(cond)
	if find != nil {
		return nil, fmt.Errorf("当前拣货框编号不可用，请检查")
	}
	err := u.PickingBasketRepository.Create(pickingBasket, nil)
	if err != nil {
		return nil, fmt.Errorf("拣货框创建失败: %w", err)
	}
	return pickingBasket, nil
}

func (u *PickingBasketService) Update(pickingBasket *model.PickingBasket) (*model.PickingBasket, error) {

	return nil, nil
}

func (u *PickingBasketService) Delete(id uint) error {
	err := u.PickingBasketRepository.DeleteById(id, nil)
	if err != nil {
		return fmt.Errorf("拣货篮删除失败: %w", err)
	}
	return nil
}

func (u *PickingBasketService) FindById(id uint) (*model.PickingBasket, error) {
	pickingBasket, err := u.PickingBasketRepository.FindById(id)
	if err != nil {
		return nil, fmt.Errorf("获取拣货篮失败: %w", err)
	}
	return pickingBasket, nil
}
