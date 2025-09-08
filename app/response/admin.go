package response

import (
	"github.com/jinzhu/copier"
	"github.com/maxlcoder/homework-backend/model"
)

type AdminResponse struct {
	BaseResponse
	Name  string         `json:"name"`
	Email string         `json:"email"`
	Age   uint8          `json:"age"`
	Roles []RoleResponse `json:"roles"`
}

func NewAdminResponse() *AdminResponse {
	return &AdminResponse{}
}

func (r *AdminResponse) FromModel(m model.Admin) {
	copier.Copy(&r.BaseResponse, &m)
	copier.Copy(r, &m)
}
