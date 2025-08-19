package model

type Tenant struct {
	BaseModel
	Name string `gorm:"size:60;not null;default:'';unique"`
}

type TenantUser struct {
	BaseModel
	TenantId uint `gorm:"not null;default:0;comment:租户 ID"`
	UserId   uint `gorm:"not null;default:0;comment:用户 ID"`
}

type TenantAdmin struct {
	BaseModel
	TenantId uint `gorm:"not null;default:0;comment:租户 ID"`
	AdminId  uint `gorm:"not null;default:0;comment:管理员 ID"`
}
