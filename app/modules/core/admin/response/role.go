package response

import (
	"github.com/jinzhu/copier"
	"github.com/maxlcoder/homework-backend/app/modules/core/model"
	"github.com/maxlcoder/homework-backend/app/response"
)

type RoleResponse struct {
	response.BaseResponse
	Name string `json:"name"`
}

func NewRoleResponse() *RoleResponse {
	return &RoleResponse{}
}

func (r *RoleResponse) FromModel(m model.Role) {
	copier.Copy(&r.BaseResponse, &m)
	copier.Copy(r, &m)
}
