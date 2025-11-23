-- 初始化测试用户
-- 密码: 123456 (bcrypt 哈希值)

-- 注意：这个哈希值对应密码 "123456"
-- 如果需要其他密码，可以使用 Go 的 bcrypt 生成

INSERT INTO users (uid, password, name, avatar, created_at, updated_at) 
VALUES ('test', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', '测试用户', 'https://cetide-1325039295.cos.ap-chengdu.myqcloud.com/west/default_avatar01.png', NOW(), NOW())
ON DUPLICATE KEY UPDATE updated_at = NOW();

-- 创建测试存档
INSERT INTO archives (user_id, name, job, avatar, last_login_time, created_at, updated_at)
SELECT id, '孙悟空', '战士', 'https://cetide-1325039295.cos.ap-chengdu.myqcloud.com/west/default_avatar01.png', NOW(), NOW(), NOW()
FROM users WHERE uid = 'test'
LIMIT 1;

