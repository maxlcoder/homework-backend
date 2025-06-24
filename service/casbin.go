package service

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
)

var Enforcer *casbin.Enforcer

// 初始化 casbin
func InitCasbin(db *gorm.DB) error {
	adapter, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		return err
	}
	enforcer, err := casbin.NewEnforcer("./casbin_model.conf", adapter)
	if err != nil {
		return err
	}
	err = enforcer.LoadPolicy()
	if err != nil {
		return err
	}
	Enforcer = enforcer
	return nil
}
