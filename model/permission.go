package model

type Permission struct {
	Model
	Name   string `gorm:"size:60;not null;default:''"`
	PATH   string `gorm:"size:60;not null;default:''"`
	Method string `gorm:"size:10;not null;default:''"`
}
