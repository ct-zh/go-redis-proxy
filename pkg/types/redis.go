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
