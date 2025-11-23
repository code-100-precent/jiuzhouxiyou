#!/bin/bash

echo "启动九州西游后端服务..."
echo ""

# 检查 .env 文件
if [ ! -f .env ]; then
    echo "错误: .env 文件不存在"
    echo "请复制 .env.example 并修改配置:"
    echo "  cp .env.example .env"
    exit 1
fi

# 加载环境变量
export $(cat .env | grep -v '^#' | xargs)

# 运行服务
go run main.go

