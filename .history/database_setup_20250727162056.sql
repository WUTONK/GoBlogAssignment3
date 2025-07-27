-- PostgreSQL 数据库设置脚本
-- 数据库名: wutonkdb
-- 表名: user_context
-- 用户: postgres
-- 密码: 123456

-- 1. 连接到默认数据库 postgres
\c postgres;

-- 2. 创建数据库（如果不存在）
CREATE DATABASE wutonkdb;

-- 3. 连接到新创建的数据库
\c wutonkdb;

-- 4. 创建用户上下文表
CREATE TABLE IF NOT EXISTS user_context (
    username VARCHAR(80),
    context VARCHAR(800)
);

-- 5. 创建索引以提高查询性能
CREATE INDEX IF NOT EXISTS idx_user_context_username ON user_context(user_name);
CREATE INDEX IF NOT EXISTS idx_user_context_created_at ON user_context(created_at);

-- 6. 创建更新时间触发器
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_user_context_updated_at 
    BEFORE UPDATE ON user_context 
    FOR EACH ROW 
    EXECUTE FUNCTION update_updated_at_column();

-- 7. 插入一些测试数据（可选）
INSERT INTO user_context (user_name, context) VALUES 
    ('test_user', '这是测试用户的初始上下文'),
    ('admin', '管理员用户的上下文');

-- 8. 显示创建结果
SELECT '数据库 wutonkdb 和表 user_context 创建成功！' as message;
SELECT COUNT(*) as table_count FROM user_context; 