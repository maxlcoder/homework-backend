package repository

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// 定义完整条件对象，避免多重条件参数过多
type ConditionScope struct {
	StructCond interface{}
	MapCond    map[string]interface{}
	Scopes     []func(*gorm.DB) *gorm.DB
	Preloads   []string
	Order      []string
	OrderBy    clause.OrderBy
	Group      string
	Having     string
}

func (cs ConditionScope) Apply(db *gorm.DB) *gorm.DB {
	query := db
	// struct 条件
	if cs.StructCond != nil {
		query = query.Where(cs.StructCond)
	}
	// map 条件
	if len(cs.MapCond) > 0 {
		query = query.Where(cs.MapCond)
	}
	// 额外的 scopes 条件
	if len(cs.Scopes) > 0 {
		query = query.Scopes(cs.Scopes...)
	}
	// preload
	for _, preload := range cs.Preloads {
		query = query.Preload(preload)
	}
	// 排序
	for _, order := range cs.Order {
		query = query.Order(order)
	}

	if len(cs.Group) > 0 {
		query = query.Group(cs.Group)
	}

	return query
}

// 定义查询接口
type QueryCondition[T any] interface {
	Apply(*gorm.DB) *gorm.DB
}

// 实现的三种形式
// 1. struct 条件
type StructCondition[T any] struct {
	Cond     T
	Preloads []string
}

// 组装查询条件
func (s StructCondition[T]) Apply(db *gorm.DB) *gorm.DB {
	query := db.Where(&s.Cond)
	if len(s.Preloads) > 0 {
		for _, preload := range s.Preloads {
			query = query.Preload(preload)
		}
	}
	return query
}

// 2. map 条件
type MapCondition[T any] struct {
	Cond     map[string]interface{}
	Preloads []string
}

func (m MapCondition[T]) Apply(db *gorm.DB) *gorm.DB {
	query := db.Where(m.Cond)
	if len(m.Preloads) > 0 {
		for _, preload := range m.Preloads {
			query = query.Preload(preload)
		}
	}
	return query
}

// 3. 自定义 builder
type FuncCondition[T any] struct {
	Builder func(db *gorm.DB) *gorm.DB // 内建函数
}

func (f FuncCondition[T]) Apply(db *gorm.DB) *gorm.DB {
	return f.Builder(db)
}
