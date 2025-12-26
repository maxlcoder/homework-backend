package request

// BinStoreRequest 库位创建请求（公共）
type BinStoreRequest struct {
	Code   string `form:"code" json:"code" binding:"required" label:"库位编号"`
	Name   string `form:"name" json:"name" binding:"required" label:"库位名称"`
	Status int    `form:"status" json:"status" binding:"omitempty" label:"状态"`
}

// BinUpdateRequest 库位更新请求（公共）
type BinUpdateRequest struct {
	Code   string `form:"code" json:"code" binding:"omitempty" label:"库位编号"`
	Name   string `form:"name" json:"name" binding:"omitempty" label:"库位名称"`
	Status int    `form:"status" json:"status" binding:"omitempty" label:"状态"`
}

// BinListRequest 库位列表请求（公共）
type BinListRequest struct {
	Page    int    `form:"page" binding:"required,min=1" label:"页码"`
	PerPage int    `form:"per_page" binding:"required,min=1,max=100" label:"每页数量"`
	Code    string `form:"code" json:"code" binding:"omitempty" label:"库位编号"`
	Status  int    `form:"status" json:"status" binding:"omitempty" label:"状态"`
}
