package model

import model2 "github.com/maxlcoder/homework-backend/model"

type Permission struct {
	model2.BaseModel
	Name   string `gorm:"size:60;not null;default:''"`
	PATH   string `gorm:"size:60;not null;default:''"`
	Method string `gorm:"size:10;not null;default:''"`
}
