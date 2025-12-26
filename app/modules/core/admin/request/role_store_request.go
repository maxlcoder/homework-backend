package request

import (
	base_request "github.com/maxlcoder/homework-backend/app/request"
)

type RoleStoreRequest struct {
	Name  string                   `json:"name" binding:"required,min=1,max=30" label:"角色名"`
	Menus []base_request.IdRequest `json:"menus" binding:"required,dive" label:"菜单"`
}
