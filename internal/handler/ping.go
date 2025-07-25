package handler

import (
	"net/http"
	"time"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	
	// 手动构建JSON响应
	jsonResponse := `{"success":true,"data":{"message":"pong","timestamp":"` + time.Now().Format(time.RFC3339) + `"}}`
	w.Write([]byte(jsonResponse))
}