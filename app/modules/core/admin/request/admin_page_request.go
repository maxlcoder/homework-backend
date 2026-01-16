package request

// AdminPageRequest 管理员列表请求（公共）
type AdminPageRequest struct {
	Page    int     `form:"page" binding:"min=1" label:"页码" default:"1"`
	PerPage int     `form:"per_page" binding:"min=1,max=100" label:"每页数量" default:"10"`
	Name    *string `form:"name" json:"name" binding:"omitempty" label:"用户名"`
}

// SetDefaults 为 AdminPageRequest 设置默认值
func (r *AdminPageRequest) SetDefaults() {
	if r.Page <= 0 {
		r.Page = 1
	}
	if r.PerPage <= 0 || r.PerPage > 100 {
		r.PerPage = 10
	}
}
