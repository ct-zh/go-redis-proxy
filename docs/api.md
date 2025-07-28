# Go Redis Proxy API 文档

## 概述

Go Redis Proxy 提供RESTful API来与Redis进行HTTP交互。本文档描述了所有可用的API端点。

## 基础信息

- **Base URL**: `http://localhost:8080`
- **API版本**: `v1`
- **API前缀**: `/api/v1`

## 认证

当前版本暂未实现认证机制，后续版本将支持API Key或JWT认证。

## 响应格式

所有API响应都遵循统一的JSON格式：

### 成功响应
```json
{
  "code": 0,
  "msg": "success",
  "data": {
    // 具体数据
  }
}
```

### 错误响应
```json
{
  "error": "错误描述信息"
}
```

## API端点

### 健康检查

#### Ping
- **URL**: `/ping`
- **方法**: `GET`
- **描述**: 检查服务是否正常运行
- **响应示例**:
```json
{
  "success": true,
  "data": {
    "message": "pong",
    "timestamp": "2025-07-28T10:00:00Z"
  }
}
```

#### Health Check
- **URL**: `/health`
- **方法**: `GET`
- **描述**: 健康检查端点（同ping）

### Redis字符串操作

#### GET字符串
- **URL**: `/api/v1/redis/string/get`
- **方法**: `POST`
- **描述**: 获取Redis中指定key的字符串值
- **请求体**:
```json
{
  "addr": "localhost:6379",
  "password": "",
  "db": 0,
  "key": "mykey"
}
```
- **响应示例**:
```json
{
  "code": 0,
  "msg": "success",
  "data": {
    "value": "myvalue"
  }
}
```

当key不存在时：
```json
{
  "code": 0,
  "msg": "success",
  "data": {
    "value": null
  }
}
```

## Swagger文档

项目集成了Swagger文档系统，提供交互式API文档：

- **Swagger UI**: `http://localhost:8080/swagger/index.html`
- **JSON格式**: `http://localhost:8080/swagger/doc.json`
- **YAML格式**: `http://localhost:8080/swagger/doc.yaml`

## 错误码

| 错误码 | HTTP状态码 | 描述 |
|--------|------------|------|
| 400 | Bad Request | 请求参数错误 |
| 500 | Internal Server Error | 服务器内部错误或Redis连接失败 |

## 示例

### 使用curl测试API

```bash
# 健康检查
curl http://localhost:8080/ping

# 获取Redis字符串
curl -X POST http://localhost:8080/api/v1/redis/string/get \
  -H "Content-Type: application/json" \
  -d '{
    "addr": "localhost:6379",
    "password": "",
    "db": 0,
    "key": "test_key"
  }'
```

## 后续计划

### 即将支持的功能：
- Redis Hash操作
- Redis List操作
- Redis Set操作
- Redis ZSet操作
- 批量操作
- 事务支持
- 用户认证
- 速率限制
- 监控指标