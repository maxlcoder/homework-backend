package service

import (
	"fmt"

	"github.com/maxlcoder/homework-backend/app/modules/wms/model"
	"github.com/maxlcoder/homework-backend/app/modules/wms/repository"
	base_repository "github.com/maxlcoder/homework-backend/repository"
	"gorm.io/gorm"
)

type BinServiceInterface interface {
	Create(model *model.Bin) (*model.Bin, error)
	GetByID(id uint) (*model.Bin, error)
	Update(model *model.Bin) (*model.Bin, error)
}

type BinService struct {
	db            *gorm.DB
	BinRepository repository.BinRepository
}

func NewBinService(db *gorm.DB, binRepository repository.BinRepository) BinServiceInterface {
	return &BinService{
		db:            db,
		BinRepository: binRepository,
	}
}

func (u *BinService) Create(bin *model.Bin) (*model.Bin, error) {
	// 判断是否存在已经适用的名称
	filer := model.Bin{
		Code: bin.Code,
	}
	cond := base_repository.ConditionScope{
		StructCond: filer,
	}
	find, _ := u.BinRepository.FindBy(cond)
	if find != nil {
		return nil, fmt.Errorf("当前库位编号不可用，请检查")
	}
	err := u.BinRepository.Create(bin, nil)
	if err != nil {
		return nil, fmt.Errorf("库位创建失败: %w", err)
	}
	return bin, nil
}

func (u *BinService) GetByID(id uint) (*model.Bin, error) {
	bin, err := u.BinRepository.FindById(id)
	if err != nil {
		return nil, fmt.Errorf("库位查询失败: %w", err)
	}
	if bin == nil {
		return nil, fmt.Errorf("库位不存在")
	}
	return bin, nil
}

func (u *BinService) Update(bin *model.Bin) (*model.Bin, error) {
	// 判断是否存在已经适用的名称（排除自身）
	filer := model.Bin{
		Code: bin.Code,
	}
	cond := base_repository.ConditionScope{
		StructCond: filer,
	}
	find, _ := u.BinRepository.FindBy(cond)
	if find != nil && find.ID != bin.ID {
		return nil, fmt.Errorf("当前库位编号不可用，请检查")
	}

	err := u.BinRepository.Update(bin, u.db)
	if err != nil {
		return nil, fmt.Errorf("库位更新失败: %w", err)
	}
	return bin, nil
}
