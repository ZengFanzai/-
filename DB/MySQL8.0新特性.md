# MySQL8.0 新特性

1. 默认字符集有 latin1 变为 utf8
2. MyISAM 系统表变为 InnoDB 表
3. 变量自增持久性
   1. 在 8.0 之前的版本，自增主键 AUTO_INCREMENT 的值如果大于 max(primary key)+1，在 MySQL 重启后，会重置 AUTO_INCREMENT=max(primary key)+1，这种现象在某些情况下会导致业务主键冲突或者其他难以发现的问题。自增主键重启重置的问题很早就被发现[bug 链接](https://bugs.mysql.com/bug.php?id=199)，一直到 8.0 才被解决，8.0 版本将会对 AUTO_INCREMENT 值进行持久化，MySQL 重启后，该值将不会改变。
4. DDL 原子化：InnoDB 表的 DDL 支持事务完整性，要么成功要么回滚
5. 参数修改持久化
   1. MySQL 8.0 版本支持在线修改全局参数并持久化，通过加上 PERSIST 关键字，可以将修改的参数持久化到新的配置文件（mysqld-auto.cnf）中，重启 MySQL 时，可以从该配置文件获取到最新的配置参数。
   2. `set PERSIST expire_logs_days=10 ;`
6. 新增降序索引
7. group by 不再隐式排序，如需要排序，必须显式加上 order by 子句。
8. JSON 特性增强
9. redo & undo 日志加密，参数如下
   1. innodb_undo_log_encrypt
   2. innodb_redo_log_encrypt
10. innodb select for update 跳过锁等待
11. 增加 SET_VAR 语法，在 sql 语法中增加 SET_VAR 语法，动态调整部分参数，有利于提升语句性能。
    1. `select /_+ SET_VAR(sort_buffer_size = 16M) _/ id from test order id ;`
    2. `insert /_+ SET_VAR(foreign_key_checks=OFF) _/ into test(name) values(1);`
12. 支持不可见索引
13. 支持直方图
14. 新增 innodb_dedicated_server 参数，能够让 InnoDB 根据服务器上检测到的内存大小自动配置 innodb_buffer_pool_size，innodb_log_file_size，innodb_flush_method 三个参数。
15. 日志分类更详细
16. undo 空间自动回收
17. 增加资源组
18. 增加角色管理

**参考**

> - [MySQL8.0 新特性集锦](jianshu.com/p/be29467c2b0c)
