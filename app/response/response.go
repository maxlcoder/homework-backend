package response

type Mapper[M any] interface {
	FromModel(M)
}

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

// 组装分页结果
func BuildPageResponse[M any, R Mapper[M]](models []M, total int64, page, perPage int, newFun func() R) PageResponse[R] {
	responses := make([]R, len(models))
	for i, m := range models {
		r := newFun() // XxxResponse 指针初始化
		r.FromModel(m)
		responses[i] = r
	}
	return PageResponse[R]{
		Page:    page,
		PerPage: perPage,
		Total:   total,
		Data:    responses,
	}
}
