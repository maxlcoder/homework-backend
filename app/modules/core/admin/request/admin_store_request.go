package request

import (
	base_request "github.com/maxlcoder/homework-backend/app/request"
)

type AdminStoreRequest struct {
	Name     string                   `json:"name" binding:"required,min=1,max=30" label:"用户名"`
	Password string                   `json:"password" binding:"required" label:"密码"`
	Roles    []base_request.IdRequest `json:"roles" binding:"required,dive" label:"角色"`
}
