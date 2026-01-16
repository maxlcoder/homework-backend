package request

// BinStoreRequest 库位创建请求（公共）
type BinStoreRequest struct {
	Code  string `form:"code" json:"code" binding:"required" label:"库位编号"`
	Num   uint   `form:"code" json:"num" label:"商品数量"`
	SkuId uint   `form:"sku_id" json:"sku_id" label:"SKU ID"`
}

// BinUpdateRequest 库位更新请求（公共）
type BinUpdateRequest struct {
	Code  string `form:"code" json:"code" binding:"omitempty" label:"库位编号"`
	Num   uint   `form:"code" json:"num" label:"商品数量"`
	SkuId uint   `form:"sku_id" json:"sku_id" label:"SKU ID"`
}

// BinListRequest 库位列表请求（公共）
type BinPageRequest struct {
	Page    int     `form:"page" binding:"required,min=1" label:"页码"`
	PerPage int     `form:"per_page" binding:"required,min=1,max=100" label:"每页数量"`
	Code    *string `form:"code" json:"code" binding:"omitempty" label:"库位编号"`
}
