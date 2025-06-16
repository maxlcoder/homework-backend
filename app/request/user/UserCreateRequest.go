package request

type UserCreateRequest struct {
	Name  string `json:"name" binding:"required,max=30" label:"用户名"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}
