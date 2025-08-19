package seed

import (
	"errors"
	"fmt"
	"strings"
	"sync"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/maxlcoder/homework-backend/model"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func InitSeed(db *gorm.DB, r *gin.Engine, enforcer *casbin.Enforcer) error {

	// 添加超管
	if err := seedSuperAdmin(db); err != nil {
		return err
	}
	// 添加超管角色
	if err := seedSuperRole(db); err != nil {
		return err
	}
	// 添加超管与超管角色关联
	if err := seedSuperAdminRole(db); err != nil {
		return err
	}
	// 权限初始化，将全部 /admin 开头的路由转移到 permission 表
	if err := seedPermissions(db, r); err != nil {
		return err
	}
	// 菜单初始化，将菜单权限数组中的内容同步到 menu 和 menu_permission 表
	if err := seedMenus(db); err != nil {
		return err
	}

	// 将超管角色与菜单权限等关联，并补充 casbin 记录
	if err := seedRoleMenuPermissions(db, enforcer); err != nil {
		return err
	}

	return nil
}

// 超管
func seedSuperAdmin(db *gorm.DB) error {
	// 检查相关表是否存在
	has := db.Migrator().HasTable(&model.Admin{})
	if !has {
		return fmt.Errorf("%s table not found", "admin")
	}
	// 检查是否存在
	var admin model.Admin
	err := db.Where("id = ?", 1).First(&admin).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		admin.ID = 1
		admin.Name = "admin"
		admin.Email = "admin@homework.com"
		defaultPassword := viper.GetString("default_password")
		password, err := model.HashPassword(defaultPassword)
		if err != nil {
			return err
		}
		admin.Password = password
		db.Create(&admin)
	}
	return nil
}

// 角色
func seedSuperRole(db *gorm.DB) error {
	// 检查相关表是否存在
	has := db.Migrator().HasTable(&model.Role{})
	if !has {
		return fmt.Errorf("%s table not found", "role")
	}
	// 检查是否存在
	var role model.Role
	err := db.Where("id = ?", 1).First(&role).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		role.ID = 1
		role.Name = "super_admin"
		db.Create(&role)
	}
	return nil
}

// 超管角色分配
func seedSuperAdminRole(db *gorm.DB) error {
	db.Clauses(clause.OnConflict{
		UpdateAll: false,
	}).Create(&model.AdminRole{
		AdminId: 1,
		RoleId:  1,
	})
	return nil
}

// 路由权限
func seedPermissions(db *gorm.DB, r *gin.Engine) error {
	// 检查相关表是否存在
	has := db.Migrator().HasTable(&model.Permission{})
	if !has {
		return fmt.Errorf("%s table not found", "permission")
	}

	// 现有权限数组
	var permissionIds []uint
	// 检查是否存在
	for _, route := range r.Routes() {
		// 是否是后台接口
		if !strings.HasPrefix(route.Path, "/admin/") {
			continue
		}
		var permission model.Permission
		err := db.Where("path = ?", route.Path).Where("method = ?", route.Method).First(&permission).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			permission.PATH = route.Path
			permission.Method = route.Method
			db.Create(&permission)
		}
		permissionIds = append(permissionIds, permission.ID)
	}

	// 删除非现有权限相关关联，默认不存在一个权限没有的情况
	if len(permissionIds) > 0 {
		// 删除权限
		db.Where("id NOT IN ?", permissionIds).Delete(&model.Permission{})
		// 删除菜单权限关联
		db.Where("permission_id NOT IN ?", permissionIds).Delete(&model.MenuPermission{})
		// 删除角色权限关联
		db.Where("permission_id NOT IN ?", permissionIds).Delete(&model.RolePermission{})
	}

	return nil
}

// 菜单
func seedMenus(db *gorm.DB) error {
	// 检查相关表是否存在
	has := db.Migrator().HasTable(&model.Menu{})
	if !has {
		return fmt.Errorf("%s table not found", "menu")
	}

	// 加载菜单
	menus := loadMenus()

	// 存在的菜单列表，删除不存在的
	menuIds := []uint{}
	menuIdCh := make(chan uint)
	var wg sync.WaitGroup

	for _, menu := range menus {
		wg.Add(1)
		go insertUpdateMenu(db, menuIdCh, &menu, 0, &wg)
	}

	go func() {
		wg.Wait()
		close(menuIdCh)
	}()

	menuId := <-menuIdCh
	menuIds = append(menuIds, menuId)
	for menuId := range menuIdCh {
		menuIds = append(menuIds, menuId)
	}

	// 删除不在系统中的菜单
	db.Where("menu_id NOT IN ?", menuIds).Delete(&model.RoleMenu{})
	db.Where("menu_id NOT IN ?", menuIds).Delete(&model.MenuPermission{})

	return nil
}

