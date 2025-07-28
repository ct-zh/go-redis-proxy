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

### Phase 1: 基础框架 (Week 1-2) ✅ 已完成
- [x] 项目初始化和目录结构
- [x] 基础配置管理
- [x] HTTP服务器搭建（Gin）
- [x] Redis客户端连接
- [x] 基本的String操作（GET/SET/DEL/EXISTS/INCR/DECR/EXPIRE）
- [x] List操作实现（LPUSH/RPUSH/LPOP/RPOP/LREM/LINDEX/LRANGE/LLEN/LTRIM）
- [x] 错误处理和基础日志
- [x] API文档系统集成（Swagger/OpenAPI）

### Phase 2: 架构重构优化 (当前阶段)
**目标：改善代码架构分层，提高可维护性和可测试性**

#### 2.1 分层架构重构
- [ ] **Handler层重构**：只负责HTTP请求解析、参数验证、响应格式化
- [ ] **Service/Logic层抽象**：封装业务逻辑，提供统一接口给Handler层
- [ ] **DAO层实现**：封装数据访问操作，屏蔽Redis客户端实现细节
- [ ] **依赖注入模式**：使用依赖注入管理各层依赖关系

#### 2.2 接口设计和实现
- [ ] **Service接口定义**：RedisStringService、RedisListService等业务接口
- [ ] **DAO接口定义**：RedisDAO统一数据访问接口
- [ ] **错误处理统一化**：定义统一错误类型，分层错误转换
- [ ] **配置统一管理**：Redis连接配置集中管理，连接池复用

#### 2.3 响应封装和错误码管理 (新增优化方案)
- [ ] **Handler响应封装**：封装c.JSON方法，自动处理data和error参数
  - 统一响应格式：`{code, msg, data}`
  - Service层只需返回业务数据和错误，无需构建完整response
  - 支持自定义错误类型和标准error的处理
- [ ] **错误码统一管理**：建立集中式错误码管理机制
  - 错误码分类定义（系统级1000+，Redis连接2000+，业务操作2100+）
  - 错误注册机制，避免code冲突
  - 启动时验证错误码完整性
  - 支持错误码按模块分组管理
- [ ] **BusinessError接口**：定义业务错误标准接口
  - Code()方法获取错误码
  - Message()方法获取错误消息
  - 继承标准error接口

#### 2.4 中间件和基础设施
- [ ] **日志中间件**：请求日志记录和结构化日志
- [ ] **性能监控中间件**：请求耗时、QPS等指标收集
- [ ] **请求追踪中间件**：分布式追踪支持
- [ ] **统一响应格式**：标准化API响应结构（已包含在2.3中）

#### 2.5 测试体系完善
- [ ] **单元测试**：各层独立测试，Mock依赖
- [ ] **集成测试**：端到端API测试
- [ ] **基准测试**：性能基准测试

### Phase 3: 核心功能扩展 (Week 3-4)
- [ ] **完整Redis数据类型支持**：Hash、Set、ZSet操作
- [ ] **连接池优化**：连接池配置和监控
- [ ] **事务支持**：Redis事务操作
- [ ] **批量操作**：Pipeline支持
- [ ] **健康检查增强**：Redis连接状态检查

### Phase 4: 扩展功能 (Week 5-6)
- [ ] **请求统计和监控**：实时统计和历史数据
- [ ] **访问控制和认证**：API Key、JWT认证
- [ ] **速率限制**：基于用户/IP的请求限制
- [ ] **请求缓存策略**：智能缓存机制
- [ ] **配置热重载**：运行时配置更新

### Phase 5: 高级功能 (Week 7-8)
- [ ] **多Redis实例支持**：主从、集群、分片
- [ ] **负载均衡**：智能路由和故障转移
- [ ] **发布订阅支持**：Redis Pub/Sub功能
- [ ] **数据备份恢复**：自动备份策略
- [ ] **性能优化**：内存优化、并发优化

### Phase 6: 生产就绪 (Week 9-10)
- [ ] **Docker支持**：容器化部署
- [ ] **Prometheus监控**：指标收集和可视化
- [ ] **压力测试**：性能基准和瓶颈分析
- [ ] **安全审计**：安全扫描和漏洞修复
- [ ] **文档完善**：部署文档、运维手册

### 架构设计原则

#### 分层架构
```
HTTP Request
     ↓
Handler Layer (HTTP处理、参数验证、响应格式化)
     ↓  
Service Layer (业务逻辑、流程控制)
     ↓
DAO Layer (数据访问抽象)
     ↓
Redis Client (具体实现)
```

#### 设计目标
- **单一职责**：每层只关注自己的职责
- **依赖倒置**：高层不依赖低层，都依赖抽象
- **开闭原则**：对扩展开放，对修改封闭
- **可测试性**：每层可独立测试
- **可维护性**：清晰的代码结构和文档

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