# Go Redis Proxy

一个基于Go语言的Redis HTTP代理服务，通过RESTful API使Redis支持HTTP请求，并提供丰富的扩展功能。

## 项目概述

Go Redis Proxy是一个高性能的HTTP到Redis的代理服务，允许客户端通过HTTP协议与Redis进行交互，同时提供额外的功能如请求统计、缓存策略、访问控制等。

## 核心功能

### 基础功能
- [ ] HTTP到Redis命令的映射和转换
- [ ] 支持Redis常用数据类型操作（String、Hash、List、Set、ZSet）
- [ ] RESTful API设计
- [ ] JSON格式的请求响应
- [ ] 连接池管理
- [ ] 基本的错误处理和日志记录

### 扩展功能
- [ ] 请求统计和监控
- [ ] 访问频率限制（Rate Limiting）
- [ ] 请求缓存策略
- [ ] 多Redis实例支持（主从、集群）
- [ ] 用户认证和权限控制
- [ ] 请求路由和负载均衡
- [ ] 配置热重载
- [ ] 健康检查接口

### 高级功能
- [ ] Redis命令批量执行
- [ ] 事务支持
- [ ] 发布订阅模式支持
- [ ] 数据备份和恢复
- [ ] 性能指标收集（Prometheus集成）
- [ ] Docker部署支持

## 项目结构

```
go-redis-proxy/
├── cmd/                    # 应用程序入口
│   └── server/
│       └── main.go
├── internal/               # 内部包，不对外暴露
│   ├── config/            # 配置管理
│   │   ├── config.go
│   │   └── config.yaml
│   ├── handler/           # HTTP处理器
│   │   ├── redis.go
│   │   ├── health.go
│   │   └── metrics.go
│   ├── middleware/        # 中间件
│   │   ├── auth.go
│   │   ├── ratelimit.go
│   │   ├── logging.go
│   │   └── cors.go
│   ├── redis/             # Redis客户端封装
│   │   ├── client.go
│   │   ├── pool.go
│   │   └── commands.go
│   ├── router/            # 路由定义
│   │   └── router.go
│   └── utils/             # 工具函数
│       ├── response.go
│       └── validator.go
├── pkg/                   # 可对外使用的包
│   ├── types/             # 类型定义
│   │   └── response.go
│   └── errors/            # 错误定义
│       └── errors.go
├── api/                   # API文档
│   ├── openapi.yaml       # OpenAPI 3.0规范文件
│   └── swagger/           # Swagger自动生成文档
│       ├── docs.go
│       ├── swagger.json
│       └── swagger.yaml
├── configs/               # 配置文件
│   ├── config.yaml
│   └── config.example.yaml
├── scripts/               # 部署和工具脚本
│   ├── build.sh
│   └── docker/
│       ├── Dockerfile
│       └── docker-compose.yml
├── test/                  # 测试文件
│   ├── integration/
│   └── benchmark/
├── docs/                  # 文档
│   ├── api.md
│   ├── deployment.md
│   ├── configuration.md
│   └── swagger-ui/        # 自定义Swagger UI资源
├── go.mod
├── go.sum
├── Makefile
└── README.md
```

## API文档系统

### OpenAPI/Swagger集成方案

本项目采用 **OpenAPI 3.0规范** 和 **Swagger** 生态系统来提供完整的API文档解决方案。

#### 技术选型

- **swaggo/swag**: Go语言的Swagger文档生成工具
- **swaggo/gin-swagger**: Gin框架的Swagger中间件
- **swaggo/files**: 静态文件服务支持
- **OpenAPI 3.0**: 现代API规范标准

#### 文档功能特性

1. **自动化文档生成**
   - 通过代码注释自动生成API文档
   - 支持实时更新和版本管理
   - 类型安全的参数和响应模型

2. **交互式API界面**
   - Swagger UI提供可视化接口文档
   - 在线API测试和调试功能
   - 支持多种认证方式测试

3. **多格式支持**
   - JSON格式API规范
   - YAML格式API规范
   - 可导出OpenAPI标准文档

4. **开发者友好**
   - 代码示例和SDK生成
   - 多语言客户端支持
   - API变更追踪和兼容性检查

#### 文档访问方式

```bash
# Swagger UI界面
GET /swagger/index.html

# JSON格式API文档
GET /swagger/doc.json

# YAML格式API文档  
GET /swagger/doc.yaml

# 自定义API文档页面
GET /api/docs
```

