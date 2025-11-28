package model

type Config struct {
	BaseModel
	ItemKey   string `gorm:"size:20;not null;default:''"`
	ItemValue string `gorm:"size:200;not null;default:''"`
}
