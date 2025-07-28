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

## API文档系统

本项目集成了 **Swagger/OpenAPI 3.0** 文档系统，提供完整的API文档和在线测试功能。

### 文档访问

```bash
# 启动服务
go run cmd/server/main.go

# 访问Swagger UI
open http://localhost:8080/swagger/index.html
```

### 技术栈
- **swaggo/swag**: Go语言的Swagger文档生成工具
- **swaggo/gin-swagger**: Gin框架的Swagger中间件  
- **OpenAPI 3.0**: 现代API规范标准

### 文档生成

```bash
# 安装swag工具
go install github.com/swaggo/swag/cmd/swag@latest

# 生成文档
swag init -g cmd/server/main.go -o api/swagger/
```

详细集成流程参见 [NOTES.md](NOTES.md#swaggeropenapi-文档系统集成流程)

## 贡献

欢迎提交Issue和Pull Request来帮助改进项目。

## 许可证

MIT License