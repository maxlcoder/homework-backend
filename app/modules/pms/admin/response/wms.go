package response

import (
	"github.com/maxlcoder/homework-backend/app/modules/wms/model"
	"github.com/maxlcoder/homework-backend/app/response"
)

// 通用响应

// BinResponse 库位响应结构（公共）
type BinResponse struct {
	response.BaseResponse
	Code string       `json:"code"`
	Sku  *SkuResponse `json:"sku"`
	Num  int          `json:"num"`
}

type SkuResponse struct {
	response.BaseResponse
	Name string `json:"name"`
	Code string `json:"code"`
}

// PickingCarResponse 拣货车响应结构（公共）
type PickingCarResponse struct {
	response.BaseResponse
	Code           string `json:"code"`
	MaxBasketCount int8   `json:"max_basket_count"`
}

// StaffResponse 员工响应结构（公共）
type StaffResponse struct {
	response.BaseResponse
	Name   string `json:"name"`
	Code   string `json:"code"`
	Status int8   `json:"status"`
}

// PickingBasketResponse 拣货框响应结构（公共）
type PickingBasketResponse struct {
	response.BaseResponse
	Code string `json:"code"`
}

// 转换函数 - 将 Model 转换为 Response

func ToBinResponse(m model.Bin) BinResponse {
	var r BinResponse
	r.FromBaseModel(m.BaseModel)
	r.Code = m.Code
	// 注意：Model 中没有 Name 和 Status 字段，可能需要从其他关联数据获取
	// r.Name = m.Name
	// r.Status = m.Status
	return r
}

func ToPickingCarResponse(m model.PickingCar) PickingCarResponse {
	var r PickingCarResponse
	r.FromBaseModel(m.BaseModel)
	r.Code = m.Code
	r.MaxBasketCount = m.MaxBasketCount
	return r
}

func ToStaffResponse(m model.Staff) StaffResponse {
	var r StaffResponse
	r.FromBaseModel(m.BaseModel)
	r.Name = m.Name
	r.Code = m.Code // 注意：Model 中没有 Code 字段，可能需要生成或关联其他数据
	r.Status = int8(m.State)
	return r
}

func ToPickingBasketResponse(m model.PickingBasket) PickingBasketResponse {
	var r PickingBasketResponse
	r.FromBaseModel(m.BaseModel)
	r.Code = m.Code
	return r
}

// 批量转换函数
func ToBinResponses(models []model.Bin) []BinResponse {
	return response.ConvertSlice[model.Bin, BinResponse](models)
}

func ToPickingCarResponses(models []model.PickingCar) []PickingCarResponse {
	return response.ConvertSlice[model.PickingCar, PickingCarResponse](models)
}

func ToStaffResponses(models []model.Staff) []StaffResponse {
	return response.ConvertSlice[model.Staff, StaffResponse](models)
}

func ToPickingBasketResponses(models []model.PickingBasket) []PickingBasketResponse {
	return response.ConvertSlice[model.PickingBasket, PickingBasketResponse](models)
}
