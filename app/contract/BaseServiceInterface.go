package contract

type BaseServiceInterface interface {
	// 增删改查
	Create(i interface{}) (interface{}, error)
}
