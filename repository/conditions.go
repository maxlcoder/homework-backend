package repository

import "gorm.io/gorm"

// 定义查询接口
type QueryCondition[T any] interface {
	Apply(*gorm.DB) *gorm.DB
}

// 实现的三种形式
// 1. struct 条件
type StructCondition[T any] struct {
	Cond T
}

// 组装查询条件
func (s StructCondition[T]) Apply(db *gorm.DB) *gorm.DB {
	return db.Where(&s.Cond)
}

// 2. map 条件
type MapCondition[T any] struct {
	Cond map[string]interface{}
}

func (m MapCondition[T]) Apply(db *gorm.DB) *gorm.DB {
	return db.Where(m.Cond)
}

// 3. 自定义 builder
type FuncCondition[T any] struct {
	Builder func(db *gorm.DB) *gorm.DB // 内建函数
}

func (f FuncCondition[T]) Apply(db *gorm.DB) *gorm.DB {
	return f.Builder(db)
}
