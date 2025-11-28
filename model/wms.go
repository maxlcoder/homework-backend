package model

import "time"

// 仓库系统

type Warehouse struct {
	BaseModel
	Name         string `gorm:"size:100;not null;default:'';comment:名称"`
	Address      string `gorm:"size:200;not null;default:'';comment:详细地址"`
	ProvinceCode string `gorm:"size:30;not null;default:'';comment:省编号"`
	CityCode     string `gorm:"size:30;not null;default:'';comment:市编号"`
	CountyCode   string `gorm:"size:30;not null;default:'';comment:区编号"`
	Area         string `gorm:"size:30;not null;default:'';comment:面积"`
}

// Bin  库位
type Bin struct {
	BaseModel
	Code           string `gorm:"size:60;not null;default:'';comment:库位编号"`
	SkuId          uint   `gorm:"not null;default:0;comment:当前存放 SKU ID"`
	Num            int16  `gorm:"not null;default:0;comment:SKU 商品数量"`
	ExpirationDate string `gorm:"comment:过期时间"`
}

// 仓库人员
type Staff struct {
	BaseModel
	Name  string `gorm:"size:60;not null;default:'';comment:姓名"`
	State int8   `gorm:"not null;default:0;comment:状态"`
}

// ---------- 入库相关 ----------

// 入库单子
type StockOrder struct {
	BaseModel
	Code   string `gorm:"size:60;not null;default:'';comment:编号"`
	InDate string `gorm:"size:30;not null;default:'';入库日期"`
}

// 拆包入库
type StockOrderProduct struct {
	BaseModel
	StockOrderId uint  `gorm:"not null;default:0;comment:入库单 ID"`
	ProductId    uint  `gorm:"not null;default:0;comment:商品 ID"`
	Num          int16 `gorm:"not null;default:0;comment:数量"`
	BinId        uint  `gorm:"not null;default:0;comment:库位 ID"`
}

// StockTask 入库任务
type StockTask struct {
	BaseModel
	Code         string `gorm:"size:60;not null;default:'';comment:编号"`
	StockOrderId uint   `gorm:"not null;default:0;comment:入库单 ID"`
}

type StockTaskProduct struct {
	BaseModel
	StockTaskId uint  `gorm:"not null;default:0;comment:入库任务 ID"`
	ProductId   uint  `gorm:"not null;default:0;comment:商品 ID"`
	Num         int16 `gorm:"not null;default:0;comment:数量"`
	BinId       uint  `gorm:"not null;default:0;comment:库位 ID"`
}

// ---------- 出库相关 ----------

// 拣货车
type PickingCar struct {
	BaseModel
	Code           string `gorm:"size:60;not null;default:'';comment:编号"`
	MaxBasketCount int8   `gorm:"not null;default:0;comment:最大拣货框数'"`
}

type PickingCarFilter struct {
	ID        *uint
	Code      *string `form:"code"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

// 拣货框 （一车多框，也就是一次拣多个订单）
type PickingBasket struct {
	BaseModel
	Code string `gorm:"size:60;not null;default:'';comment:编号"`
}

// 拣货任务
type PickingTask struct {
	BaseModel
	PickingCarId uint `gorm:"not null;default:0;comment:拣货车 ID"`
	StaffId      uint `gorm:"not null;default:0;comment:拣货员工 ID"`
}

// 拣货任务关联拣货框
type PickingTaskBasket struct {
	BaseModel
	PickingTaskId uint `gorm:"not null;default:0;comment:拣货任务 ID"`
	PickingCarId  uint `gorm:"not null;default:0;comment:拣货车 ID"`
	OrderId       uint `gorm:"not null;default:0;comment:订单 ID"`
}

// 拣货框订单商品
type PickingTaskBasketProduct struct {
	BaseModel
	PickingTaskBasketId uint   `gorm:"not null;default:0;comment:拣货框 ID"`
	PickingTaskId       uint   `gorm:"not null;default:0;comment:拣货任务 ID"`
	PickingCarId        uint   `gorm:"not null;default:0;comment:拣货车 ID"`
	OrderId             uint   `gorm:"not null;default:0;comment:订单 ID"`
	SkuCode             string `gorm:"not null;default:'';comment:SKU CODE"`
	Num                 int16  `gorm:"not null;default:0;comment:应拣数量"`
	ActualNum           int16  `gorm:"not null;default:0;comment:实拣数量"`
}
