package types

// RedisRequest 包含连接Redis所需的基础参数
type RedisRequest struct {
	Addr     string `json:"addr"`
	Password string `json:"password"`
	DB       int    `json:"db"`
}

// StringGetRequest 定义了GET string类型value的请求体
type StringGetRequest struct {
	RedisRequest
	Key string `json:"key"`
}

// StringSetRequest 定义了SET string类型value的请求体
type StringSetRequest struct {
	RedisRequest
	Key    string `json:"key"`
	Value  string `json:"value"`
	TTL    int    `json:"ttl,omitempty"` // 过期时间，单位秒，0表示不过期
}

// StringDelRequest 定义了DEL key的请求体
type StringDelRequest struct {
	RedisRequest
	Key string `json:"key"`
}

// StringExistsRequest 定义了EXISTS key的请求体
type StringExistsRequest struct {
	RedisRequest
	Key string `json:"key"`
}

// StringIncrRequest 定义了INCR key的请求体
type StringIncrRequest struct {
	RedisRequest
	Key string `json:"key"`
}

// StringDecrRequest 定义了DECR key的请求体
type StringDecrRequest struct {
	RedisRequest
	Key string `json:"key"`
}

// StringExpireRequest 定义了EXPIRE key的请求体
type StringExpireRequest struct {
	RedisRequest
	Key string `json:"key"`
	TTL int    `json:"ttl"` // 过期时间，单位秒
}

// List操作请求类型

// ListLPushRequest 定义了LPUSH操作的请求体
type ListLPushRequest struct {
	RedisRequest
	Key    string   `json:"key"`
	Values []string `json:"values"` // 要推入的值数组
}

// ListRPushRequest 定义了RPUSH操作的请求体
type ListRPushRequest struct {
	RedisRequest
	Key    string   `json:"key"`
	Values []string `json:"values"` // 要推入的值数组
}

// ListLPopRequest 定义了LPOP操作的请求体
type ListLPopRequest struct {
	RedisRequest
	Key string `json:"key"`
}

// ListRPopRequest 定义了RPOP操作的请求体
type ListRPopRequest struct {
	RedisRequest
	Key string `json:"key"`
}

// ListLRemRequest 定义了LREM操作的请求体
type ListLRemRequest struct {
	RedisRequest
	Key   string `json:"key"`
	Count int64  `json:"count"` // 删除的数量，0表示删除所有
	Value string `json:"value"` // 要删除的值
}

// ListLIndexRequest 定义了LINDEX操作的请求体
type ListLIndexRequest struct {
	RedisRequest
	Key   string `json:"key"`
	Index int64  `json:"index"` // 索引位置
}

// ListLRangeRequest 定义了LRANGE操作的请求体
type ListLRangeRequest struct {
	RedisRequest
	Key   string `json:"key"`
	Start int64  `json:"start"` // 开始索引
	Stop  int64  `json:"stop"`  // 结束索引
}

// ListLLenRequest 定义了LLEN操作的请求体
type ListLLenRequest struct {
	RedisRequest
	Key string `json:"key"`
}

// ListLTrimRequest 定义了LTRIM操作的请求体
type ListLTrimRequest struct {
	RedisRequest
	Key   string `json:"key"`
	Start int64  `json:"start"` // 开始索引
	Stop  int64  `json:"stop"`  // 结束索引
}

// Set操作请求类型

// RedisSAddRequest 定义了SADD操作的请求体
type RedisSAddRequest struct {
	RedisRequest
	Key     string   `json:"key"`
	Members []string `json:"members"`
}

// RedisSRemRequest 定义了SREM操作的请求体
type RedisSRemRequest struct {
	RedisRequest
	Key     string   `json:"key"`
	Members []string `json:"members"`
}

// RedisSIsMemberRequest 定义了SISMEMBER操作的请求体
type RedisSIsMemberRequest struct {
	RedisRequest
	Key    string `json:"key"`
	Member string `json:"member"`
}

// RedisSMembersRequest 定义了SMEMBERS操作的请求体
type RedisSMembersRequest struct {
	RedisRequest
	Key string `json:"key"`
}

// RedisSCardRequest 定义了SCARD操作的请求体
type RedisSCardRequest struct {
	RedisRequest
	Key string `json:"key"`
}
