package response

import (
	"github.com/jinzhu/copier"
	"github.com/maxlcoder/homework-backend/app/modules/core/model"
	"github.com/maxlcoder/homework-backend/app/response"
	"github.com/samber/lo"
)

type AdminResponse struct {
	response.BaseResponse
	Name  string         `json:"name"`
	Email string         `json:"email"`
	Age   uint8          `json:"age"`
	Roles []RoleResponse `json:"roles"`
}

func NewAdminResponse() *AdminResponse {
	return &AdminResponse{}
}

func (r *AdminResponse) FromModel(m model.Admin) {
	copier.Copy(&r.BaseResponse, &m)
	copier.Copy(r, &m)
}

type MeResponse struct {
	response.BaseResponse
	Name  string          `json:"name"`
	Email string          `json:"email"`
	Age   uint8           `json:"age"`
	Roles []RoleResponse  `json:"roles"`
	Menus []*MenuResponse `json:"menus"`
}

// 转换单个 tree
func TreeToResponse(menu *model.Menu) *MenuResponse {
	if menu == nil {
		return nil
	}
	return &MenuResponse{
		BaseResponse: response.BaseResponse{
			ID:        menu.ID,
			CreatedAt: response.JSONTime(menu.CreatedAt),
			UpdatedAt: response.JSONTime(menu.UpdatedAt),
		},
		Name:       menu.Name,
		ParentID:   menu.ParentID,
		Sort:       menu.Sort,
		IsDisabled: menu.IsDisabled,
		Number:     menu.Number,
		Children:   convertChildren(menu.Children),
	}
}

func convertChildren(children []*model.Menu) []*MenuResponse {
	return lo.FilterMap(children, func(child *model.Menu, _ int) (*MenuResponse, bool) {
		resp := TreeToResponse(child)
		return resp, resp != nil
	})
}

// 转换全树
func TreesToResponse(menus []*model.Menu) []*MenuResponse {
	return lo.FilterMap(menus, func(menu *model.Menu, _ int) (*MenuResponse, bool) {
		resp := TreeToResponse(menu)
		return resp, resp != nil
	})
}

type MenuResponse struct {
	response.BaseResponse
	Name       string          `json:"name"`
	ParentID   uint            `json:"parent_id"`
	Sort       int             `json:"sort"`
	IsDisabled bool            `json:"is_disabled"`
	Number     string          `json:"number"`
	Children   []*MenuResponse `json:"children"`
}