#### 代码注释规范

```go
// @title Go Redis Proxy API
// @version 1.0.0
// @description Redis HTTP代理服务API文档
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /api/v1

// @Summary Redis字符串GET操作
// @Description 根据指定的key获取Redis中存储的字符串值
// @Tags Redis String Operations
// @Accept json
// @Produce json
// @Param request body types.StringGetRequest true "请求参数"
// @Success 200 {object} types.StringGetResponse "成功响应"
// @Failure 400 {object} types.ErrorResponse "请求参数错误"
// @Failure 500 {object} types.ErrorResponse "服务器内部错误"
// @Router /redis/string/get [post]
func RedisStringGet(client RedisClient) gin.HandlerFunc {
    // 处理逻辑...
}
```

#### 文档生成命令

```bash
# 安装swag工具
go install github.com/swaggo/swag/cmd/swag@latest

# 生成Swagger文档
swag init -g cmd/server/main.go -o api/swagger/

# 格式化API文档
swag fmt

# 验证API规范
swag validate
```

#### 集成配置

```go
// main.go
import (
    "github.com/gin-gonic/gin"
    ginSwagger "github.com/swaggo/gin-swagger"
    "github.com/swaggo/gin-swagger/swaggerFiles"
    
    _ "github.com/ct-zh/go-redis-proxy/api/swagger"
)

func main() {
    r := gin.Default()
    
    // Swagger路由
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
    
    // API路由...
    r.Run(":8080")
}
```

## API设计示例

### 基础Redis操作

```bash
# 设置键值
POST /api/v1/keys/{key}
Content-Type: application/json
{
  "value": "hello world",
  "ttl": 3600
}

# 获取键值
GET /api/v1/keys/{key}

# 删除键
DELETE /api/v1/keys/{key}

# Hash操作
POST /api/v1/hashes/{key}
{
  "field": "name",
  "value": "john"
}

GET /api/v1/hashes/{key}/{field}

# List操作
POST /api/v1/lists/{key}/push
{
  "values": ["item1", "item2"],
  "direction": "left"
}

GET /api/v1/lists/{key}?start=0&stop=-1
```

### 管理接口

```bash
# 健康检查
GET /health

# 指标监控
GET /metrics

# 配置信息
GET /api/v1/config
```

## 开发计划

### Phase 1: 基础框架 (Week 1-2)
- [x] 项目初始化和目录结构
- [ ] 基础配置管理
- [ ] HTTP服务器搭建
- [ ] Redis客户端连接
- [ ] 基本的GET/SET操作
- [ ] 错误处理和日志

### Phase 2: 核心功能 (Week 3-4)
- [ ] 完整的Redis数据类型支持
- [ ] RESTful API设计实现
- [ ] 连接池优化
- [ ] 中间件系统
- [ ] API文档系统集成（Swagger/OpenAPI）
- [ ] 单元测试

### Phase 3: 扩展功能 (Week 5-6)
- [ ] 请求统计和监控
- [ ] 访问控制和认证
- [ ] 速率限制
- [ ] 配置热重载
- [ ] 集成测试

### Phase 4: 高级功能 (Week 7-8)
- [ ] 多Redis实例支持
- [ ] 负载均衡
- [ ] 事务支持
- [ ] 性能优化
- [ ] 部署文档

### Phase 5: 生产就绪 (Week 9-10)
- [ ] Docker支持
- [ ] 监控集成
- [ ] 压力测试
- [ ] 安全审计
- [ ] 文档完善

## 技术栈

- **语言**: Go 1.24.3+
- **Web框架**: Gin
- **Redis客户端**: go-redis/redis/v8
- **配置管理**: Viper
- **日志**: Logrus
- **测试**: Testify
- **API文档**: Swagger/OpenAPI 3.0 (swaggo/swag)
- **容器化**: Docker

## 快速开始

```bash
# 克隆项目
git clone git@github.com:ct-zh/go-redis-proxy.git
cd go-redis-proxy

# 初始化依赖
go mod init github.com/ct-zh/go-redis-proxy
go mod tidy

# 启动Redis（如果需要）
docker run -d -p 6379:6379 redis:latest

# 运行项目
go run cmd/server/main.go

# 测试API
curl http://localhost:8080/health

# 访问API文档
open http://localhost:8080/swagger/index.html
```

## 贡献

欢迎提交Issue和Pull Request来帮助改进项目。

## 许可证

MIT License