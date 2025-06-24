package request

type PageRequest struct {
	Page    int `form:"page" binding:"min=1"`
	PerPage int `form:"per_page" binding:"min=1,max=100"` // 最大支持 100 TODO 配置项
}
