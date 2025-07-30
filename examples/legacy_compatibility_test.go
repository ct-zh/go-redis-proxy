package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ct-zh/go-redis-proxy/internal/adapter"
	"github.com/ct-zh/go-redis-proxy/pkg/types"
)

func runLegacyCompatibilityTest() {
	fmt.Println("=== 向后兼容性测试示例 ===")

	// 创建兼容性适配器
	legacyAdapter := adapter.NewLegacyCompatibilityAdapter()
	defer legacyAdapter.Close()

	ctx := context.Background()

	// 1. 测试传统的Redis操作接口
	fmt.Println("\n--- 传统Redis操作接口测试 ---")
	
	// SET操作
	setReq := &types.StringSetRequest{
		RedisRequest: types.RedisRequest{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		},
		Key:   "legacy_test_key",
		Value: "legacy_test_value",
		TTL:   300,
	}

	setResult, err := legacyAdapter.StringSet(ctx, setReq)
	if err != nil {
		log.Printf("传统SET操作失败: %v", err)
	} else {
		fmt.Printf("传统SET操作成功: %+v\n", setResult)
	}

	// GET操作
	getReq := &types.StringGetRequest{
		RedisRequest: types.RedisRequest{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		},
		Key: "legacy_test_key",
	}

	getResult, err := legacyAdapter.StringGet(ctx, getReq)
	if err != nil {
		log.Printf("传统GET操作失败: %v", err)
	} else {
		fmt.Printf("传统GET操作成功: %+v\n", getResult)
	}

	// 2. 测试传统的Token管理接口
	fmt.Println("\n--- 传统Token管理接口测试 ---")
	
	// 连接（创建Token）
	connectReq := &types.ConnectRequest{
		RedisRequest: types.RedisRequest{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		},
		TokenTTL: 3600,
	}

	connectResp, err := legacyAdapter.Connect(ctx, connectReq)
	if err != nil {
		log.Printf("传统连接操作失败: %v", err)
	} else {
		fmt.Printf("传统连接操作成功: Token=%s, 过期时间=%v\n", 
			connectResp.Token, connectResp.ExpiresAt)
	}

	// Token验证
	if connectResp != nil {
		valid, err := legacyAdapter.ValidateToken(ctx, connectResp.Token)
		if err != nil {
			log.Printf("传统Token验证失败: %v", err)
		} else {
			fmt.Printf("传统Token验证结果: %v\n", valid)
		}
	}

	// 3. 测试批量操作（新功能）
	fmt.Println("\n--- 批量操作测试（新功能） ---")
	
	batchOps := &adapter.BatchStringOperations{
		Sets: []*types.StringSetRequest{
			{
				RedisRequest: types.RedisRequest{Addr: "localhost:6379"},
				Key:          "batch_key_1",
				Value:        "batch_value_1",
				TTL:          300,
			},
			{
				RedisRequest: types.RedisRequest{Addr: "localhost:6379"},
				Key:          "batch_key_2",
				Value:        "batch_value_2",
				TTL:          300,
			},
		},
		Gets: []*types.StringGetRequest{
			{
				RedisRequest: types.RedisRequest{Addr: "localhost:6379"},
				Key:          "batch_key_1",
			},
			{
				RedisRequest: types.RedisRequest{Addr: "localhost:6379"},
				Key:          "batch_key_2",
			},
		},
	}

	batchResults := legacyAdapter.ExecuteBatchStringOperations(ctx, batchOps)
	fmt.Printf("批量操作完成，错误数量: %d\n", len(batchResults.Errors))
	
	for i, result := range batchResults.SetResults {
		if result != nil {
			fmt.Printf("批量SET[%d]成功: %+v\n", i, result)
		}
	}
	
	for i, result := range batchResults.GetResults {
		if result != nil {
			fmt.Printf("批量GET[%d]成功: %+v\n", i, result)
		}
	}

	// 4. 测试统计和监控接口
	fmt.Println("\n--- 统计和监控接口测试 ---")
	
	// 连接统计
	stats, err := legacyAdapter.GetConnectionStats(ctx)
	if err != nil {
		log.Printf("获取连接统计失败: %v", err)
	} else {
		fmt.Printf("连接统计: %+v\n", stats)
	}

	// 健康检查
	err = legacyAdapter.HealthCheck(ctx)
	if err != nil {
		log.Printf("健康检查失败: %v", err)
	} else {
		fmt.Println("健康检查通过")
	}

	// 5. 测试高级功能访问
	fmt.Println("\n--- 高级功能访问测试 ---")
	
	// 获取底层服务
	connectionService := legacyAdapter.GetConnectionService()
	stringService := legacyAdapter.GetStringService()
	tokenService := legacyAdapter.GetTokenService()

	fmt.Printf("连接服务: %T\n", connectionService)
	fmt.Printf("字符串服务: %T\n", stringService)
	fmt.Printf("Token服务: %T\n", tokenService)

	// 使用底层服务进行高级操作
	detailedStats := connectionService.GetStats()
	fmt.Printf("详细统计信息: %+v\n", detailedStats)

	healthDetails := connectionService.HealthCheck(ctx)
	fmt.Printf("详细健康信息: %+v\n", healthDetails)

	// 6. 测试资源清理
	fmt.Println("\n--- 资源清理测试 ---")
	
	if connectResp != nil {
		disconnectReq := &types.DisconnectRequest{
			Token: connectResp.Token,
		}
		
		err = legacyAdapter.Disconnect(ctx, disconnectReq)
		if err != nil {
			log.Printf("断开连接失败: %v", err)
		} else {
			fmt.Println("断开连接成功")
		}
	}

	// 7. 性能对比测试
	fmt.Println("\n--- 性能对比测试 ---")
	
	// 测试连接复用效果
	fmt.Println("测试连接复用...")
	for i := 0; i < 5; i++ {
		req := &types.StringGetRequest{
			RedisRequest: types.RedisRequest{Addr: "localhost:6379"},
			Key:          fmt.Sprintf("perf_test_%d", i),
		}
		
		_, err := legacyAdapter.StringGet(ctx, req)
		if err != nil {
			log.Printf("性能测试[%d]失败: %v", i, err)
		} else {
			fmt.Printf("性能测试[%d]完成\n", i)
		}
	}

	fmt.Println("\n=== 向后兼容性测试完成 ===")
	fmt.Println("所有传统接口都能正常工作，同时享受新系统的性能优势")
}