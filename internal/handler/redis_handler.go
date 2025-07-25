
package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/ct-zh/go-redis-proxy/pkg/types"
	"github.com/go-redis/redis/v8"
)

// RedisClient defines the interface for a Redis client. It's used to allow mocking in tests.
type RedisClient interface {
	Get(ctx context.Context, key string) *redis.StringCmd
}

func RedisStringGet(client RedisClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.StringGetRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// If the client is nil, create a new one. This allows the handler to be used in production
		// without having to pass a client to it.
		var rdb RedisClient
		if client == nil {
			rdb = redis.NewClient(&redis.Options{
				Addr:     req.Addr,
				Password: req.Password,
				DB:       req.DB,
			})
		} else {
			rdb = client
		}

		val, err := rdb.Get(context.Background(), req.Key).Result()
		if err != nil {
			if err == redis.Nil {
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(map[string]interface{}{
					"code": 0,
					"msg":  "success",
					"data": map[string]interface{}{
						"value": nil,
					},
				})
				return
			}
			http.Error(w, "error connecting to redis: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"code": 0,
			"msg":  "success",
			"data": map[string]interface{}{
				"value": val,
			},
		})
	}
}
