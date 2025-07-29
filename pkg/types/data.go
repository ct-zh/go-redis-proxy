package types

// 业务数据类型定义，用于Service层返回

// StringGetData String GET操作的业务数据
type StringGetData struct {
	Value interface{} `json:"value"`
}

// StringSetData String SET操作的业务数据
type StringSetData struct {
	Result string `json:"result"`
}

// StringDelData String DEL操作的业务数据
type StringDelData struct {
	Deleted int64 `json:"deleted"`
}

// StringExistsData String EXISTS操作的业务数据
type StringExistsData struct {
	Exists bool `json:"exists"`
}

// StringIncrData String INCR操作的业务数据
type StringIncrData struct {
	Value int64 `json:"value"`
}

// StringDecrData String DECR操作的业务数据
type StringDecrData struct {
	Value int64 `json:"value"`
}

// StringExpireData String EXPIRE操作的业务数据
type StringExpireData struct {
	Success bool `json:"success"`
}

// List操作的业务数据类型

// ListLPushData List LPUSH操作的业务数据
type ListLPushData struct {
	Length int64 `json:"length"`
}

// ListRPushData List RPUSH操作的业务数据
type ListRPushData struct {
	Length int64 `json:"length"`
}

// ListLPopData List LPOP操作的业务数据
type ListLPopData struct {
	Value interface{} `json:"value"`
}

// ListRPopData List RPOP操作的业务数据
type ListRPopData struct {
	Value interface{} `json:"value"`
}

// ListLRemData List LREM操作的业务数据
type ListLRemData struct {
	Removed int64 `json:"removed"`
}

// ListLIndexData List LINDEX操作的业务数据
type ListLIndexData struct {
	Value interface{} `json:"value"`
}

// ListLRangeData List LRANGE操作的业务数据
type ListLRangeData struct {
	Values []string `json:"values"`
}

// ListLLenData List LLEN操作的业务数据
type ListLLenData struct {
	Length int64 `json:"length"`
}

// ListLTrimData List LTRIM操作的业务数据
type ListLTrimData struct {
	Result string `json:"result"`
}

// ZSet操作的业务数据类型

// ZSetMember 表示ZSet成员及其分数
type ZSetMember struct {
	Member string  `json:"member"`
	Score  float64 `json:"score"`
}

// ZSetZAddData ZSet ZADD操作的业务数据
type ZSetZAddData struct {
	Added int64 `json:"added"` // 新添加的成员数量
}

// ZSetZIncrByData ZSet ZINCRBY操作的业务数据
type ZSetZIncrByData struct {
	Score float64 `json:"score"` // 新的分数
}

// ZSetZScoreData ZSet ZSCORE操作的业务数据
type ZSetZScoreData struct {
	Score interface{} `json:"score"` // 分数，如果成员不存在则为nil
}

// ZSetZCardData ZSet ZCARD操作的业务数据
type ZSetZCardData struct {
	Count int64 `json:"count"` // 集合中元素的数量
}

// ZSetZCountData ZSet ZCOUNT操作的业务数据
type ZSetZCountData struct {
	Count int64 `json:"count"` // 指定分数范围内的元素数量
}

// ZSetZRankData ZSet ZRANK操作的业务数据
type ZSetZRankData struct {
	Rank interface{} `json:"rank"` // 排名，如果成员不存在则为nil
}

// ZSetZRevRankData ZSet ZREVRANK操作的业务数据
type ZSetZRevRankData struct {
	Rank interface{} `json:"rank"` // 排名，如果成员不存在则为nil
}

// ZSetZRangeData ZSet ZRANGE操作的业务数据
type ZSetZRangeData struct {
	Members []interface{} `json:"members"` // 成员列表，可能包含分数
}

// ZSetZRevRangeData ZSet ZREVRANGE操作的业务数据
type ZSetZRevRangeData struct {
	Members []interface{} `json:"members"` // 成员列表，可能包含分数
}

// ZSetZRangeByScoreData ZSet ZRANGEBYSCORE操作的业务数据
type ZSetZRangeByScoreData struct {
	Members []interface{} `json:"members"` // 成员列表，可能包含分数
}

// ZSetZRevRangeByScoreData ZSet ZREVRANGEBYSCORE操作的业务数据
type ZSetZRevRangeByScoreData struct {
	Members []interface{} `json:"members"` // 成员列表，可能包含分数
}

// ZSetZRemData ZSet ZREM操作的业务数据
type ZSetZRemData struct {
	Removed int64 `json:"removed"` // 被移除的成员数量
}

// ZSetZRemRangeByRankData ZSet ZREMRANGEBYRANK操作的业务数据
type ZSetZRemRangeByRankData struct {
	Removed int64 `json:"removed"` // 被移除的成员数量
}

// ZSetZRemRangeByScoreData ZSet ZREMRANGEBYSCORE操作的业务数据
type ZSetZRemRangeByScoreData struct {
	Removed int64 `json:"removed"` // 被移除的成员数量
}

// Hash操作的业务数据类型

// HashHSetData Hash HSET操作的业务数据
type HashHSetData struct {
	Set int64 `json:"set"` // 被设置的字段数量
}

// HashHGetData Hash HGET操作的业务数据
type HashHGetData struct {
	Value interface{} `json:"value"` // 字段值，如果字段不存在则为nil
}

// HashHMGetData Hash HMGET操作的业务数据
type HashHMGetData struct {
	Values []interface{} `json:"values"` // 字段值列表，不存在的字段值为nil
}

// HashHGetAllData Hash HGETALL操作的业务数据
type HashHGetAllData struct {
	Fields map[string]string `json:"fields"` // 所有字段和值的映射
}

// HashHDelData Hash HDEL操作的业务数据
type HashHDelData struct {
	Deleted int64 `json:"deleted"` // 被删除的字段数量
}

// HashHExistsData Hash HEXISTS操作的业务数据
type HashHExistsData struct {
	Exists bool `json:"exists"` // 字段是否存在
}

// HashHLenData Hash HLEN操作的业务数据
type HashHLenData struct {
	Length int64 `json:"length"` // 哈希表中字段的数量
}

// HashHKeysData Hash HKEYS操作的业务数据
type HashHKeysData struct {
	Keys []string `json:"keys"` // 所有字段名列表
}

// HashHValsData Hash HVALS操作的业务数据
type HashHValsData struct {
	Values []string `json:"values"` // 所有字段值列表
}

// HashHIncrByData Hash HINCRBY操作的业务数据
type HashHIncrByData struct {
	Value int64 `json:"value"` // 增量操作后的值
}