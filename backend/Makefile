.PHONY: run build test clean deps

# 运行服务
run:
	go run main.go

# 构建
build:
	go build -o bin/server main.go

# 测试
test:
	go test ./...

# 清理
clean:
	rm -rf bin/

# 安装依赖
deps:
	go mod download
	go mod tidy

# 初始化数据库（需要先配置 .env）
init-db:
	@echo "请确保数据库已创建，然后运行服务会自动迁移表结构"

