
# 框架介绍
### 背景

使用 Gin 框架构建一个 web 服务项目


### 组件

#### 配置项管理

使用 [Viper](https://github.com/spf13/viper) 进行配置管理

### 项目结构
 * 前台 `api`
 * 后台 `admin`
 * 其他端

### 功能
 * ORM
   > 使用 gorm
 * 权限校验 
   > 使用 casbin 作为 rbac 形式的权限校验，同时补充 menu，permission 相关表来做菜单业务权限控制
   > vo -> 角色 v1-> object v2-> action  这个顺序和 casbin 配置设置的 `r = sub, obj, act` 对应。
   
 * 权限分配
   >设计 `menu` ，`menu_permission` ，`permission` 表进行菜单和权限的管理
   > 
   >设计 `role` ， `role_menu` ， `role_permission` 表进行角色菜单和角色权限的分配，**其中角色在菜单和权限的对应关系表现上分别对应可视界面和接口 API 的允许范围**
   > 同时针对权限的校验通过 casbin 进行判断，尽管 casbin 也可以作为权限关系的处理服务，但是考虑到业务组件的可替代性，暂时同时保持数据库和 casbin 双重可处理形式
   > 
   >设计 `admin_role` 进行管理员角色分配，支持一个管理员多种角色场景


#### 全局变量

* 请求临时中间的全局变量统一采用 `c.Set` 进行设置，每一个请求单独设置  

- [ ] 用户全局变量缓存，避免每个请求重复请求数据库来获取用户信息
- [ ] -

#### 路由
 * 路由分组
 * Restful 风格

#### 密码
 * 使用 `bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)` 生成密码，使用 `err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))` 检查密码

#### 入参校验
 * 独立的 Request 
 * 自定义验证
 * 入参默认值设置
   
    **---------注意---------**
    >ShouldBind:   
     如果 Content-Type=application/json → 调用 ShouldBindJSON 
     如果 Content-Type=application/x-www-form-urlencoded → 调用 ShouldBindForm
 
    >为此在入参校验中通过判断请求方式来兼容请求类型

    ```aiignore
    # 1. 默认值标签设置
    type PageRequest struct {
      Page     *int   `json:"page" default:"1" binding:"omitempty,gte=1"`
      PageSize *int   `json:"page_size" default:"10" binding:"omitempty,gte=1,max=100"`
      Name     string `json:"name" default:"guest" binding:"omitempty,min=2"`
    }
     
    # 2. 使用定义全局处理默认值函数，同时兼容参数 validator 判断，并对判断结果进行中文转换
    func BindAndSetDefaults(c *gin.Context, req interface{}) error {
        # // 判断一下请求方式
        ct := c.ContentType()
        if ct == "application/json" {
            if c.Request.Method == "GET" {
                if err := c.ShouldBindQuery(req); err != nil {
                    errorTrans := validator.TranslateError(err)
                    return fmt.Errorf("%s", strings.Join(errorTrans, ","))
    
                }
            }
        } else {
            if err := c.ShouldBind(req); err != nil {
                errorTrans := validator.TranslateError(err)
                return fmt.Errorf("%s", strings.Join(errorTrans, ","))
            }
        }
    
        # // 应用默认值（用 creasty/defaults 或你自己写的 applyDefaults）
        if err := defaults.Set(req); err != nil {
            return err
        }
        return nil
    }
    # 3. 在方法中调用
    func (controller *AdminController) Page(c *gin.Context) {
       var pagination model.Pagination
       var filter model.AdminFilter
        
       if err := request.BindAndSetDefaults(c, &pagination); err != nil {
           controller.Error(c, http.StatusBadRequest, err.Error())
           return
       }
        
       _ = c.ShouldBindJSON(&filter)
       total, admins, err := controller.adminService.GetPageByFilter(filter, pagination)
           if err != nil {
           controller.Error(c, http.StatusBadRequest, err.Error())
       }
        
       pageResponse := response.BuildPageResponse[model.Admin, *response.AdminResponse](admins, total, pagination.Page, pagination.PerPage, response.NewAdminResponse)
       controller.Success(c, pageResponse)
        
    }

    ```

#### 验证中间件
 * [gin-jwt](https://github.com/appleboy/gin-jwt) JWT 验证中间件，改造中间支持多用户类型登录（user,admin,jsc），登录时增加 user_type ，token 解析即判断是否是对应的类型。
通过 identityHandler 指定 user_type

#### 控制器
 * 返回格式统一

#### 本地化
 * 参数校验提示本地化
 * 

#### 缓存


#### 数据库


### 菜单权限设计

#### 设计要求
* 支持多租户
* RBAC模式权限（兼容弱ABAC模式）
* 区分显示层（菜单按钮等）和操作层（API） 等权限控制

#### 设计原理

1. 设计相关表存储基础角色、菜单、权限相关的内容
2. 使用 casbin 作为后端权限的校验方式
3. 设计超管拥有全部权限，能够给租户分配租户可操作权限
4. 权限数据初始化设计，每次升级将权限的初次分配用升级脚本操作，避免人工操作

### 数据初始化

1. 初始化超级管理员账号，初始化超级管理员角色
2. 分配超管角色权限


