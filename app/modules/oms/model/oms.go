package model

import (
	model2 "github.com/maxlcoder/homework-backend/model"
	"github.com/shopspring/decimal"
)

// Platform 平台

type Platform struct {
	model2.BaseModel
	Name string `gorm:"size:60;not null;default:'';comment:名称"`
}

// Order 订单
type Order struct {
	model2.BaseModel
	StoreOrderId   uint            `gorm:"not null;default:0;comment:店铺订单 ID"`
	Address        string          `gorm:"size:200;not null;default:'';comment:详细地址"`
	ProvinceCode   string          `gorm:"size:30;not null;default:'';comment:省编号"`
	CityCode       string          `gorm:"size:30;not null;default:'';comment:市编号"`
	CountyCode     string          `gorm:"size:30;not null;default:'';comment:区编号"`
	TotalAmount    decimal.Decimal `gorm:"type:decimal(14,4);comment:金额"`
	Currency       string          `gorm:"size:10;not null;default:'';comment:币种"`
	TotalAmountCny decimal.Decimal `gorm:"type:decimal(14,4);comment:人民币金额"`
	TotalAmountUsd decimal.Decimal `gorm:"type:decimal(14,4);comment:美元金额"`
	State          int             `gorm:"type:tinyint;comment:状态"`
}

// OrderItem 订单项
type OrderItem struct {
	model2.BaseModel
	ProductName string          `gorm:"size:60;not null;default:'';comment:商品名称"`
	Quantity    int16           `gorm:"type:int;not null;default:0;comment:商品数量"`
	Price       decimal.Decimal `gorm:"type:decimal(14,4);comment:单价"`
	Currency    string          `gorm:"size:10;not null;default:'';comment:币种"`
	PriceCny    decimal.Decimal `gorm:"type:decimal(14,4);comment:人民币价格"`
	PriceUsd    decimal.Decimal `gorm:"type:decimal(14,4);comment:美元价格"`
}

// 订单操作日志
type OrderOperateLog struct {
	model2.BaseModel
	AdminId uint   `gorm:"not null;default:0;管理员 ID"`
	Content string `gorm:"comment:内容"`
}

// StoreOrder 店铺订单
type StoreOrder struct {
	model2.BaseModel
	Address        string          `gorm:"size:200;not null;default:'';comment:详细地址"`
	ProvinceCode   string          `gorm:"size:30;not null;default:'';comment:省编号"`
	CityCode       string          `gorm:"size:30;not null;default:'';comment:市编号"`
	CountyCode     string          `gorm:"size:30;not null;default:'';comment:区编号"`
	TotalAmount    decimal.Decimal `gorm:"type:decimal(14,4);comment:金额"`
	Currency       string          `gorm:"size:10;not null;default:'';comment:币种"`
	TotalAmountCny decimal.Decimal `gorm:"type:decimal(14,4);comment:人民币金额"`
	TotalAmountUsd decimal.Decimal `gorm:"type:decimal(14,4);comment:美元金额"`
	State          int             `gorm:"type:tinyint;comment:状态"`
}

// StoreOrderItem 店铺订单项
type StoreOrderItem struct {
	model2.BaseModel
	ProductName string          `gorm:"size:60;not null;default:'';comment:商品名称"`
	Quantity    int16           `gorm:"type:int;not null;default:0;comment:商品数量"`
	Price       decimal.Decimal `gorm:"type:decimal(14,4);comment:单价"`
	Currency    string          `gorm:"size:10;not null;default:'';comment:币种"`
	PriceCny    decimal.Decimal `gorm:"type:decimal(14,4);comment:人民币价格"`
	PriceUsd    decimal.Decimal `gorm:"type:decimal(14,4);comment:美元价格"`
}

type WebhookLog struct {
	model2.BaseModel
	UniqueNum    string `gorm:"size:20;not null;default:'';comment:唯一编号"`
	PlatformType string `gorm:"size:20;not null;default:'';comment:平台类型"`
	Content      string `gorm:"comment:内容"`
}
