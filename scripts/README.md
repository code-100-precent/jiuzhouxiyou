# 初始化脚本说明

## 初始化测试用户

运行以下命令创建初始测试账号：

```bash
cd backend
go run scripts/init_users.go
```

或者：

```bash
cd backend/scripts
go run init_users.go
```

### 默认创建的账号

| 账号 | 密码 | 昵称 |
|------|------|------|
| test | 123456 | 测试用户 |
| wukong | 123456 | 孙悟空 |
| bajie | 123456 | 猪八戒 |
| wujing | 123456 | 沙悟净 |
| tangseng | 123456 | 唐僧 |

所有账号的默认头像都是：`https://cetide-1325039295.cos.ap-chengdu.myqcloud.com/west/default_avatar01.png`

## 注意事项

- 如果账号已存在，脚本会跳过创建
- 密码使用 bcrypt 加密存储
- 可以修改 `init_users.go` 中的用户列表来添加更多测试账号

