package types

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

type PingResponse struct {
	Success bool `json:"success"`
	Data    struct {
		Message   string `json:"message"`
		Timestamp string `json:"timestamp"`
	} `json:"data"`
}

// String操作响应类型

type StringGetResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Value interface{} `json:"value"`
	} `json:"data"`
}

type StringSetResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Result string `json:"result"`
	} `json:"data"`
}

type StringDelResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Deleted int `json:"deleted"` // 删除的键数量
	} `json:"data"`
}

type StringExistsResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Exists bool `json:"exists"`
	} `json:"data"`
}

type StringIncrResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Value int64 `json:"value"` // 增加后的值
	} `json:"data"`
}

type StringDecrResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Value int64 `json:"value"` // 减少后的值
	} `json:"data"`
}

type StringExpireResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Success bool `json:"success"` // 是否设置成功
	} `json:"data"`
}

// List操作响应类型

type ListLPushResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Length int64 `json:"length"` // 推入后列表的长度
	} `json:"data"`
}

type ListRPushResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Length int64 `json:"length"` // 推入后列表的长度
	} `json:"data"`
}

type ListLPopResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Value interface{} `json:"value"` // 弹出的值，可能为null
	} `json:"data"`
}

type ListRPopResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Value interface{} `json:"value"` // 弹出的值，可能为null
	} `json:"data"`
}

type ListLRemResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Removed int64 `json:"removed"` // 删除的元素数量
	} `json:"data"`
}

type ListLIndexResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Value interface{} `json:"value"` // 指定索引的值，可能为null
	} `json:"data"`
}

type ListLRangeResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Values []string `json:"values"` // 范围内的值数组
	} `json:"data"`
}

type ListLLenResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Length int64 `json:"length"` // 列表长度
	} `json:"data"`
}

type ListLTrimResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Result string `json:"result"` // 操作结果，通常为"OK"
	} `json:"data"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
