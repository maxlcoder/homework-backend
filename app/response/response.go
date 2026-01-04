package response

import (
	"github.com/jinzhu/copier"
	base_model "github.com/maxlcoder/homework-backend/model"
)

// Mapper 定义模型到响应的转换接口
type Mapper[M any] interface {
	FromModel(M)
}

// BaseMapper 提供通用的转换功能
type BaseMapper[M any] struct{}

// FromModel 通用的模型到响应转换方法
func (bm *BaseMapper[M]) FromModel(m M) {
	// 默认实现，由具体的响应类型重写
}

// BaseResponse 基础响应结构体，嵌入到所有响应中
type BaseResponse struct {
	ID        uint     `json:"id"`
	CreatedAt JSONTime `json:"created_at"`
	UpdatedAt JSONTime `json:"updated_at"`
}

// FromBaseModel 从基础模型转换基础响应字段
func (br *BaseResponse) FromBaseModel(bm base_model.BaseModel) {
	br.ID = bm.ID
	br.CreatedAt = JSONTime(bm.CreatedAt)
	br.UpdatedAt = JSONTime(bm.UpdatedAt)
}

// ConvertModelToResponse 通用的模型到响应转换函数
func ConvertModelToResponse[M any, R any](model M) R {
	var response R
	copier.Copy(&response, &model)
	return response
}

// ConvertSlice 通用的切片转换函数
func ConvertSlice[M any, R any](models []M) []R {
	responses := make([]R, len(models))
	for i, m := range models {
		responses[i] = ConvertModelToResponse[M, R](m)
	}
	return responses
}

// 分页响应
type PageResponse[T any] struct {
	Page    int   `json:"page"`
	PerPage int   `json:"per_page"`
	Total   int64 `json:"total"`
	Data    []T   `json:"data"`
}

// 一般作为创建对象后返回
type DataId struct {
	ID uint `json:"id"`
}

// BuildPageResponse 组装分页结果（支持泛型转换）
func BuildPageResponse[M any, R any](models []M, total int64, page, perPage int) PageResponse[R] {
	return PageResponse[R]{
		Page:    page,
		PerPage: perPage,
		Total:   total,
		Data:    ConvertSlice[M, R](models),
	}
}

// BuildPageResponseWithMapper 使用自定义转换函数组装分页结果
func BuildPageResponseWithMapper[M any, R any](models []M, total int64, page, perPage int, mapper func(M) R) PageResponse[R] {
	responses := make([]R, len(models))
	for i, m := range models {
		responses[i] = mapper(m)
	}
	return PageResponse[R]{
		Page:    page,
		PerPage: perPage,
		Total:   total,
		Data:    responses,
	}
}
