package wms_admin_response

// AdminBinResponse 管理员专用库位响应结构
type AdminBinResponse struct {
    ID        uint   `json:"id"`
    Code      string `json:"code"`
    Name      string `json:"name"`
    Status    int    `json:"status"`
    CreatedBy uint   `json:"created_by"`
    UpdatedBy uint   `json:"updated_by"`
    CreatedAt string `json:"created_at"`
    UpdatedAt string `json:"updated_at"`
}

// AdminPickingCarResponse 管理员专用拣货车响应结构
type AdminPickingCarResponse struct {
    ID        uint   `json:"id"`
    Code      string `json:"code"`
    Status    int    `json:"status"`
    CurrentUserID uint `json:"current_user_id"`
    CreatedBy uint   `json:"created_by"`
    UpdatedBy uint   `json:"updated_by"`
    CreatedAt string `json:"created_at"`
    UpdatedAt string `json:"updated_at"`
}

// AdminStaffResponse 管理员专用员工响应结构
type AdminStaffResponse struct {
    ID        uint   `json:"id"`
    Name      string `json:"name"`
    Code      string `json:"code"`
    Position  string `json:"position"`
    Status    int    `json:"status"`
    CreatedBy uint   `json:"created_by"`
    UpdatedBy uint   `json:"updated_by"`
    CreatedAt string `json:"created_at"`
    UpdatedAt string `json:"updated_at"`
}

// AdminPickingBasketResponse 管理员专用拣货框响应结构
type AdminPickingBasketResponse struct {
    ID        uint   `json:"id"`
    Code      string `json:"code"`
    Status    int    `json:"status"`
    CreatedBy uint   `json:"created_by"`
    UpdatedBy uint   `json:"updated_by"`
    CreatedAt string `json:"created_at"`
    UpdatedAt string `json:"updated_at"`
}