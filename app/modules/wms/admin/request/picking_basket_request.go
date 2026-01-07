package request

import base_request "github.com/maxlcoder/homework-backend/app/request"

type PickingBasketPageRequest struct {
	base_request.PageRequest
	Code string `form:"code"`
}

type PickingBasketStoreRequest struct {
	Code string `json:"code" binding:"required,min=1,max=60" label:"编号"`
}

type PickingBasketUpdateRequest struct {
	ID   uint   `json:"id" label:"ID"`
	Code string `json:"code" binding:"required,min=1,max=60" label:"编号"`
}
