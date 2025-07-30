package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ct-zh/go-redis-proxy/internal/connection"
	"github.com/ct-zh/go-redis-proxy/internal/dao"
	"github.com/ct-zh/go-redis-proxy/internal/service"
	"github.com/ct-zh/go-redis-proxy/pkg/types"
)

func main() {
	fmt.Println("=== 连接管理系统集成示例 ===")

	// 1. 初始化连接服务
	connectionService := connection.NewService()
	defer connectionService.Close()

	// 2. 初始化DAO层
	dynamicDAO := dao.NewDynamicRedisDAO()

	// 3. 初始化业务服务层
	stringService := service.NewDynamicRedisStringService(connectionService, dynamicDAO)
	tokenService := service.NewSimpleTokenService(connectionService)

	ctx := context.Background()

	// 4. 演示Token管理
	fmt.Println("\n--- Token管理演示 ---")
	
	// 创建Token
	connectReq := &types.ConnectRequest{
		RedisRequest: types.RedisRequest{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		},
		TokenTTL: 3600, // 1小时
	}

	tokenResp, err := tokenService.CreateToken(ctx, connectReq)
	if err != nil {
		log.Printf("创建Token失败: %v", err)
	} else {
		fmt.Printf("Token创建成功: %s, 过期时间: %v\n", tokenResp.Token, tokenResp.ExpiresAt)
	}

	// 验证Token
	if tokenResp != nil {
		valid, err := tokenService.ValidateToken(ctx, tokenResp.Token)
		if err != nil {
			log.Printf("验证Token失败: %v", err)
		} else {
			fmt.Printf("Token验证结果: %v\n", valid)
		}
	}

	// 5. 演示基于Token的Redis操作（注意：当前DynamicRedisStringService不支持Token操作）
	fmt.Println("\n--- 基于Token的Redis操作演示 ---")
	fmt.Println("注意：当前版本的DynamicRedisStringService不支持Token操作")
	fmt.Println("Token操作需要使用专门的TokenRedisService")

	// 6. 演示直接Redis请求操作
	fmt.Println("\n--- 直接Redis请求操作演示 ---")
	
	// SET操作
	directSetReq := &types.StringSetRequest{
		RedisRequest: types.RedisRequest{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		},
		Key:   "direct_key",
		Value: "direct_value",
		TTL:   300,
	}

	directSetResult, err := stringService.Set(ctx, directSetReq)
	if err != nil {
		log.Printf("直接SET操作失败: %v", err)
	} else {
		fmt.Printf("直接SET操作成功: %+v\n", directSetResult)
	}

	// GET操作
	directGetReq := &types.StringGetRequest{
		RedisRequest: types.RedisRequest{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		},
		Key: "direct_key",
	}

	directGetResult, err := stringService.Get(ctx, directGetReq)
	if err != nil {
		log.Printf("直接GET操作失败: %v", err)
	} else {
		fmt.Printf("直接GET操作成功: %+v\n", directGetResult)
	}

	// 7. 获取统计信息
	fmt.Println("\n--- 统计信息 ---")
	
	stats, err := tokenService.GetConnectionStats(ctx)
	if err != nil {
		log.Printf("获取统计信息失败: %v", err)
	} else {
		fmt.Printf("连接统计: %+v\n", stats)
	}

	// 8. 健康检查
	fmt.Println("\n--- 健康检查 ---")
	
	err = tokenService.HealthCheck(ctx)
	if err != nil {
		log.Printf("健康检查失败: %v", err)
	} else {
		fmt.Println("健康检查通过")
	}

	// 9. 清理资源
	fmt.Println("\n--- 清理资源 ---")
	
	if tokenResp != nil {
		disconnectReq := &types.DisconnectRequest{
			Token: tokenResp.Token,
		}
		
		err = tokenService.DeleteToken(ctx, disconnectReq)
		if err != nil {
			log.Printf("删除Token失败: %v", err)
		} else {
			fmt.Println("Token删除成功")
		}
	}

	fmt.Println("\n=== 集成示例完成 ===")
}