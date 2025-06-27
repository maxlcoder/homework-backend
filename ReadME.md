
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
   > vo -> 角色 v1-> object v2-> action  这个顺序和 casbin 配置设置的 `r = sub, obj, act` 对应

#### 路由
 * 路由分组
 * Restful 风格

#### 密码
 * 使用 `bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)` 生成密码，使用 `err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))` 检查密码

#### 入参校验
 * 独立的 Request 
 * 自定义验证

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


