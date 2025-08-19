package response

type RoleResponse struct {
	BaseResponse
	Name string `json:"name"`
}
