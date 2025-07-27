# 使用须知

本软件需要的数据库表格式为：username(varchar 80); context(varchar80)
默认数据库名为: wutonkdb 
表名为:user_context
密码为:123456 用户为:postgres

你可以复制以下脚本到你的psql中来创建一个符合要求的数据库
```sh
CREATE DATABASE wutonkdb;
```

然后创建一个符合要求的表
```sh
CREATE TABLE IF NOT EXISTS user_context (
    username VARCHAR(80) UNIQUE,
    context  VARCHAR(800),
);
```
