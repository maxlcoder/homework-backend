package response

import (
	"github.com/jinzhu/copier"
	"github.com/maxlcoder/homework-backend/app/modules/wms/model"
	"github.com/maxlcoder/homework-backend/app/response"
)

// 通用响应
// BinResponse 库位响应结构（公共）
type BinResponse struct {
	response.BaseResponse
	Code   string `json:"code"`
	Name   string `json:"name"`
	Status int    `json:"status"`
}

// PickingCarResponse 拣货车响应结构（公共）
type PickingCarResponse struct {
	response.BaseResponse
	Code   string `json:"code"`
	Status int    `json:"status"`
}

func NewPickingCarResponse() *PickingCarResponse {
	return &PickingCarResponse{}
}

func (r *PickingCarResponse) FromModel(m model.PickingCar) {
	copier.Copy(&r.BaseResponse, &m)
	copier.Copy(r, &m)
}

// StaffResponse 员工响应结构（公共）
type StaffResponse struct {
	response.BaseResponse
	Name      string `json:"name"`
	Code      string `json:"code"`
	Position  string `json:"position"`
	Status    int    `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func NewStaffResponse() *StaffResponse {
	return &StaffResponse{}
}

func (r *StaffResponse) FromModel(m model.Staff) {
	copier.Copy(&r.BaseResponse, &m)
	copier.Copy(r, &m)
}

// PickingBasketResponse 拣货框响应结构（公共）
type PickingBasketResponse struct {
	response.BaseResponse
	ID        uint   `json:"id"`
	Code      string `json:"code"`
	Status    int    `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func NewPickingBasketResponse() *PickingBasketResponse {
	return &PickingBasketResponse{}
}

func (r *PickingBasketResponse) FromModel(m model.PickingBasket) {
	copier.Copy(&r.BaseResponse, &m)
	copier.Copy(r, &m)
}
