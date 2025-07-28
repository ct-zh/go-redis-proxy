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

// ZSet操作请求类型

// ZSetZAddRequest 定义了ZADD操作的请求体
type ZSetZAddRequest struct {
	RedisRequest
	Key     string            `json:"key"`
	Members map[string]float64 `json:"members"` // member -> score 映射
}

// ZSetZIncrByRequest 定义了ZINCRBY操作的请求体
type ZSetZIncrByRequest struct {
	RedisRequest
	Key       string  `json:"key"`
	Increment float64 `json:"increment"`
	Member    string  `json:"member"`
}

// ZSetZScoreRequest 定义了ZSCORE操作的请求体
type ZSetZScoreRequest struct {
	RedisRequest
	Key    string `json:"key"`
	Member string `json:"member"`
}

// ZSetZCardRequest 定义了ZCARD操作的请求体
type ZSetZCardRequest struct {
	RedisRequest
	Key string `json:"key"`
}

// ZSetZCountRequest 定义了ZCOUNT操作的请求体
type ZSetZCountRequest struct {
	RedisRequest
	Key string  `json:"key"`
	Min float64 `json:"min"`
	Max float64 `json:"max"`
}

// ZSetZRankRequest 定义了ZRANK操作的请求体
type ZSetZRankRequest struct {
	RedisRequest
	Key    string `json:"key"`
	Member string `json:"member"`
}

// ZSetZRevRankRequest 定义了ZREVRANK操作的请求体
type ZSetZRevRankRequest struct {
	RedisRequest
	Key    string `json:"key"`
	Member string `json:"member"`
}

// ZSetZRangeRequest 定义了ZRANGE操作的请求体
type ZSetZRangeRequest struct {
	RedisRequest
	Key        string `json:"key"`
	Start      int64  `json:"start"`
	Stop       int64  `json:"stop"`
	WithScores bool   `json:"with_scores,omitempty"` // 是否返回分数
}

// ZSetZRevRangeRequest 定义了ZREVRANGE操作的请求体
type ZSetZRevRangeRequest struct {
	RedisRequest
	Key        string `json:"key"`
	Start      int64  `json:"start"`
	Stop       int64  `json:"stop"`
	WithScores bool   `json:"with_scores,omitempty"` // 是否返回分数
}

// ZSetZRangeByScoreRequest 定义了ZRANGEBYSCORE操作的请求体
type ZSetZRangeByScoreRequest struct {
	RedisRequest
	Key        string `json:"key"`
	Min        string `json:"min"` // 支持 -inf, +inf, (value 等格式
	Max        string `json:"max"` // 支持 -inf, +inf, (value 等格式
	WithScores bool   `json:"with_scores,omitempty"`
	Offset     int64  `json:"offset,omitempty"`
	Count      int64  `json:"count,omitempty"`
}

// ZSetZRevRangeByScoreRequest 定义了ZREVRANGEBYSCORE操作的请求体
type ZSetZRevRangeByScoreRequest struct {
	RedisRequest
	Key        string `json:"key"`
	Max        string `json:"max"` // 支持 -inf, +inf, (value 等格式
	Min        string `json:"min"` // 支持 -inf, +inf, (value 等格式
	WithScores bool   `json:"with_scores,omitempty"`
	Offset     int64  `json:"offset,omitempty"`
	Count      int64  `json:"count,omitempty"`
}

// ZSetZRemRequest 定义了ZREM操作的请求体
type ZSetZRemRequest struct {
	RedisRequest
	Key     string   `json:"key"`
	Members []string `json:"members"`
}

// ZSetZRemRangeByRankRequest 定义了ZREMRANGEBYRANK操作的请求体
type ZSetZRemRangeByRankRequest struct {
	RedisRequest
	Key   string `json:"key"`
	Start int64  `json:"start"`
	Stop  int64  `json:"stop"`
}

// ZSetZRemRangeByScoreRequest 定义了ZREMRANGEBYSCORE操作的请求体
type ZSetZRemRangeByScoreRequest struct {
	RedisRequest
	Key string `json:"key"`
	Min string `json:"min"` // 支持 -inf, +inf, (value 等格式
	Max string `json:"max"` // 支持 -inf, +inf, (value 等格式
}
