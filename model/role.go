package model

type Role struct {
	Model
	Name string `gorm:"size:60;not null;default:''"`
}
