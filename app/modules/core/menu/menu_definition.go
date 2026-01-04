package menu

import (
	"github.com/maxlcoder/homework-backend/app/contract"
	"github.com/maxlcoder/homework-backend/model"
)

// CoreMenuDefinition 核心模块菜单定义实现
type CoreMenuDefinition struct{}

// NewCoreMenuDefinition 创建核心模块菜单定义实例
func NewCoreMenuDefinition() contract.MenuDefinition {
	return &CoreMenuDefinition{}
}

// GetMenus 返回核心模块的菜单定义
func (d *CoreMenuDefinition) GetMenus() []model.Menu {
	return []model.Menu{
		{
			Number: "system-setting",
			Name:   "系统设置",
			Sort:   1,
			Children: []*model.Menu{
				{
					Number: "admin-management",
					Name:   "账号管理",
					Children: []*model.Menu{
						{
							Number: "admin-list",
							Name:   "列表",
							Permissions: []*model.Permission{
								{
									Name:   "列表",
									PATH:   "/admin/admins",
									Method: "GET",
								},
							},
						},
						{
							Number: "admin-add",
							Name:   "新增",
							Permissions: []*model.Permission{
								{
									Name:   "新增",
									PATH:   "/admin/admins",
									Method: "POST",
								},
							},
						},
						{
							Number: "admin-update",
							Name:   "更新",
							Permissions: []*model.Permission{
								{
									Name:   "更新",
									PATH:   "/admin/admins/:id",
									Method: "PUT",
								},
							},
						},
						{
							Number: "admin-detail",
							Name:   "详情",
							Permissions: []*model.Permission{
								{
									Name:   "详情",
									PATH:   "/admin/admins/:id",
									Method: "GET",
								},
							},
						},
						{
							Number: "admin-delete",
							Name:   "删除",
							Permissions: []*model.Permission{
								{
									Name:   "删除",
									PATH:   "/admin/admins/:id",
									Method: "DELETE",
								},
							},
						},
					},
				},
				{
					Number: "role-management",
					Name:   "角色管理",
					Children: []*model.Menu{
						{
							Number: "role-list",
							Name:   "列表",
							Permissions: []*model.Permission{
								{
									Name:   "列表",
									PATH:   "/admin/roles",
									Method: "GET",
								},
							},
						},
						{
							Number: "role-add",
							Name:   "新增",
							Permissions: []*model.Permission{
								{
									Name:   "新增",
									PATH:   "/admin/roles",
									Method: "POST",
								},
							},
						},
						{
							Number: "role-update",
							Name:   "更新",
							Permissions: []*model.Permission{
								{
									Name:   "更新",
									PATH:   "/admin/roles/:id",
									Method: "PUT",
								},
							},
						},
						{
							Number: "role-detail",
							Name:   "详情",
							Permissions: []*model.Permission{
								{
									Name:   "详情",
									PATH:   "/admin/roles/:id",
									Method: "GET",
								},
							},
						},
						{
							Number: "role-delete",
							Name:   "删除",
							Permissions: []*model.Permission{
								{
									Name:   "删除",
									PATH:   "/admin/roles/:id",
									Method: "DELETE",
								},
							},
						},
					},
				},
			},
		},
	}
}
