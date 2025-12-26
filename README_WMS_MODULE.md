# WMS 仓库管理系统模块

## 1. 模块概述

WMS (Warehouse Management System) 仓库管理系统模块是本项目的核心业务模块之一，用于管理仓库的基本信息、货位、库存、拣货任务等核心功能。

### 主要功能
- 仓库基本信息管理
- 货位管理
- 拣货车管理
- 拣货任务管理
- 库存管理
- 入库/出库管理

## 2. 目录结构

```
homework-backend/
├── app/
│   ├── controller/
│   │   └── wms_controller/          # WMS 控制器层
│   │       └── picking_car_controller.go
│   ├── request/
│   │   └── wms_request/             # WMS 请求验证层
│   │       └── picking_car_store_request.go
│   └── route/
│       └── route.go                 # 路由配置
├── model/
│   └── wms.go                       # WMS 数据模型
├── repository/
│   ├── wms_bin_repository.go        # 货位仓库
│   ├── wms_picking_basket_repository.go
│   ├── wms_picking_car_repository.go # 拣货车仓库
│   ├── wms_picking_task_repository.go
│   └── wms_warehouse_repository.go  # 仓库仓库
└── service/
    └── wms_service/
        ├── bin_service.go           # 货位服务
        └── picking_car_service.go   # 拣货车服务
```

## 3. 核心数据模型

WMS 模块的核心数据模型定义在 `model/wms.go` 文件中，主要包括以下实体：

### 3.1 基础模型
- `Warehouse` - 仓库信息
- `Bin` - 货位信息
- `Staff` - 仓库人员

### 3.2 入库相关
- `StockOrder` - 入库订单
- `StockOrderItem` - 入库订单明细
- `StockOperationLog` - 入库操作日志

### 3.3 出库相关
- `PickingCar` - 拣货车
- `PickingBasket` - 拣货篮
- `PickingTask` - 拣货任务
- `PickingTaskBasket` - 拣货任务与拣货篮关联
- `PickingTaskBasketProduct` - 拣货任务商品明细

## 4. 仓库层 (Repository)

仓库层负责数据访问操作，采用泛型基础仓库 `BaseRepository` 实现通用的 CRUD 操作，并在此基础上扩展特定业务方法。

### 4.1 PickingCarRepository
```go
// 拣货车仓库接口
type PickingCarRepository interface {
    // 通用仓库接口
    BaseRepository[model.PickingCar]
    // 特定业务方法
    FindByCode(code string) (*model.PickingCar, error)
}

// 拣货车仓库实现
func NewPickingCarRepository(db *gorm.DB) PickingCarRepository {
    return &pickingCarRepository{
        BaseRepository: repository.NewBaseRepository[model.PickingCar](db),
        db:            db,
    }
}
```

### 4.2 BinRepository
```go
// 货位仓库接口
type BinRepository interface {
    BaseRepository[model.Bin]
    FindByCode(code string) (*model.Bin, error)
    FindByWarehouseID(warehouseID uint) ([]model.Bin, error)
}
```

### 4.3 其他仓库
- `WarehouseRepository` - 仓库信息仓库
- `PickingTaskRepository` - 拣货任务仓库
- `PickingBasketRepository` - 拣货篮仓库

## 5. 服务层 (Service)

服务层负责业务逻辑处理，封装复杂的业务规则，协调多个仓库的操作。

### 5.1 PickingCarService
```go
// 拣货车服务接口
type PickingCarService interface {
    CreatePickingCar(request request.PickingCarStoreRequest) (*model.PickingCar, error)
    GetPickingCar(id uint) (*model.PickingCar, error)
    UpdatePickingCar(id uint, request request.PickingCarUpdateRequest) (*model.PickingCar, error)
    DeletePickingCar(id uint) error
    ListPickingCars(page, pageSize int) ([]model.PickingCar, int64, error)
}

// 拣货车服务实现
func NewPickingCarService(db *gorm.DB, pickingCarRepo repository.PickingCarRepository) *PickingCarService {
    return &PickingCarService{
        db:              db,
        pickingCarRepo:  pickingCarRepo,
    }
}
```

