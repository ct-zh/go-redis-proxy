# 日志中间件使用指南

## 概述

本项目实现了完整的日志系统，包括分级日志、配置化存储和请求中间件。

## 功能特性

### 1. 分级日志
- **Access日志** (`logs/access.log`): 记录所有HTTP请求的访问信息
- **Error日志** (`logs/error.log`): 记录错误信息和5xx状态码请求
- **Info日志** (`logs/info.log`): 记录一般信息，如系统启动等
- **Debug日志** (`logs/debug.log`): 记录详细的调试信息

### 2. 配置化存储
日志存储位置可通过环境变量配置：
```bash
export LOG_LEVEL=debug        # 日志级别: debug, info, warn, error
export LOG_DIR=logs          # 日志目录，默认为项目根目录下的logs文件夹
export LOG_MAX_SIZE=100      # 单个日志文件最大大小(MB)
export LOG_MAX_AGE=30        # 日志文件保留天数
export LOG_COMPRESS=true     # 是否压缩旧日志文件
```

### 3. 中间件设计

#### 3.1 日志中间件 (`LoggingMiddleware`)
自动记录每个HTTP请求的详细信息：

**记录内容：**
- 请求方法、路径、查询参数
- 客户端IP、User-Agent
- 请求体和响应体（JSON格式）
- 请求耗时（毫秒）
- 状态码
- 请求和响应大小

**使用方法：**
```go
// 在路由设置中添加中间件
engine.Use(middleware.LoggingMiddleware())
```

#### 3.2 请求ID中间件 (`RequestIDMiddleware`)
为每个请求生成唯一ID，便于追踪：

```go
engine.Use(middleware.RequestIDMiddleware())
```

#### 3.3 恢复中间件 (`RecoveryMiddleware`)
捕获panic并记录错误日志：

```go
engine.Use(middleware.RecoveryMiddleware())
```

### 4. 完整的中间件配置示例

```go
func SetupWithContainer(engine *gin.Engine, container *container.Container) {
    // 添加全局中间件（顺序很重要）
    engine.Use(middleware.RequestIDMiddleware())  // 1. 请求ID中间件
    engine.Use(middleware.RecoveryMiddleware())   // 2. 恢复中间件
    engine.Use(middleware.LoggingMiddleware())    // 3. 日志中间件

    // 设置路由...
}
```

### 5. 日志格式示例

#### Access日志示例：
```json
{
  "client_ip": "::1",
  "duration_ms": 70,
  "level": "info",
  "message": "HTTP Request",
  "method": "POST",
  "path": "/api/v1/redis/string/set",
  "query": "",
  "request_body": {
    "key": "test",
    "value": "hello world"
  },
  "request_size": 36,
  "response_body": {
    "code": 2000,
    "data": null,
    "message": "Redis连接失败"
  },
  "response_size": 55,
  "status_code": 200,
  "timestamp": "2025-07-29 10:06:55",
  "type": "access",
  "user_agent": "curl/8.7.1"
}
```

#### Debug日志示例：
```json
{
  "headers": {
    "Accept": ["*/*"],
    "Content-Length": ["36"],
    "Content-Type": ["application/json"],
    "User-Agent": ["curl/8.7.1"]
  },
  "level": "debug",
  "message": "Request Details",
  "params": null,
  "timestamp": "2025-07-29 10:06:55",
  "type": "debug"
}
```

### 6. 在代码中使用日志

#### 6.1 使用全局日志函数
```go
import "github.com/ct-zh/go-redis-proxy/pkg/logger"

// 记录信息日志
logger.Info("用户登录成功", logrus.Fields{
    "user_id": 123,
    "ip": "192.168.1.1",
})

// 记录错误日志
logger.Error("数据库连接失败", logrus.Fields{
    "error": err.Error(),
    "database": "redis",
})

// 格式化日志
logger.Infof("处理请求耗时: %dms", duration)
```

#### 6.2 获取日志器实例
```go
logger := logger.GetLogger()
logger.Access("API调用", logrus.Fields{
    "endpoint": "/api/v1/users",
    "method": "GET",
})
```

### 7. 日志级别说明

- **Debug**: 详细的调试信息，包括请求headers、参数等
- **Info**: 一般信息，如系统启动、正常的API调用
- **Warn**: 警告信息，如4xx状态码的请求
- **Error**: 错误信息，如5xx状态码的请求、系统异常

### 8. 最佳实践

1. **中间件顺序**: RequestID → Recovery → Logging，确保每个请求都有ID且异常能被捕获
2. **敏感信息**: 避免在日志中记录密码、token等敏感信息
3. **日志轮转**: 生产环境建议配置日志轮转，避免单个文件过大
4. **性能考虑**: 大量请求时，可以考虑异步写入日志

### 9. 环境变量配置示例

```bash
# 开发环境
export LOG_LEVEL=debug
export LOG_DIR=logs

# 生产环境
export LOG_LEVEL=info
export LOG_DIR=/var/log/go-redis-proxy
export LOG_MAX_SIZE=500
export LOG_MAX_AGE=7
export LOG_COMPRESS=true
```

## 总结

通过这套日志系统，你可以：
1. **只需添加中间件**即可为所有请求自动记录日志
2. **分级管理**不同类型的日志信息
3. **配置化存储**，适应不同环境需求
4. **结构化日志**，便于后续分析和监控