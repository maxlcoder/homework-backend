package request

type AdminCreateRequest struct {
	Name     string `json:"name" binding:"required,min=1,max=30" label:"用户名"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
	Password string `json:"password" binding:"required" label:"密码"`
}
