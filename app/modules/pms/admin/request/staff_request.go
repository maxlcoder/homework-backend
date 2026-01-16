package request

// StaffStoreRequest 仓库人员创建请求
type StaffStoreRequest struct {
	Name  string `json:"name" binding:"required,min=1,max=60"`
	State int8   `json:"state" binding:"omitempty,oneof=0 1 2 3"`
}

// StaffUpdateRequest 仓库人员更新请求
type StaffUpdateRequest struct {
	Name  string `json:"name" binding:"omitempty,min=1,max=60"`
	State int8   `json:"state" binding:"omitempty,oneof=0 1 2 3"`
}

// StaffFilterRequest 仓库人员列表过滤请求
type StaffFilterRequest struct {
	Name    string `form:"name" binding:"omitempty,max=60"`
	State   int8   `form:"state" binding:"omitempty,oneof=0 1 2 3"`
	Page    int    `form:"page" binding:"omitempty,min=1"`
	PerPage int    `form:"limit" binding:"omitempty,min=1,max=100"`
}
