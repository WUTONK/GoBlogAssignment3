# 使用须知

本软件需要的数据库表格式为：username(varchar 80); context(varchar80)    

标准表名为:user_context

以下是后端链接表的标准信息 可参考: "user=postgres password=123456 host=localhost port=5432 dbname=wutonkdb sslmode=disable"

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
