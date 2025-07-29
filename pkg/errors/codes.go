package errors

// 错误码分类定义

// 系统级错误 1000-1999
const (
	CodeInternalError = 1000 // 内部服务器错误
	CodeInvalidParams = 1001 // 参数验证失败
	CodeMethodNotAllowed = 1002 // 方法不允许
	CodeUnauthorized = 1003 // 未授权访问
)

// Redis连接错误 2000-2099
const (
	CodeRedisConnectFailed = 2000 // Redis连接失败
	CodeRedisTimeout       = 2001 // Redis操作超时
	CodeRedisAuthFailed    = 2002 // Redis认证失败
	CodeRedisDBSelectFailed = 2003 // Redis数据库选择失败
)

// String操作错误 2100-2199
const (
	CodeStringKeyNotFound   = 2100 // 键不存在
	CodeStringTypeMismatch  = 2101 // 类型不匹配
	CodeStringSetFailed     = 2102 // 设置失败
	CodeStringGetFailed     = 2103 // 获取失败
	CodeStringDelFailed     = 2104 // 删除失败
	CodeStringIncrFailed    = 2105 // 自增失败
	CodeStringDecrFailed    = 2106 // 自减失败
	CodeStringExpireFailed  = 2107 // 设置过期时间失败
)

// List操作错误 2200-2299
const (
	CodeListIndexOutOfRange = 2200 // 索引超出范围
	CodeListTypeMismatch    = 2201 // 类型不匹配
	CodeListPushFailed      = 2202 // 推入失败
	CodeListPopFailed       = 2203 // 弹出失败
	CodeListRemoveFailed    = 2204 // 删除失败
	CodeListTrimFailed      = 2205 // 裁剪失败
)

// Hash操作错误 2300-2399
const (
	CodeHashKeyNotFound     = 2300 // Hash键不存在
	CodeHashFieldNotFound   = 2301 // Hash字段不存在
	CodeHashTypeMismatch    = 2302 // 类型不匹配
	CodeHashSetFailed       = 2303 // 设置失败
	CodeHashGetFailed       = 2304 // 获取失败
	CodeHashDelFailed       = 2305 // 删除失败
	CodeHashDeleteFailed    = 2306 // 删除失败（alias）
	CodeHashIncrementFailed = 2307 // 增量操作失败
)

// Set操作错误 2400-2499
const (
	CodeSetTypeMismatch     = 2400 // 类型不匹配
	CodeSetAddFailed        = 2401 // 添加失败
	CodeSetRemoveFailed     = 2402 // 删除失败
	CodeSetMemberNotFound   = 2403 // 成员不存在
)

// ZSet操作错误 2500-2599
const (
	CodeZSetTypeMismatch    = 2500 // 类型不匹配
	CodeZSetAddFailed       = 2501 // 添加失败
	CodeZSetRemoveFailed    = 2502 // 删除失败
	CodeZSetMemberNotFound  = 2503 // 成员不存在
	CodeZSetRankNotFound    = 2504 // 排名不存在
)