func insertUpdateMenu(db *gorm.DB, ch chan<- uint, menu *model.Menu, parentId uint, wg *sync.WaitGroup) {
	defer wg.Done()
	// 根据编号查询是否存在，存在则更新，不存在则插入
	var findMenu model.Menu
	err := db.Where("number = ?", menu.Number).First(&findMenu).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		findMenu.Number = menu.Number
		findMenu.Name = menu.Name
		findMenu.ParentID = parentId
		db.Create(&findMenu)
	} else {
		db.Model(&findMenu).Updates(model.Menu{
			Name:     menu.Name,
			ParentID: parentId,
		})
	}
	// 检查是否有 permissions，有则需要写入权限关系
	for _, permision := range menu.Permissions {
		var findPermission model.Permission
		err := db.Where("path = ?", permision.PATH).Where("method = ?", permision.Method).First(&findPermission).Error
		if err == nil {
			if strings.TrimSpace(permision.Name) != "" && findPermission.Name != permision.Name {
				db.Model(&findPermission).Updates(model.Permission{
					Name: permision.Name,
				})
			}
			// 菜单与权限的关联，不存在则创建
			db.Clauses(clause.OnConflict{
				UpdateAll: false,
			}).Create(&model.MenuPermission{
				MenuID:       findMenu.ID,
				PermissionID: findPermission.ID,
			})
		}
	}

	ch <- findMenu.ID

	for _, child := range menu.Children {
		wg.Add(1)
		go insertUpdateMenu(db, ch, child, findMenu.ID, wg)
	}

}

// 角色菜单,权限关联
func seedRoleMenuPermissions(db *gorm.DB, enforcer *casbin.Enforcer) error {
	// 检查相关表是否存在
	has := db.Migrator().HasTable(&model.RoleMenu{})
	if !has {
		return fmt.Errorf("%s table not found", "role_menu")
	}
	// 超管角色与全部菜单关联
	var menus []model.Menu
	result := db.Find(&menus)
	if result.Error != nil {
		return fmt.Errorf("find menu error: %v", result.Error)
	}
	for _, menu := range menus {
		var roleMenu model.RoleMenu
		err := db.Where("id = ?", menu.ID).First(&roleMenu).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			roleMenu.RoleID = 1
			roleMenu.MenuID = menu.ID
			db.Create(&roleMenu)
		}

		// casbin police 添加

	}

	// 超管角色与全部权限关联
	var permissions []model.Permission
	result = db.Find(&permissions)
	if result.Error != nil {
		return fmt.Errorf("find permission error: %v", result.Error)
	}
	for _, permission := range permissions {
		var rolePermission model.RolePermission
		err := db.Where("id = ?", permission.ID).First(&rolePermission).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			rolePermission.RoleID = 1
			rolePermission.PermissionID = permission.ID
			db.Create(&rolePermission)
		}
	}

	return nil
}

func loadMenus() []model.Menu {
	menus := []model.Menu{
		{
			Number: "n1",
			Name:   "菜单 1",
			Sort:   1,
			Children: []*model.Menu{
				{
					Number: "n1-1",
					Name:   "菜单 1-1",
					Children: []*model.Menu{
						{
							Number: "n1-1-1",
							Name:   "菜单 1-1-1",
							Permissions: []*model.Permission{
								{
									Name:   "用户列表",
									PATH:   "/admin/users",
									Method: "GET",
								},
							},
						},
					},
				},
			},
		},
		{
			Number: "n2",
			Name:   "菜单 2",
			Sort:   1,
			Children: []*model.Menu{
				{
					Number: "n2-1",
					Name:   "菜单 2-1",
				},
			},
		},
		{
			Number: "n3",
			Name:   "菜单 3",
			Sort:   1,
			Children: []*model.Menu{
				{
					Number: "n3-1",
					Name:   "菜单 3-1",
				},
			},
		},
	}
	return menus
}
