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

type StringGetResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Value interface{} `json:"value"`
	} `json:"data"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
