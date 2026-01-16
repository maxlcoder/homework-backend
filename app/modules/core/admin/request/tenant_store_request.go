package request

type TenantStoreRequest struct {
	Name string `json:"name" binding:"required,min=1,max=60" label:"租户名称"`
}
