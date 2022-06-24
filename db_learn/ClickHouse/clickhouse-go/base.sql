# 创建数据库
create database statics engine=Ordinary;

# 切换数据库
use statics;

# 创建表
create table mysql_slow_log (
    id UInt16,
    user_name String,
    host String,
    sql String,
    rows_examined UInt16,
    exec_time UInt16,
    query_time String,
    create_time date
)engine=MergeTree(create_time, (id), 8192);

# 插入数据
insert into mysql_slow_log  (user_name, host, sql, rows_examined, exec_time, query_time, create_time) values('xiaoming', '127.0.0.1', 'select * from music', 3000, 1587021607, '0.333', '2020-04-16 15:32:17');
