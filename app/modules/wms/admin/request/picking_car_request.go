package request

type PickingCarStoreRequest struct {
	Code string `json:"code" binding:"required,min=1,max=60" label:"编号"`
}

type PickingCarUpdateRequest struct {
	ID   uint   `json:"id" label:"ID"`
	Code string `json:"code" binding:"required,min=1,max=60" label:"编号"`
}
