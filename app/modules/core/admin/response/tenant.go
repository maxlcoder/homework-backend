package response

import (
	"github.com/maxlcoder/homework-backend/app/modules/core/model"
	"github.com/maxlcoder/homework-backend/app/response"
)

// TenantResponse 租户响应结构（公共）
type TenantResponse struct {
	response.BaseResponse
	Name string `json:"name"`
}

// 转换函数 - 将 Model 转换为 Response
func ToTenantResponse(m model.Tenant) TenantResponse {
	var r TenantResponse
	r.FromBaseModel(m.BaseModel)
	r.Name = m.Name
	return r
}

// 批量转换函数
func ToTenantResponses(models []model.Tenant) []TenantResponse {
	return response.ConvertSlice[model.Tenant, TenantResponse](models)
}
