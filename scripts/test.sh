#!/bin/bash

# 测试运行脚本
# 用于运行项目的各种测试

set -e

# 颜色定义
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# 打印带颜色的消息
print_message() {
    local color=$1
    local message=$2
    echo -e "${color}${message}${NC}"
}

# 检查Redis是否运行
check_redis() {
    print_message $BLUE "检查Redis服务状态..."
    
    if redis-cli ping > /dev/null 2>&1; then
        print_message $GREEN "✓ Redis服务正在运行"
        return 0
    else
        print_message $YELLOW "⚠ Redis服务未运行，某些测试将被跳过"
        return 1
    fi
}

# 运行单元测试
run_unit_tests() {
    print_message $BLUE "运行单元测试..."
    
    # 运行所有单元测试
    go test -v -race -coverprofile=coverage.out ./internal/... ./pkg/...
    
    if [ $? -eq 0 ]; then
        print_message $GREEN "✓ 单元测试通过"
    else
        print_message $RED "✗ 单元测试失败"
        exit 1
    fi
}

# 运行集成测试
run_integration_tests() {
    print_message $BLUE "运行集成测试..."
    
    # 运行集成测试
    go test -v -race ./test/integration/...
    
    if [ $? -eq 0 ]; then
        print_message $GREEN "✓ 集成测试通过"
    else
        print_message $RED "✗ 集成测试失败"
        exit 1
    fi
}

# 运行基准测试
run_benchmark_tests() {
    print_message $BLUE "运行基准测试..."
    
    # 运行基准测试
    go test -bench=. -benchmem ./internal/... ./pkg/...
    
    if [ $? -eq 0 ]; then
        print_message $GREEN "✓ 基准测试完成"
    else
        print_message $RED "✗ 基准测试失败"
        exit 1
    fi
}

# 生成测试覆盖率报告
generate_coverage_report() {
    print_message $BLUE "生成测试覆盖率报告..."
    
    if [ -f coverage.out ]; then
        # 生成HTML覆盖率报告
        go tool cover -html=coverage.out -o coverage.html
        
        # 显示覆盖率统计
        go tool cover -func=coverage.out | tail -1
        
        print_message $GREEN "✓ 覆盖率报告已生成: coverage.html"
    else
        print_message $YELLOW "⚠ 未找到覆盖率文件"
    fi
}

# 运行代码质量检查
run_quality_checks() {
    print_message $BLUE "运行代码质量检查..."
    
    # 检查是否安装了golangci-lint
    if ! command -v golangci-lint &> /dev/null; then
        print_message $YELLOW "⚠ golangci-lint未安装，跳过代码质量检查"
        return
    fi
    
    # 运行golangci-lint
    golangci-lint run
    
    if [ $? -eq 0 ]; then
        print_message $GREEN "✓ 代码质量检查通过"
    else
        print_message $RED "✗ 代码质量检查失败"
        exit 1
    fi
}

# 运行性能测试
run_performance_tests() {
    print_message $BLUE "运行性能测试..."
    
    # 检查Redis是否可用
    if ! check_redis; then
        print_message $YELLOW "⚠ Redis不可用，跳过性能测试"
        return
    fi
    
    # 运行性能基准测试示例
    go run examples/performance_benchmark.go
    
    if [ $? -eq 0 ]; then
        print_message $GREEN "✓ 性能测试完成"
    else
        print_message $RED "✗ 性能测试失败"
        exit 1
    fi
}

# 运行兼容性测试
run_compatibility_tests() {
    print_message $BLUE "运行兼容性测试..."
    
    # 检查Redis是否可用
    if ! check_redis; then
        print_message $YELLOW "⚠ Redis不可用，跳过兼容性测试"
        return
    fi
    
    # 运行兼容性测试示例
    go run examples/legacy_compatibility_test.go
    
    if [ $? -eq 0 ]; then
        print_message $GREEN "✓ 兼容性测试完成"
    else
        print_message $RED "✗ 兼容性测试失败"
        exit 1
    fi
}

# 清理测试文件
cleanup() {
    print_message $BLUE "清理测试文件..."
    
    # 删除临时文件
    rm -f coverage.out
    rm -f coverage.html
    
    print_message $GREEN "✓ 清理完成"
}

# 显示帮助信息
show_help() {
    echo "测试运行脚本"
    echo ""
    echo "用法: $0 [选项]"
    echo ""
    echo "选项:"
    echo "  unit          运行单元测试"
    echo "  integration   运行集成测试"
    echo "  benchmark     运行基准测试"
    echo "  coverage      生成覆盖率报告"
    echo "  quality       运行代码质量检查"
    echo "  performance   运行性能测试"
    echo "  compatibility 运行兼容性测试"
    echo "  all           运行所有测试"
    echo "  cleanup       清理测试文件"
    echo "  help          显示此帮助信息"
    echo ""
    echo "示例:"
    echo "  $0 unit                # 只运行单元测试"
    echo "  $0 all                 # 运行所有测试"
    echo "  $0 unit integration    # 运行单元测试和集成测试"
}

# 主函数
main() {
    # 如果没有参数，显示帮助
    if [ $# -eq 0 ]; then
        show_help
        exit 0
    fi
    
    # 检查Go环境
    if ! command -v go &> /dev/null; then
        print_message $RED "✗ Go未安装或不在PATH中"
        exit 1
    fi
    
    print_message $GREEN "开始运行测试..."
    
    # 处理参数
    for arg in "$@"; do
        case $arg in
            unit)
                run_unit_tests
                ;;
            integration)
                check_redis
                run_integration_tests
                ;;
            benchmark)
                run_benchmark_tests
                ;;
            coverage)
                generate_coverage_report
                ;;
            quality)
                run_quality_checks
                ;;
            performance)
                run_performance_tests
                ;;
            compatibility)
                run_compatibility_tests
                ;;
            all)
                check_redis
                run_unit_tests
                run_integration_tests
                run_benchmark_tests
                generate_coverage_report
                run_quality_checks
                run_performance_tests
                run_compatibility_tests
                ;;
            cleanup)
                cleanup
                ;;
            help)
                show_help
                ;;
            *)
                print_message $RED "未知选项: $arg"
                show_help
                exit 1
                ;;
        esac
    done
    
    print_message $GREEN "所有测试完成！"
}

# 运行主函数
main "$@"