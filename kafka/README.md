# Kafka 功能测试

本目录提供了用于测试项目Kafka功能的工具和示例代码。

## 目录结构

- `kafka_test.go` - 单元测试文件，使用Go的testing框架
- `kafka_test_main.go` - 可执行的测试程序，提供多种测试模式
- `kakfa.go` - Kafka初始化和配置
- `producer.go` - Kafka生产者实现
- `consumer.go` - Kafka消费者实现
- `handler_order.go` - 订单相关消息处理示例

## 测试准备

1. 确保Kafka服务器已启动并运行在默认端口（localhost:9092）
2. 如果使用不同的Kafka服务器，可以修改 `config/config.yaml` 中的Kafka配置

## 测试方法

### 1. 单元测试

使用Go的testing框架运行单元测试：

```bash
cd /Users/woody/workspace/homework-backend
# 运行所有Kafka相关测试
go test -v ./kafka
# 运行特定测试函数
go test -v ./kafka -run TestKafka
```

### 2. 可执行测试程序

使用 `kafka_test_main.go` 可以灵活地测试Kafka功能：

```bash
# 编译测试程序
go build -o kafka_test ./kafka/kafka_test_main.go

# 运行测试程序（默认模式：both，既测试生产者又测试消费者）
./kafka_test

# 只测试生产者
./kafka_test -mode producer

# 只测试消费者
./kafka_test -mode consumer

# 指定Kafka服务器地址
./kafka_test -brokers "kafka1:9092,kafka2:9092"

# 指定测试主题
./kafka_test -topic "my_test_topic"
```

## 测试功能说明

### 生产者测试

- 同步发送消息
- 异步发送消息
- 发送多种格式的消息（目前主要是OrderCreatedMsg格式）

### 消费者测试

- 消费指定主题的消息
- 显示消息的详细信息（主题、分区、偏移量、键、值等）
- 尝试将消息解析为OrderCreatedMsg格式
- 持续运行直到收到中断信号（Ctrl+C）

## 配置说明

Kafka配置位于 `config/config.yaml` 文件中：

```yaml
kafka:
  brokers:       # Kafka服务器地址列表
    - 127.0.0.1:9092
  async: false   # 默认发送模式（true为异步，false为同步）
  topics:        # 应用中使用的主题列表
    - test_topic
    - order_created
```

## 代码说明

### 消息格式

目前主要测试的消息格式是 `OrderCreatedMsg`：

```go
type OrderCreatedMsg struct {
    OrderID string `json:"orderId"`
    UserID  string `json:"userId"`
}
```

### 主要函数

- `InitKafka()` - 从配置初始化Kafka
- `SendSync(topic string, value []byte)` - 同步发送消息
- `SendAsync(topic string, value []byte)` - 异步发送消息
- `StartConsumer(topic, groupID string)` - 启动消费者
- `RegisterHandler(topic string, handler func(kafka.Message) error)` - 注册消息处理函数

## 注意事项

1. 确保Kafka服务器已正确配置并运行
2. 如果遇到连接问题，请检查网络设置和Kafka服务器配置
3. 测试程序默认使用临时的消费者组，每次运行都会从最新或最开始的消息开始消费
4. 可以通过修改代码中的配置和参数来满足不同的测试需求
