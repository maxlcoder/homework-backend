package model

type Menu struct {
	BaseModel
	Name        string        `gorm:"size:60;not null;default:''"`
	ParentID    uint          `gorm:"not null;default:0"`
	Sort        int           `gorm:"not null;default:0"`
	IsDisabled  bool          `gorm:"default:0"`
	Number      string        `gorm:"not null;default:'';uniqueIndex"`
	Children    []*Menu       `gorm:"-"`
	Permissions []*Permission `gorm:"-"`
}

type MenuPermission struct {
	MenuID       uint `gorm:"not null;default:0;uniqueIndex:uq_menu_permission"`
	PermissionID uint `gorm:"not null;default:0;uniqueIndex:uq_menu_permission"`
}
