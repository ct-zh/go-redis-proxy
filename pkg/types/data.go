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