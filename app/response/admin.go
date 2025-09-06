package response

import "github.com/jinzhu/copier"

type AdminResponse struct {
	BaseResponse
	Name  string         `json:"name"`
	Email string         `json:"email"`
	Age   uint8          `json:"age"`
	Roles []RoleResponse `json:"roles"`
}

func ToAdminResponse(T any) AdminResponse {
	var adminResponse AdminResponse
	copier.Copy(&adminResponse.BaseResponse, T)
	copier.Copy(&adminResponse, T)
	return adminResponse
}
