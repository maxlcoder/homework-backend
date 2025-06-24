package response

// 分页
type PageResponse[T any] struct {
	Page    int   `json:"page"`
	PerPage int   `json:"per_page"`
	Total   int64 `json:"total"`
	Data    []T   `json:"data"`
}

// 基础
type BaseResponse struct {
	ID        uint     `json:"id"`
	CreatedAt JSONTime `json:"created_at"`
	UpdatedAt JSONTime `json:"updated_at"`
}

// 一般作为创建对象后返回
type DataId struct {
	ID int `json:"id"`
}