### 5.2 BinService
```go
// 货位服务接口
type BinService interface {
    CreateBin(request request.BinStoreRequest) (*model.Bin, error)
    GetBin(id uint) (*model.Bin, error)
    UpdateBin(id uint, request request.BinUpdateRequest) (*model.Bin, error)
    DeleteBin(id uint) error
    ListBins(warehouseID uint, page, pageSize int) ([]model.Bin, int64, error)
}
```

## 6. 控制器层 (Controller)

控制器层负责处理 HTTP 请求，解析请求参数，调用服务层方法，并返回响应结果。

### 6.1 PickingCarController
```go
// 拣货车控制器
type PickingCarController struct {
    pickingCarService wms_service.PickingCarServiceInterface
}

// 创建拣货车
func (c *PickingCarController) Store(ctx *gin.Context) {
    var request request.PickingCarStoreRequest
    if err := ctx.ShouldBindJSON(&request); err != nil {
        response.Error(ctx, response.BadRequest, "参数错误", err.Error())
        return
    }

    pickingCar, err := c.pickingCarService.CreatePickingCar(request)
    if err != nil {
        response.Error(ctx, response.InternalServerError, "创建失败", err.Error())
        return
    }

    response.Success(ctx, response.Created, "创建成功", pickingCar)
}
```

## 7. 请求验证层 (Request)

请求验证层定义了 API 请求的参数结构和验证规则。

### 7.1 PickingCarStoreRequest
```go
// 拣货车创建请求
type PickingCarStoreRequest struct {
    Code string `json:"code" binding:"required,min=1,max=20"`
}
```

## 8. 模块初始化

WMS 模块通过 `initWmsModule` 函数初始化，按模块组织仓库、服务和控制器的创建与依赖注入。

```go
// WMS 模块初始化
func initWmsModule(db *gorm.DB) *WmsModule {
    // 初始化仓库
    pickingCarRepository := repository.NewPickingCarRepository(db)

    // 初始化服务
    pickingCarService := wms_service.NewPickingCarService(db, pickingCarRepository)

    // 初始化控制器
    pickingCarController := wms_controller.NewPickingCarController(*pickingCarService)
    binController := wms_controller.NewBinController(*pickingCarService)

    return &WmsModule{
        PickcarController: pickingCarController,
        BinController:     binController,
    }
}
```

## 9. 路由注册

WMS 模块的路由需要在 `ApiRoutes` 函数中注册。

```go
func ApiRoutes(r *gin.Engine, enforcer *casbin.Enforcer) {
    // 初始化所有模块
    appModules := initAllModules(database.DB, enforcer)

    // 注册 WMS 模块路由
    registerWmsRoutes(r, appModules.WmsModule)
}

// WMS 路由注册函数
func registerWmsRoutes(r *gin.Engine, wmsModule *WmsModule) {
    wmsApi := r.Group("/api/wms")
    
    // 拣货车路由
    pickingCarGroup := wmsApi.Group("/picking-cars")
    {
        pickingCarGroup.POST("", wmsModule.PickcarController.Store)
        pickingCarGroup.GET("/:id", wmsModule.PickcarController.Show)
        pickingCarGroup.PUT("/:id", wmsModule.PickcarController.Update)
        pickingCarGroup.DELETE("/:id", wmsModule.PickcarController.Destroy)
        pickingCarGroup.GET("", wmsModule.PickcarController.Index)
    }
    
    // 货位路由
    binGroup := wmsApi.Group("/bins")
    {
        binGroup.POST("", wmsModule.BinController.Store)
        binGroup.GET("/:id", wmsModule.BinController.Show)
        binGroup.PUT("/:id", wmsModule.BinController.Update)
        binGroup.DELETE("/:id", wmsModule.BinController.Destroy)
        binGroup.GET("", wmsModule.BinController.Index)
    }
}
```

## 10. 使用示例

