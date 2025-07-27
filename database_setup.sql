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
    username VARCHAR(80) UNIQUE,
    context VARCHAR(800)
);

-- 5. 显示创建结果
SELECT '数据库 wutonkdb 和表 user_context 创建成功！' as message; 