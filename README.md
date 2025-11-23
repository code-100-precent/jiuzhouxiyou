# 九州西游 - 后端服务

基于 Go + Gin + GORM + MySQL 的游戏后端服务。

## 功能特性

- ✅ 用户登录认证（JWT）
- ✅ 心跳检测
- ✅ 存档管理
- ✅ CORS 支持

## 技术栈

- **框架**: Gin
- **ORM**: GORM
- **数据库**: MySQL
- **认证**: JWT
- **密码加密**: bcrypt

## 项目结构

```
backend/
├── config/          # 配置管理
├── database/        # 数据库连接
├── handlers/        # 请求处理器
├── middleware/      # 中间件
├── models/          # 数据模型
├── routes/          # 路由定义
├── utils/           # 工具函数
├── .env            # 环境配置
├── go.mod          # Go 模块
└── main.go         # 入口文件
```

## 快速开始

### 1. 安装依赖

```bash
cd backend
go mod download
```

### 2. 配置环境变量

编辑 `.env` 文件：

```env
DB_DRIVER=mysql
DSN=root:password@tcp(localhost:3306)/lingecho_db?charset=utf8mb4&parseTime=True&loc=Local
PORT=7070
JWT_SECRET=your-secret-key-change-in-production
```

### 3. 运行服务

**方式一：直接运行**
```bash
go run main.go
```

**方式二：使用启动脚本**
```bash
./start.sh
```

**方式三：使用 Makefile**
```bash
make run
```

服务器将在 `http://localhost:7070` 启动。

### 4. 创建测试用户

服务启动后会自动创建数据表。你可以通过以下方式创建测试用户：

**方式一：使用 SQL 脚本**
```bash
mysql -h cd-cynosdbmysql-grp-lfa6zfg0.sql.tencentcdb.com -P 23771 -u root -p lingecho_db < scripts/init_test_user.sql
```

**方式二：手动插入（密码: 123456）**
```sql
INSERT INTO users (uid, password, name, avatar, created_at, updated_at) 
VALUES ('test', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', '测试用户', '', NOW(), NOW());
```

**方式三：使用 Go 代码生成密码哈希**
创建一个临时文件生成密码哈希：
```go
package main
import (
    "fmt"
    "golang.org/x/crypto/bcrypt"
)
func main() {
    hash, _ := bcrypt.GenerateFromPassword([]byte("your-password"), bcrypt.DefaultCost)
    fmt.Println(string(hash))
}
```

## API 文档

### 1. 用户登录

**POST** `/api/v1/users/login`

**请求体:**
```json
{
  "uid": "username",
  "pwd": "password"
}
```

**响应:**
```json
{
  "code": 200,
  "msg": "登录成功",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }
}
```

### 2. 心跳检测

**POST** `/api/v1/users/ping`

**Headers:**
```
Authorization: Bearer <token>
```

**响应:**
```json
{
  "code": 200,
  "msg": "pong",
  "data": {
    "user_id": 1,
    "uid": "username"
  }
}
```

### 3. 获取存档列表

**GET** `/api/v1/users/archive`

**Headers:**
```
Authorization: Bearer <token>
```

**响应:**
```json
{
  "code": 200,
  "msg": "获取成功",
  "data": {
    "archives": [
      {
        "id": 1,
        "user_id": 1,
        "name": "角色名",
        "job": "职业",
        "avatar": "头像URL",
        "last_login_time": "2025-01-01T12:00:00Z",
        "created_at": "2025-01-01T10:00:00Z",
        "updated_at": "2025-01-01T12:00:00Z"
      }
    ]
  }
}
```

## 数据库表结构

### users 表
- `id`: 主键
- `uid`: 用户名（唯一）
- `password`: 加密后的密码
- `name`: 用户昵称
- `avatar`: 头像URL
- `created_at`: 创建时间
- `updated_at`: 更新时间

### archives 表
- `id`: 主键
- `user_id`: 用户ID（外键）
- `name`: 角色名
- `job`: 职业
- `avatar`: 头像URL
- `last_login_time`: 最后登录时间
- `created_at`: 创建时间
- `updated_at`: 更新时间

## 开发说明

### 创建测试用户

可以通过数据库直接插入，或使用以下 SQL：

```sql
-- 密码是 "123456" 的 bcrypt 哈希
INSERT INTO users (uid, password, name, avatar, created_at, updated_at) 
VALUES ('test', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', '测试用户', '', NOW(), NOW());
```

### 编译

```bash
go build -o server main.go
```

### 运行编译后的程序

```bash
./server
```

## 注意事项

1. 生产环境请修改 `.env` 中的 `JWT_SECRET`
2. 确保数据库连接信息正确
3. 首次运行会自动创建数据表
4. 密码使用 bcrypt 加密存储

## License

MIT

# jiuzhouxiyou
