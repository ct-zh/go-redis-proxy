package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ct-zh/go-redis-proxy/internal/adapter"
	"github.com/ct-zh/go-redis-proxy/internal/connection"
	"github.com/ct-zh/go-redis-proxy/internal/dao"
	"github.com/ct-zh/go-redis-proxy/internal/service"
	"github.com/ct-zh/go-redis-proxy/pkg/types"
)

// PerformanceBenchmark 性能基准测试
type PerformanceBenchmark struct {
	legacyAdapter     *adapter.LegacyCompatibilityAdapter
	connectionService *connection.Service
	stringService     *service.DynamicRedisStringService
}

// NewPerformanceBenchmark 创建性能基准测试实例
func NewPerformanceBenchmark() *PerformanceBenchmark {
	// 初始化服务
	connectionService := connection.NewService()
	dynamicDAO := dao.NewDynamicRedisDAO()
	stringService := service.NewDynamicRedisStringService(connectionService, dynamicDAO)
	legacyAdapter := adapter.NewLegacyCompatibilityAdapter()

	return &PerformanceBenchmark{
		legacyAdapter:     legacyAdapter,
		connectionService: connectionService,
		stringService:     stringService,
	}
}

// Close 关闭资源
func (pb *PerformanceBenchmark) Close() {
	if pb.legacyAdapter != nil {
		pb.legacyAdapter.Close()
	}
	if pb.connectionService != nil {
		pb.connectionService.Close()
	}
}

// BenchmarkStringOperations 字符串操作性能测试
func (pb *PerformanceBenchmark) BenchmarkStringOperations(iterations int) {
	fmt.Printf("\n=== 字符串操作性能测试 (迭代次数: %d) ===\n", iterations)
	
	ctx := context.Background()
	redisReq := types.RedisRequest{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	}

	// 测试传统接口性能
	fmt.Println("\n--- 传统接口性能测试 ---")
	start := time.Now()
	
	for i := 0; i < iterations; i++ {
		setReq := &types.StringSetRequest{
			RedisRequest: redisReq,
			Key:          fmt.Sprintf("legacy_perf_key_%d", i),
			Value:        fmt.Sprintf("legacy_perf_value_%d", i),
			TTL:          300,
		}
		
		_, err := pb.legacyAdapter.StringSet(ctx, setReq)
		if err != nil {
			log.Printf("传统SET操作失败[%d]: %v", i, err)
			continue
		}
		
		getReq := &types.StringGetRequest{
			RedisRequest: redisReq,
			Key:          fmt.Sprintf("legacy_perf_key_%d", i),
		}
		
		_, err = pb.legacyAdapter.StringGet(ctx, getReq)
		if err != nil {
			log.Printf("传统GET操作失败[%d]: %v", i, err)
		}
	}
	
	legacyDuration := time.Since(start)
	fmt.Printf("传统接口完成时间: %v\n", legacyDuration)
	fmt.Printf("传统接口平均每次操作: %v\n", legacyDuration/time.Duration(iterations*2))

	// 测试新接口性能
	fmt.Println("\n--- 新接口性能测试 ---")
	start = time.Now()
	
	for i := 0; i < iterations; i++ {
		setReq := &types.StringSetRequest{
			RedisRequest: redisReq,
			Key:          fmt.Sprintf("new_perf_key_%d", i),
			Value:        fmt.Sprintf("new_perf_value_%d", i),
			TTL:          300,
		}
		
		_, err := pb.stringService.Set(ctx, setReq)
		if err != nil {
			log.Printf("新SET操作失败[%d]: %v", i, err)
			continue
		}
		
		getReq := &types.StringGetRequest{
			RedisRequest: redisReq,
			Key:          fmt.Sprintf("new_perf_key_%d", i),
		}
		
		_, err = pb.stringService.Get(ctx, getReq)
		if err != nil {
			log.Printf("新GET操作失败[%d]: %v", i, err)
		}
	}
	
	newDuration := time.Since(start)
	fmt.Printf("新接口完成时间: %v\n", newDuration)
	fmt.Printf("新接口平均每次操作: %v\n", newDuration/time.Duration(iterations*2))
	
	// 性能对比
	if legacyDuration > 0 {
		improvement := float64(legacyDuration-newDuration) / float64(legacyDuration) * 100
		fmt.Printf("\n性能提升: %.2f%%\n", improvement)
	}
}

