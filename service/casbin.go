package service

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
)

// 初始化 casbin
func NewCasbin(db *gorm.DB) (*casbin.Enforcer, error) {
	// 初始化 casbin 相关表
	adapter, err := gormadapter.NewAdapterByDBUseTableName(db, "", "casbin_rule")
	if err != nil {
		return nil, err
	}
	enforcer, err := casbin.NewEnforcer("config/rbac_with_domains_model.conf", adapter)
	if err != nil {
		return nil, err
	}
	err = enforcer.LoadPolicy()
	if err != nil {
		return nil, err
	}
	return enforcer, nil
}
