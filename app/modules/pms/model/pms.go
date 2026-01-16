package model

import model2 "github.com/maxlcoder/homework-backend/model"

type Product struct {
	model2.BaseModel
	Name        string  `gorm:"size:60;not null;default:'';comment:产品名称"`
	Description string  `gorm:"size:2000;comment:产品描述"`
	Price       float64 `gorm:"type:decimal(12,2);not null;default:0"`
	State       int8    `gorm:"type:int;not null;default:0;comment:'状态 0:待生效 1:生效中'"`
}

type SpecificationType struct {
	model2.BaseModel
	Name string `gorm:"size:60;not null;default:'';comment:规格类型名称"`
}

type SpecificationValue struct {
	model2.BaseModel
	SpecificationTypeId uint   `gorm:"type:int;not null;default:0;comment:规格类型 ID"`
	Name                string `gorm:"size:60;not null;default:'';comment:规格类型值"`
}

// ProductSpecification 产品规格属性表
type ProductSpecification struct {
	model2.BaseModel
	ProductId            uint `gorm:"type:bigint;not null;default:0;comment:产品 ID"`
	SpecificationTypeId  uint `gorm:"type:bigint;not null;default:0;comment:规格类型 ID"`
	SpecificationValueId uint `gorm:"type:bigint;not null;default:0;comment:规格值 ID"`
}

type ProductSku struct {
	model2.BaseModel
	ProductId uint   `gorm:"type:bigint;not null;default:0;comment:产品 ID"`
	SkuCode   string `gorm:"size:60;not null;default:'';comment:sku code"`
}

// ProductSkuSpecification 多维的规格属性决定一个 sku
type ProductSkuSpecification struct {
	model2.BaseModel
	SkuId                uint `gorm:"type:bigint;not null;default:0;comment:SKU ID"`
	ProductId            uint `gorm:"type:bigint;not null;default:0;comment:产品 ID"`
	SpecificationTypeId  uint `gorm:"type:bigint;not null;default:0;comment:规格类型 ID"`
	SpecificationValueId uint `gorm:"type:bigint;not null;default:0;comment:规格值 ID"`
}

type ProductImage struct {
	model2.BaseModel
	ProductId uint   `gorm:"type:bigint;not null;default:0;comment:产品 ID"`
	Url       string `gorm:"not null;default:'';comment:图片地址"`
}

type ProductSkuImage struct {
	model2.BaseModel
	SkuId     uint   `gorm:"type:bigint;not null;default:0;comment:SKU ID"`
	ProductId uint   `gorm:"type:bigint;not null;default:0;comment:产品 ID"`
	Url       string `gorm:"not null;default:'';comment:图片地址"`
}
