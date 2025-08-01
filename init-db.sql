-- 创建符合要求的表
CREATE TABLE IF NOT EXISTS user_context (
    username VARCHAR(80) UNIQUE,
    context VARCHAR(800)
); 