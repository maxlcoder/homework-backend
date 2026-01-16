package request

type TenantUpdateRequest struct {
	Name string `json:"name" binding:"omitempty,min=1,max=60" label:"租户名称"`
}

// TenantPageRequest 租户列表请求（公共）
type TenantPageRequest struct {
	Page    int     `form:"page" binding:"required,min=1" label:"页码"`
	PerPage int     `form:"per_page" binding:"required,min=1,max=100" label:"每页数量"`
	Name    *string `form:"name" json:"name" binding:"omitempty" label:"租户名称"`
}
