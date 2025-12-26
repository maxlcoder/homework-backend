package request

type RoleUpdateRequest struct {
	Name  string      `json:"name" binding:"required,min=1,max=30" label:"角色名"`
	Menus []IdRequest `json:"menus" binding:"required,dive" label:"菜单"`
}
