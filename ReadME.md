# Homework Backend

基于 Gin 框架构建的后台管理系统，提供用户管理、权限控制、菜单管理等功能。

## 📋 目录

- [项目介绍](#项目介绍)
- [技术栈](#技术栈)
- [项目结构](#项目结构)
- [核心功能](#核心功能)
- [权限设计](#权限设计)
- [快速开始](#快速开始)
- [API 文档](#api-文档)
- [开发指南](#开发指南)

## 🚀 项目介绍

这是一个基于 Gin 框架开发的后台管理系统，采用现代化的架构设计，支持多租户、RBAC 权限模型，提供完整的用户管理和权限控制功能。

### 主要特性

- 🏗️ **架构清晰**：采用 MVC 分层架构，代码结构清晰
- 🔐 **权限完善**：基于 RBAC 的权限控制，支持菜单和 API 权限
- 🏢 **多租户**：支持多租户架构设计
- 📝 **参数校验**：完善的请求参数校验和错误处理
- 🌐 **国际化**：支持多语言本地化
- 🚀 **高性能**：基于 Gin 框架，性能优异

## 🛠️ 技术栈

- **Web 框架**：[Gin](https://gin-gonic.com/)
- **ORM**：[GORM](https://gorm.io/)
- **权限控制**：[Casbin](https://casbin.org/)
- **配置管理**：[Viper](https://github.com/spf13/viper)
- **JWT 认证**：[gin-jwt](https://github.com/appleboy/gin-jwt)
- **密码加密**：bcrypt
- **参数校验**：validator v10

## 📁 项目结构

```
.
├── app/                    # 应用层
│   ├── contract/          # 接口定义
│   ├── controller/        # 控制器
│   ├── middleware/        # 中间件
│   ├── request/          # 请求结构体
│   ├── response/         # 响应结构体
│   └── route/            # 路由定义
├── cmd/                   # 命令行工具
├── config/               # 配置文件
├── database/             # 数据库相关
├── model/                # 数据模型
├── pkg/                  # 公共包
├── repository/           # 数据访问层
├── service/              # 业务逻辑层
└── main.go              # 程序入口
```

### 架构说明

- **前台端**：`api` - 面向普通用户的接口
- **后台端**：`admin` - 面向管理员的接口
- **扩展端**：支持其他业务端扩展

## ⚡ 核心功能

### 配置

- 系统使用 etcd 作为远程配置中心，部分非敏感信息且几乎不太会调整配置放在代码仓库的 yaml 文件件中
- 使用 Viper 进行配置获取
- 系统代码保留配置更新监听（监听和轮询机制）代码，考虑项目部署的简便性，暂不考虑开启配置监听，而是直接重启项目形式进行配置更新

### 🗄️ ORM 数据层

使用 GORM 作为 ORM 框架，提供：
- 数据库连接管理
- 模型定义和关联
- 查询构建器
- 事务支持

### 🔒 权限校验

采用 Casbin 实现 RBAC 权限模型：
- **权限模型**：`subject(用户/角色) -> object(资源) -> action(操作)`
- **配置格式**：对应 Casbin 的 `r = sub, obj, act` 配置
- **双重保障**：数据库存储 + Casbin 校验

### 👥 权限分配

完善的权限管理体系：

#### 核心表设计
- `menu`：菜单表
- `menu_permission`：菜单权限关联表  
- `permission`：权限表
- `role`：角色表
- `role_menu`：角色菜单关联表
- `role_permission`：角色权限关联表
- `admin_role`：管理员角色关联表

#### 权限层级
- **显示权限**：控制菜单、按钮等 UI 元素的可见性
- **操作权限**：控制 API 接口的访问权限
- **多角色支持**：一个管理员可拥有多个角色

### 🌐 全局变量管理

- **请求级变量**：使用 `c.Set()` 设置请求上下文变量
- **用户缓存**：避免重复查询数据库获取用户信息

### 🛣️ 路由管理

- **分组路由**：按功能模块分组管理
- **RESTful API**：遵循 REST 规范设计接口
- **中间件支持**：认证、权限、错误处理等中间件

### 🔐 密码安全

使用 bcrypt 进行密码加密：

```go
// 密码加密
bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

// 密码验证
err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
```

### ✅ 参数校验

完善的请求参数校验机制：

#### 特性
- **独立的 Request 结构**：每个接口对应独立的请求结构体
- **自定义验证规则**：支持业务相关的验证逻辑
- **默认值设置**：自动设置参数默认值
- **错误信息本地化**：支持中文错误提示

#### 使用示例

```go
// 1. 定义请求结构体
type PageRequest struct {
    Page     *int   `json:"page" default:"1" binding:"omitempty,gte=1"`
    PageSize *int   `json:"page_size" default:"10" binding:"omitempty,gte=1,max=100"`
    Name     string `json:"name" default:"guest" binding:"omitempty,min=2"`
}

// 2. 通用绑定和验证函数
func BindAndSetDefaults(c *gin.Context, req interface{}) error {
    // 根据 Content-Type 选择绑定方式
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
    
    // 应用默认值
    if err := defaults.Set(req); err != nil {
        return err
    }
    return nil
}

// 3. 在控制器中使用
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
        return
    }
    
    pageResponse := response.BuildPageResponse[model.Admin, *response.AdminResponse](
        admins, total, pagination.Page, pagination.PerPage, response.NewAdminResponse)
    controller.Success(c, pageResponse)
}
```

#### 重要说明

> **ShouldBind 机制**：
> - `Content-Type=application/json` → 调用 `ShouldBindJSON`
> - `Content-Type=application/x-www-form-urlencoded` → 调用 `ShouldBindForm`
>
> 为了兼容不同的请求类型，在参数校验中通过判断请求方式来选择合适的绑定方法。

### 🎫 JWT 认证中间件

基于 [gin-jwt](https://github.com/appleboy/gin-jwt) 实现的认证系统：

- **多用户类型**：支持 `user`、`admin`、`jsc` 等多种用户类型登录
- **类型识别**：登录时指定 `user_type`，token 解析时验证用户类型
- **身份处理**：通过 `identityHandler` 指定用户类型

### 📤 统一响应格式

所有 API 接口采用统一的响应格式，确保前端处理的一致性。

### 🌍 国际化支持

- **参数校验本地化**：错误提示信息支持中文显示
- **多语言扩展**：支持其他语言的本地化扩展

### 💾 缓存机制

提供多层次的缓存支持，提升系统性能。

### 🗃️ 数据库管理

完善的数据库管理功能，包括连接池、事务管理等。

## 🔐 权限设计

### 设计要求

- ✅ **多租户支持**：支持多租户架构
- ✅ **RBAC 模式**：基于角色的权限控制（兼容弱 ABAC 模式）
- ✅ **权限分层**：区分显示层（菜单按钮等）和操作层（API）权限控制

### 设计原理

1. **数据存储**：设计相关表存储基础角色、菜单、权限相关的内容
2. **权限校验**：使用 Casbin 作为后端权限的校验方式
3. **超级管理员**：设计超管拥有全部权限，能够给租户分配租户可操作权限
4. **权限初始化**：每次升级将权限的初次分配用升级脚本操作，避免人工操作

### 数据初始化

1. **超级管理员**：初始化超级管理员账号和角色
2. **权限分配**：为超管分配所有权限
3. **基础数据**：初始化基础菜单和权限数据

## 🚀 快速开始

### 环境要求

- Go 1.19+
- MySQL 5.7+
- Redis 6.0+

### 安装步骤

1. **克隆项目**
```bash
git clone <repository-url>
cd homework-backend
```

2. **安装依赖**
```bash
go mod download
```

3. **配置文件**
```bash
cp config/config.yaml.example config/config.yaml
# 编辑配置文件，设置数据库连接等信息
```

4. **数据库初始化**
```bash
go run cmd/main.go migrate
go run cmd/main.go seed
```

5. **启动服务**
```bash
go run main.go
```


## 🔧 开发指南

### 代码规范

- 遵循 Go 官方代码规范
- 使用 gofmt 格式化代码
- 添加必要的注释和文档

### 业务

- restful api 入口对应 controller
- controller 注入 service，service 注入 repository ，repository 通过 gorm 处理 model 对应的数据
- 三层业务架构方案，controller 不直接和 Repository 接触，均通过 service 进行中间层的数据整合
- 定义通用的 repository 操作方案
- 

### 提交规范

- 遵循 Conventional Commits 规范
- 提交前运行测试
- 确保代码质量

### 测试

```bash
# 运行测试
go test ./...

# 运行测试并生成覆盖率报告
go test -cover ./...
```

---

## 📄 许可证

本项目采用 MIT 许可证，详情请查看 [LICENSE](LICENSE) 文件。

## 🤝 贡献

欢迎提交 Issue 和 Pull Request 来帮助改进这个项目。

## 📞 联系方式

如有问题或建议，请通过以下方式联系：

- 提交 Issue
- 发送邮件到：[liurenlin77@gmail.com]

