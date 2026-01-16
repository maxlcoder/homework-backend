package model

import "gorm.io/gorm"

func Models() []interface{} {
	return []interface{}{
		// 初始化 wms 数据库
		&Admin{},
		&AdminRole{},
		&Permission{},
		&Menu{},
		&MenuPermission{},
		&Role{},
		&RoleMenu{},
		&RolePermission{},

		&Tenant{},
	}
}

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(Models()...)
}