### 10.1 创建拣货车

**请求：**
```bash
POST /api/wms/picking-cars
Content-Type: application/json

{
    "code": "PC-001"
}
```

**响应：**
```json
{
    "code": 201,
    "message": "创建成功",
    "data": {
        "id": 1,
        "code": "PC-001",
        "status": 1,
        "created_at": "2023-10-01T10:00:00Z",
        "updated_at": "2023-10-01T10:00:00Z"
    }
}
```

### 10.2 获取拣货车列表

**请求：**
```bash
GET /api/wms/picking-cars?page=1&page_size=10
```

**响应：**
```json
{
    "code": 200,
    "message": "获取成功",
    "data": {
        "list": [
            {
                "id": 1,
                "code": "PC-001",
                "status": 1,
                "created_at": "2023-10-01T10:00:00Z",
                "updated_at": "2023-10-01T10:00:00Z"
            }
        ],
        "total": 1,
        "page": 1,
        "page_size": 10
    }
}
```

## 11. 扩展指南

### 11.1 添加新功能

1. **定义数据模型**：在 `model/wms.go` 中添加新的模型结构体
2. **创建仓库层**：实现对应的 Repository 接口和结构体
3. **实现服务层**：实现业务逻辑
4. **创建控制器**：处理 HTTP 请求
5. **添加请求验证**：定义请求参数结构和验证规则
6. **注册路由**：在 `registerWmsRoutes` 函数中添加新路由

### 11.2 示例：添加库存管理功能

1. **定义库存模型**：
```go
type Inventory struct {
    BaseModel
    SKU        string  `json:"sku" gorm:"index;not null"`
    BinID      uint    `json:"bin_id" gorm:"index;not null"`
    Quantity   int     `json:"quantity" gorm:"default:0"`
    LockedQty  int     `json:"locked_qty" gorm:"default:0"`
    Price      float64 `json:"price" gorm:"type:decimal(10,2);default:0.00"`
    
    // 关联
    Bin        Bin     `json:"bin" gorm:"foreignKey:BinID"`
}
```

2. **创建库存仓库**：
```go
type InventoryRepository interface {
    BaseRepository[model.Inventory]
    FindBySKUAndBinID(sku string, binID uint) (*model.Inventory, error)
    UpdateQuantity(id uint, quantity int) error
}
```

3. **实现库存服务**：
```go
type InventoryService interface {
    StockIn(request request.InventoryStockInRequest) error
    StockOut(request request.InventoryStockOutRequest) error
    GetInventory(sku string, binID uint) (*model.Inventory, error)
}
```

4. **创建库存控制器**：
```go
type InventoryController struct {
    inventoryService wms_service.InventoryService
}
```

5. **注册库存路由**：
```go
func registerWmsRoutes(r *gin.Engine, wmsModule *WmsModule) {
    // ... 现有路由
    
    // 库存路由
    inventoryGroup := wmsApi.Group("/inventory")
    {
        inventoryGroup.POST("/stock-in", wmsModule.InventoryController.StockIn)
        inventoryGroup.POST("/stock-out", wmsModule.InventoryController.StockOut)
        inventoryGroup.GET("", wmsModule.InventoryController.Show)
    }
}
```

## 12. 注意事项

1. **事务处理**：在涉及多个数据操作的业务方法中，使用事务确保数据一致性
2. **参数验证**：所有 API 请求必须进行参数验证，防止无效数据进入系统
3. **错误处理**：统一的错误处理机制，返回标准化的错误响应
4. **权限控制**：结合 Casbin 进行细粒度的权限控制
5. **性能优化**：
   - 合理使用索引
   - 避免 N+1 查询
   - 批量操作优化

## 13. 待实现功能

- [ ] 库存管理功能
- [ ] 入库单管理功能
- [ ] 出库单管理功能
- [ ] 拣货任务分配功能
- [ ] 库存盘点功能
- [ ] 库存预警功能
- [ ] 报表统计功能

## 14. 联系方式

如有问题或建议，请联系开发团队。