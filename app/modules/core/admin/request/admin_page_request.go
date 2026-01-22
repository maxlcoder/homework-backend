package request

import "github.com/maxlcoder/homework-backend/app/request"

// AdminPageRequest 管理员列表请求（公共）
type AdminPageRequest struct {
	request.PageRequest
	Name *string `form:"name" json:"name" binding:"omitempty" label:"用户名"`
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