// BenchmarkConnectionReuse 连接复用性能测试
func (pb *PerformanceBenchmark) BenchmarkConnectionReuse(iterations int) {
	fmt.Printf("\n=== 连接复用性能测试 (迭代次数: %d) ===\n", iterations)
	
	ctx := context.Background()
	redisReq := types.RedisRequest{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	}

	// 测试连接复用效果
	fmt.Println("\n--- 连接复用测试 ---")
	start := time.Now()
	
	for i := 0; i < iterations; i++ {
		getReq := &types.StringGetRequest{
			RedisRequest: redisReq,
			Key:          "connection_reuse_test",
		}
		
		_, err := pb.stringService.Get(ctx, getReq)
		if err != nil {
			log.Printf("连接复用测试失败[%d]: %v", i, err)
		}
	}
	
	duration := time.Since(start)
	fmt.Printf("连接复用测试完成时间: %v\n", duration)
	fmt.Printf("平均每次操作: %v\n", duration/time.Duration(iterations))
	
	// 获取连接统计
	stats := pb.connectionService.GetStats()
	fmt.Printf("连接统计: %+v\n", stats)
}

// BenchmarkBatchOperations 批量操作性能测试
func (pb *PerformanceBenchmark) BenchmarkBatchOperations(batchSize int) {
	fmt.Printf("\n=== 批量操作性能测试 (批量大小: %d) ===\n", batchSize)
	
	ctx := context.Background()
	redisReq := types.RedisRequest{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	}

	// 准备批量操作数据
	batchOps := &adapter.BatchStringOperations{
		Sets: make([]*types.StringSetRequest, batchSize),
		Gets: make([]*types.StringGetRequest, batchSize),
	}
	
	for i := 0; i < batchSize; i++ {
		batchOps.Sets[i] = &types.StringSetRequest{
			RedisRequest: redisReq,
			Key:          fmt.Sprintf("batch_perf_key_%d", i),
			Value:        fmt.Sprintf("batch_perf_value_%d", i),
			TTL:          300,
		}
		
		batchOps.Gets[i] = &types.StringGetRequest{
			RedisRequest: redisReq,
			Key:          fmt.Sprintf("batch_perf_key_%d", i),
		}
	}

	// 测试批量操作性能
	fmt.Println("\n--- 批量操作测试 ---")
	start := time.Now()
	
	results := pb.legacyAdapter.ExecuteBatchStringOperations(ctx, batchOps)
	
	duration := time.Since(start)
	fmt.Printf("批量操作完成时间: %v\n", duration)
	fmt.Printf("平均每次操作: %v\n", duration/time.Duration(batchSize*2))
	fmt.Printf("成功的SET操作: %d/%d\n", len(results.SetResults), batchSize)
	fmt.Printf("成功的GET操作: %d/%d\n", len(results.GetResults), batchSize)
	fmt.Printf("错误数量: %d\n", len(results.Errors))

	// 对比单个操作性能
	fmt.Println("\n--- 单个操作对比测试 ---")
	start = time.Now()
	
	for i := 0; i < batchSize; i++ {
		_, err := pb.legacyAdapter.StringSet(ctx, batchOps.Sets[i])
		if err != nil {
			log.Printf("单个SET操作失败[%d]: %v", i, err)
		}
		
		_, err = pb.legacyAdapter.StringGet(ctx, batchOps.Gets[i])
		if err != nil {
			log.Printf("单个GET操作失败[%d]: %v", i, err)
		}
	}
	
	singleDuration := time.Since(start)
	fmt.Printf("单个操作完成时间: %v\n", singleDuration)
	fmt.Printf("平均每次操作: %v\n", singleDuration/time.Duration(batchSize*2))
	
	if singleDuration > 0 {
		improvement := float64(singleDuration-duration) / float64(singleDuration) * 100
		fmt.Printf("批量操作性能提升: %.2f%%\n", improvement)
	}
}

// BenchmarkMemoryUsage 内存使用测试
func (pb *PerformanceBenchmark) BenchmarkMemoryUsage() {
	fmt.Println("\n=== 内存使用测试 ===")
	
	// 这里可以添加内存使用监控逻辑
	// 由于Go的runtime包提供了内存统计功能，可以在这里实现
	fmt.Println("内存使用监控功能待实现...")
}

// RunAllBenchmarks 运行所有基准测试
func (pb *PerformanceBenchmark) RunAllBenchmarks() {
	fmt.Println("=== 开始性能基准测试 ===")
	
	// 字符串操作性能测试
	pb.BenchmarkStringOperations(100)
	
	// 连接复用性能测试
	pb.BenchmarkConnectionReuse(200)
	
	// 批量操作性能测试
	pb.BenchmarkBatchOperations(50)
	
	// 内存使用测试
	pb.BenchmarkMemoryUsage()
	
	fmt.Println("\n=== 性能基准测试完成 ===")
}

func runPerformanceBenchmark() {
	benchmark := NewPerformanceBenchmark()
	defer benchmark.Close()
	
	benchmark.RunAllBenchmarks()
}