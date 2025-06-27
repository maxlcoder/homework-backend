package model

type Menu struct {
	Model
	Name       string `gorm:"size:60;not null;default:''"`
	ParentID   uint   `gorm:"not null;default:0"`
	Sort       int    `gorm:"not null;default:0"`
	IsDisabled bool   `gorm:"default:0"`
}

type MenuPermission struct {
	MenuID       uint `gorm:"not null;default:0"`
	PermissionID uint `gorm:"not null;default:0"`
}
