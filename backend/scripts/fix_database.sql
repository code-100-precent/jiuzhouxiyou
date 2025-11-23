-- 修复数据库迁移问题
-- 清理无效数据并修复表结构

-- 1. 删除 uid 为空或 NULL 的记录
DELETE FROM users WHERE uid = '' OR uid IS NULL;

-- 2. 处理重复的 uid（保留 id 最小的记录）
DELETE u1 FROM users u1
INNER JOIN users u2 
WHERE u1.id > u2.id AND u1.uid = u2.uid AND u1.uid != '';

-- 3. 如果 uid 列不存在唯一索引，先删除可能存在的索引
-- ALTER TABLE users DROP INDEX IF EXISTS uid;

-- 4. 确保 uid 列有正确的约束（如果表已存在但结构不对）
-- 注意：如果表结构已经正确，这些语句不会执行
-- 如果表不存在，GORM 会自动创建

