package types

import "time"

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

type PingResponse struct {
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}
