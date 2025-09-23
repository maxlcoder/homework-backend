package request

type AdminUpdateRequest struct {
	Name     string      `json:"name" binding:"required,min=1,max=30" label:"用户名"`
	Password string      `json:"password" binding:"" label:"密码"`
	Roles    []IdRequest `json:"roles" binding:"required,dive" label:"角色"`
}
