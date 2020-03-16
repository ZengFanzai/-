# MySQL 面试题收集 1

- [MySQL 面试题收集 1](#mysql-%e9%9d%a2%e8%af%95%e9%a2%98%e6%94%b6%e9%9b%86-1)
  - [MySQL 的复制原理以及流程](#mysql-%e7%9a%84%e5%a4%8d%e5%88%b6%e5%8e%9f%e7%90%86%e4%bb%a5%e5%8f%8a%e6%b5%81%e7%a8%8b)
  - [MySQL 中 varchar 与 char 的区别以及 varchar(50)中的 50 代表的涵义](#mysql-%e4%b8%ad-varchar-%e4%b8%8e-char-%e7%9a%84%e5%8c%ba%e5%88%ab%e4%bb%a5%e5%8f%8a-varchar50%e4%b8%ad%e7%9a%84-50-%e4%bb%a3%e8%a1%a8%e7%9a%84%e6%b6%b5%e4%b9%89)
    - [varchar 与 char 的区别](#varchar-%e4%b8%8e-char-%e7%9a%84%e5%8c%ba%e5%88%ab)
    - [varchar(50)中 50 的涵义](#varchar50%e4%b8%ad-50-%e7%9a%84%e6%b6%b5%e4%b9%89)
    - [int(20)中 20 的涵义](#int20%e4%b8%ad-20-%e7%9a%84%e6%b6%b5%e4%b9%89)
    - [mysql 为什么这么设计](#mysql-%e4%b8%ba%e4%bb%80%e4%b9%88%e8%bf%99%e4%b9%88%e8%ae%be%e8%ae%a1)
  - [MySQL binlog 的几种日志录入格式以及区别](#mysql-binlog-%e7%9a%84%e5%87%a0%e7%a7%8d%e6%97%a5%e5%bf%97%e5%bd%95%e5%85%a5%e6%a0%bc%e5%bc%8f%e4%bb%a5%e5%8f%8a%e5%8c%ba%e5%88%ab)
    - [Statement](#statement)
    - [Row](#row)
    - [Mixedlevel](#mixedlevel)
  - [MySQL 数据库 cpu 飙升到 500%的话他怎么处理?](#mysql-%e6%95%b0%e6%8d%ae%e5%ba%93-cpu-%e9%a3%99%e5%8d%87%e5%88%b0-500%e7%9a%84%e8%af%9d%e4%bb%96%e6%80%8e%e4%b9%88%e5%a4%84%e7%90%86)
  - [sql 优化各种方法](#sql-%e4%bc%98%e5%8c%96%e5%90%84%e7%a7%8d%e6%96%b9%e6%b3%95)
    - [explain 出来的各种 item 意义](#explain-%e5%87%ba%e6%9d%a5%e7%9a%84%e5%90%84%e7%a7%8d-item-%e6%84%8f%e4%b9%89)
    - [profile 的意义以及使用场景](#profile-%e7%9a%84%e6%84%8f%e4%b9%89%e4%bb%a5%e5%8f%8a%e4%bd%bf%e7%94%a8%e5%9c%ba%e6%99%af)
    - [MySQL5.7 之后 profile 为过时命令,使用 performance_schema 来替代](#mysql57-%e4%b9%8b%e5%90%8e-profile-%e4%b8%ba%e8%bf%87%e6%97%b6%e5%91%bd%e4%bb%a4%e4%bd%bf%e7%94%a8-performanceschema-%e6%9d%a5%e6%9b%bf%e4%bb%a3)
  - [备份计划,mysqldump 以及 xtranbackup 的实现原理](#%e5%a4%87%e4%bb%bd%e8%ae%a1%e5%88%92mysqldump-%e4%bb%a5%e5%8f%8a-xtranbackup-%e7%9a%84%e5%ae%9e%e7%8e%b0%e5%8e%9f%e7%90%86)
    - [备份计划](#%e5%a4%87%e4%bb%bd%e8%ae%a1%e5%88%92)
    - [备份恢复时间](#%e5%a4%87%e4%bb%bd%e6%81%a2%e5%a4%8d%e6%97%b6%e9%97%b4)
    - [xtrabackup 实现原理](#xtrabackup-%e5%ae%9e%e7%8e%b0%e5%8e%9f%e7%90%86)
  - [mysqldump 中备份出来的 sql,如果我想 sql 文件中,一行只有一个 insert...value()的话,怎么办?如果备份需要带上 master 的复制点信息怎么办?](#mysqldump-%e4%b8%ad%e5%a4%87%e4%bb%bd%e5%87%ba%e6%9d%a5%e7%9a%84-sql%e5%a6%82%e6%9e%9c%e6%88%91%e6%83%b3-sql-%e6%96%87%e4%bb%b6%e4%b8%ad%e4%b8%80%e8%a1%8c%e5%8f%aa%e6%9c%89%e4%b8%80%e4%b8%aa-insertvalue%e7%9a%84%e8%af%9d%e6%80%8e%e4%b9%88%e5%8a%9e%e5%a6%82%e6%9e%9c%e5%a4%87%e4%bb%bd%e9%9c%80%e8%a6%81%e5%b8%a6%e4%b8%8a-master-%e7%9a%84%e5%a4%8d%e5%88%b6%e7%82%b9%e4%bf%a1%e6%81%af%e6%80%8e%e4%b9%88%e5%8a%9e)
  - [500 台 db,在最快时间之内重启](#500-%e5%8f%b0-db%e5%9c%a8%e6%9c%80%e5%bf%ab%e6%97%b6%e9%97%b4%e4%b9%8b%e5%86%85%e9%87%8d%e5%90%af)
  - [innodb 的读写参数优化](#innodb-%e7%9a%84%e8%af%bb%e5%86%99%e5%8f%82%e6%95%b0%e4%bc%98%e5%8c%96)
    - [读取参数](#%e8%af%bb%e5%8f%96%e5%8f%82%e6%95%b0)
    - [写入参数](#%e5%86%99%e5%85%a5%e5%8f%82%e6%95%b0)
    - [与 IO 相关的参数](#%e4%b8%8e-io-%e7%9b%b8%e5%85%b3%e7%9a%84%e5%8f%82%e6%95%b0)
    - [缓存参数以及缓存的适用场景](#%e7%bc%93%e5%ad%98%e5%8f%82%e6%95%b0%e4%bb%a5%e5%8f%8a%e7%bc%93%e5%ad%98%e7%9a%84%e9%80%82%e7%94%a8%e5%9c%ba%e6%99%af)
  - [你是如何监控你们的数据库的?你们的慢日志都是怎么查询的?](#%e4%bd%a0%e6%98%af%e5%a6%82%e4%bd%95%e7%9b%91%e6%8e%a7%e4%bd%a0%e4%bb%ac%e7%9a%84%e6%95%b0%e6%8d%ae%e5%ba%93%e7%9a%84%e4%bd%a0%e4%bb%ac%e7%9a%84%e6%85%a2%e6%97%a5%e5%bf%97%e9%83%bd%e6%98%af%e6%80%8e%e4%b9%88%e6%9f%a5%e8%af%a2%e7%9a%84)
  - [你是否做过主从一致性校验,如果有,怎么做的,如果没有,你打算怎么做？](#%e4%bd%a0%e6%98%af%e5%90%a6%e5%81%9a%e8%bf%87%e4%b8%bb%e4%bb%8e%e4%b8%80%e8%87%b4%e6%80%a7%e6%a0%a1%e9%aa%8c%e5%a6%82%e6%9e%9c%e6%9c%89%e6%80%8e%e4%b9%88%e5%81%9a%e7%9a%84%e5%a6%82%e6%9e%9c%e6%b2%a1%e6%9c%89%e4%bd%a0%e6%89%93%e7%ae%97%e6%80%8e%e4%b9%88%e5%81%9a)
  - [你们数据库是否支持 emoji 表情,如果不支持,如何操作?](#%e4%bd%a0%e4%bb%ac%e6%95%b0%e6%8d%ae%e5%ba%93%e6%98%af%e5%90%a6%e6%94%af%e6%8c%81-emoji-%e8%a1%a8%e6%83%85%e5%a6%82%e6%9e%9c%e4%b8%8d%e6%94%af%e6%8c%81%e5%a6%82%e4%bd%95%e6%93%8d%e4%bd%9c)
  - [你是如何维护数据库的数据字典的?](#%e4%bd%a0%e6%98%af%e5%a6%82%e4%bd%95%e7%bb%b4%e6%8a%a4%e6%95%b0%e6%8d%ae%e5%ba%93%e7%9a%84%e6%95%b0%e6%8d%ae%e5%ad%97%e5%85%b8%e7%9a%84)
  - [表中有大字段 X(例如:text 类型),且字段 X 不会经常更新,以读为为主,请问是否应该拆分](#%e8%a1%a8%e4%b8%ad%e6%9c%89%e5%a4%a7%e5%ad%97%e6%ae%b5-x%e4%be%8b%e5%a6%82text-%e7%b1%bb%e5%9e%8b%e4%b8%94%e5%ad%97%e6%ae%b5-x-%e4%b8%8d%e4%bc%9a%e7%bb%8f%e5%b8%b8%e6%9b%b4%e6%96%b0%e4%bb%a5%e8%af%bb%e4%b8%ba%e4%b8%ba%e4%b8%bb%e8%af%b7%e9%97%ae%e6%98%af%e5%90%a6%e5%ba%94%e8%af%a5%e6%8b%86%e5%88%86)
  - [MySQL 中 InnoDB 引擎的行锁是通过加在什么上完成(或称实现)的?为什么是这样子的?](#mysql-%e4%b8%ad-innodb-%e5%bc%95%e6%93%8e%e7%9a%84%e8%a1%8c%e9%94%81%e6%98%af%e9%80%9a%e8%bf%87%e5%8a%a0%e5%9c%a8%e4%bb%80%e4%b9%88%e4%b8%8a%e5%ae%8c%e6%88%90%e6%88%96%e7%a7%b0%e5%ae%9e%e7%8e%b0%e7%9a%84%e4%b8%ba%e4%bb%80%e4%b9%88%e6%98%af%e8%bf%99%e6%a0%b7%e5%ad%90%e7%9a%84)
  - [开放性问题:据说是腾讯的](#%e5%bc%80%e6%94%be%e6%80%a7%e9%97%ae%e9%a2%98%e6%8d%ae%e8%af%b4%e6%98%af%e8%85%be%e8%ae%af%e7%9a%84)
  - [索引是什么?有什么作用以及优缺点?](#%e7%b4%a2%e5%bc%95%e6%98%af%e4%bb%80%e4%b9%88%e6%9c%89%e4%bb%80%e4%b9%88%e4%bd%9c%e7%94%a8%e4%bb%a5%e5%8f%8a%e4%bc%98%e7%bc%ba%e7%82%b9)
  - [什么是事务?](#%e4%bb%80%e4%b9%88%e6%98%af%e4%ba%8b%e5%8a%a1)
  - [使用索引一定能提高查询的性能吗?为什么](#%e4%bd%bf%e7%94%a8%e7%b4%a2%e5%bc%95%e4%b8%80%e5%ae%9a%e8%83%bd%e6%8f%90%e9%ab%98%e6%9f%a5%e8%af%a2%e7%9a%84%e6%80%a7%e8%83%bd%e5%90%97%e4%b8%ba%e4%bb%80%e4%b9%88)
  - [简单说一下 drop,delete,truncate 的区别](#%e7%ae%80%e5%8d%95%e8%af%b4%e4%b8%80%e4%b8%8b-dropdeletetruncate-%e7%9a%84%e5%8c%ba%e5%88%ab)
  - [drop,delete,truncate 分别在什么场景下使用?](#dropdeletetruncate-%e5%88%86%e5%88%ab%e5%9c%a8%e4%bb%80%e4%b9%88%e5%9c%ba%e6%99%af%e4%b8%8b%e4%bd%bf%e7%94%a8)
  - [超键,候选键,主键,外键分别是什么?](#%e8%b6%85%e9%94%ae%e5%80%99%e9%80%89%e9%94%ae%e4%b8%bb%e9%94%ae%e5%a4%96%e9%94%ae%e5%88%86%e5%88%ab%e6%98%af%e4%bb%80%e4%b9%88)
  - [什么是视图?以及视图的使用场景有哪些?](#%e4%bb%80%e4%b9%88%e6%98%af%e8%a7%86%e5%9b%be%e4%bb%a5%e5%8f%8a%e8%a7%86%e5%9b%be%e7%9a%84%e4%bd%bf%e7%94%a8%e5%9c%ba%e6%99%af%e6%9c%89%e5%93%aa%e4%ba%9b)
  - [参考](#%e5%8f%82%e8%80%83)

> 摘自[《史上最详细的一线大厂 Mysql 面试题详解》](https://juejin.im/post/5cb6c4ef51882532b70e6ff0#heading-0)

## MySQL 的复制原理以及流程

基本原理流程，3 个线程以及之间的关联：

- **主**：binlog 线程——记录下所有改变了数据库的语句，放进 master 上的 binlog 中；
- **从**：io 线程——在使用 start slave 之后，负责从 master 上拉取 binlog 内容，放进自己的 relay log 中
- **从**：sql 执行线程——执行 relay log 中的语句；

## MySQL 中 varchar 与 char 的区别以及 varchar(50)中的 50 代表的涵义

### varchar 与 char 的区别

- char 是一种固定长度的类型；
- varchar 则是一种可变长度的类型

### varchar(50)中 50 的涵义

最多存放 50 个字符，varchar(50)和(200)存储 hello 所占空间一样，但后者在排序时会消耗更多内存，因为 order by col 采用 fixed_length 计算 col 长度(memory 引擎也一样)

### int(20)中 20 的涵义

是指显示字符的长度,20 表示最大显示宽度为 20，但仍占 4 字节存储，存储范围不变；比如它是记录行数的 id,插入 10 笔资料，它就显示 00000000001~~~00000000010，当字符的位数超过 11,它也只显示 11 位，如果你没有加那个让它未满 11 位就前面加 0 的参数，它不会在前面加 0

### mysql 为什么这么设计

对大多数应用没有意义，只是规定一些工具用来显示字符的个数；int(1)和 int(20)存储和计算均一样；

## MySQL binlog 的几种日志录入格式以及区别

### Statement

每一条会修改数据的 sql 都会记录在 binlog 中。

**优点**：**不需要记录每一行的变化，减少了 binlog 日志量，节约了 IO，提高性能**。(相比 row 能节约多少性能 与日志量，这个取决于应用的 SQL 情况，正常同一条记录修改或者插入 row 格式所产生的日志量还小于 Statement 产生的日志量，但是考虑到如果带条 件的 update 操作，以及整表删除，alter 表等操作，ROW 格式会产生大量日志，因此在考虑是否使用 ROW 格式日志时应该跟据应用的实际情况，其所 产生的日志量会增加多少，以及带来的 IO 性能问题。)

**缺点**：由于记录的只是执行语句，为了这些语句能在 slave 上正确运行，因此还必须记录每条语句在执行的时候的 一些相关信息，以保证所有语句能在 slave 得到和在 master 端执行时候相同 的结果。另外 mysql 的复制,像一些特定函数功能，slave 可与 master 上要保持一致会有很多相关问题(如 sleep()函数， last_insert_id()，以及 user-defined functions(udf)会出现问题)。

**使用以下函数的语句也无法被复制**：

- LOAD_FILE()
- UUID()
- USER()
- FOUND_ROWS()
- SYSDATE() (除非启动时启用了 --sysdate-is-now 选项)

同时在 INSERT …SELECT 会产生比 RBR 更多的行级锁

### Row

不记录 sql 语句上下文相关信息，仅保存哪条记录被修改

**优点**： binlog 中可以不记录执行的 sql 语句的上下文相关的信息，仅需要记录那一条记录被修改成什么了。所以 rowlevel 的日志内容会非常清楚的记录下 每一行数据修改的细节。而且不会出现某些特定情况下的存储过程，或 function，以及 trigger 的调用和触发无法被正确复制的问题。

**缺点**：所有的执行的语句当记录到日志中的时候，都将以每行记录的修改来记录，这样可能会产生大量的日志内容,比 如一条 update 语句，修改多条记录，则 binlog 中每一条修改都会有记录，这样造成 binlog 日志量会很大，特别是当执行 alter table 之类的语句的时候，由于表结构修改，每条记录都发生改变，那么该表每一条记录都会记录到日志中。

### Mixedlevel

以上两种 level 的混合使用，一般的语句修改使用 statment 格式保存 binlog，如一些函数，statement 无法完成主从复制的操作，则 采用 row 格式保存 binlog,MySQL 会根据执行的每一条具体的 sql 语句来区分对待记录的日志形式，也就是在 Statement 和 Row 之间选择 一种.新版本的 MySQL 中队 row level 模式也被做了优化，并不是所有的修改都会以 row level 来记录，像遇到表结构变更的时候就会以 statement 模式来记录。至于 update 或者 delete 等修改数据的语句，还是会记录所有行的变更。

## MySQL 数据库 cpu 飙升到 500%的话他怎么处理?

- 列出所有进程 `show processlist`,观察所有进程 ,多秒没有状态变化的(干掉)
- 查看超时日志或者错误日志 (做了几年开发,一般会是查询以及大批量的插入会导致 cpu 与 i/o 上涨,当然不排除网络状态突然断了,导致一个请求服务器只接受到一半，比如 where 子句或分页子句没有发送,,当然的一次被坑经历)

## sql 优化各种方法

### explain 出来的各种 item 意义

- select_type：表示查询中每个 select 子句的类型
- type：表示 MySQL 在表中找到所需行的方式，又称“访问类型”
- possible_keys：指出 MySQL 能使用哪个索引在表中找到行，查询涉及到的字段上若存在索引，则该索引将被列出，但不一定被查询使用
- key：显示 MySQL 在查询中实际用到的索引，若没有使用索引，显示为 NULL
- key_len：表示索引中使用的字节数，可以通过该列计算查询中使用的索引长度
- ref：表示上述表的连接匹配条件，即哪些或常量被用于查询索引列上的值
- Extra：包含不适合在其他列中显示但十分重要的额外信息

### profile 的意义以及使用场景

查询到 SQL 会执行多少时间, 并看出 CPU/Memory 使用量, 执行过程中 Systemlock, Table lock 花多少时间等等

### MySQL5.7 之后 profile 为过时命令,使用 performance_schema 来替代

## 备份计划,mysqldump 以及 xtranbackup 的实现原理

### 备份计划

按业务实际情况备份

### 备份恢复时间

这里跟机器，尤其是硬盘的速率有关系，以下列举几个仅供参考

- 20G 的 2 分钟（mysqldump）
- 80G 的 30 分钟(mysqldump)
- 111G 的 30 分钟（mysqldump)
- 288G 的 3 小时（xtra)
- 3T 的 4 小时（xtra)

逻辑导入时间一般是备份时间的 5 倍以上

### xtrabackup 实现原理

在 InnoDB 内部会维护一个 redo 日志文件，我们也可以叫做事务日志文件。事务日志会存储每一个 InnoDB 表数据的记录修改。当 InnoDB 启动时，InnoDB 会检查数据文件和事务日志，并执行两个步骤：它应用（前滚）已经提交的事务日志到数据文件，并将修改过但没有提交的数据进行回滚操作。

## mysqldump 中备份出来的 sql,如果我想 sql 文件中,一行只有一个 insert...value()的话,怎么办?如果备份需要带上 master 的复制点信息怎么办?

```sql
--skip-extended-insert

[root@helei-zhuanshu ~]# mysqldump -uroot -p helei --skip-extended-insert

Enter password:

KEY `idx_c1` (`c1`),

KEY `idx_c2` (`c2`)

) ENGINE=InnoDB AUTO_INCREMENT=51 DEFAULT CHARSET=latin1;

/*!40101 SET character_set_client = @saved_cs_client */;

--

-- Dumping data for table `helei`

--

LOCK TABLES `helei` WRITE;

/*!40000 ALTER TABLE `helei` DISABLE KEYS */;

INSERT INTO `helei` VALUES (1,32,37,38,'2016-10-18 06:19:24','susususususususususususu');

INSERT INTO `helei` VALUES (2,37,46,21,'2016-10-18 06:19:24','susususususu');

INSERT INTO `helei` VALUES (3,21,5,14,'2016-10-18 06:19:24','susu');
```

## 500 台 db,在最快时间之内重启

可以使用批量 ssh 工具 pssh 来对需要重启的机器执行重启命令。 也可以使用 salt（前提是客户端有安装 salt）或者 ansible（ ansible 只需要 ssh 免登通了就行）等多线程工具同时操作多台服务器

## innodb 的读写参数优化

### 读取参数

global buffer 以及 local buffer；

```text
# Global buffer：
Innodb_buffer_pool_size //理论上越大越好，建议服务器50%~80%，实际为数据大小80%~90%即可；
innodb_log_buffer_size
innodb_additional_mem_pool_size

# local buffer(下面的都是 server 层的 session 变量，不是 innodb 的)：
Read_buffer_size
Join_buffer_size
Sort_buffer_size
Key_buffer_size
Binlog_cache_size
```

### 写入参数

```text
innodb_flush_log_at_trx_commit
innodb_buffer_pool_size
insert_buffer_size
innodb_double_write
innodb_write_io_thread
innodb_flush_method
```

### 与 IO 相关的参数

```text
innodb_write_io_threads = 8
innodb_read_io_threads = 8 //根据处理器内核数决定；
innodb_thread_concurrency = 0
Sync_binlog
Innodb_flush_log_at_trx_commit
Innodb_lru_scan_depth
Innodb_io_capacity
Innodb_io_capacity_max
innodb_log_buffer_size
innodb_max_dirty_pages_pct
```

### 缓存参数以及缓存的适用场景

```text
query cache/query_cache_type
```

并不是所有表都适合使用 query cache。造成 query cache 失效的原因主要是相应的 table 发生了变更

- 第一个：读操作多的话看看比例，简单来说，如果是用户清单表，或者说是数据比例比较固定，比如说商品列表，是可以打开的，前提是这些库比较集中，数据库中的实务比较小。

- 第二个：我们“行骗”的时候，比如说我们竞标的时候压测，把 query cache 打开，还是能收到 qps 激增的效果，当然前提示前端的连接池什么的都配置一样。大部分情况下如果写入的居多，访问量并不多，那么就不要打开，例如社交网站的，10%的人产生内容，其余的 90%都在消费，打开还是效果很好的，但是你如果是 qq 消息，或者聊天，那就很要命。

- 第三个：小网站或者没有高并发的无所谓，高并发下，会看到 很多 qcache 锁 等待，所以一般高并发下，不建议打开 query cache

## 你是如何监控你们的数据库的?你们的慢日志都是怎么查询的?

监控的工具有很多，例如 zabbix，lepus，我这里用的是 lepus

## 你是否做过主从一致性校验,如果有,怎么做的,如果没有,你打算怎么做？

主从一致性校验有多种工具 例如 checksum、mysqldiff、pt-table-checksum 等

## 你们数据库是否支持 emoji 表情,如果不支持,如何操作?

如果是 utf8 字符集的话，需要升级至 utf8_mb4 方可支持(四字节)

一个汉字 utf8 占 3 字节

## 你是如何维护数据库的数据字典的?

这个大家维护的方法都不同，我一般是直接在生产库进行注释，利用工具导出成 excel 方便流通。

## 表中有大字段 X(例如:text 类型),且字段 X 不会经常更新,以读为为主,请问是否应该拆分

拆带来的问题：连接消耗 + 存储拆分空间；不拆可能带来的问题：查询性能；

1. 如果能容忍拆分带来的空间问题,拆的话最好和经常要查询的表的主键在物理结构上放置在一起(分区) 顺序 IO,减少连接消耗,最后这是一个文本列再加上一个全文索引来尽量抵消连接消耗
2. 如果能容忍不拆分带来的查询性能损失的话:上面的方案在某个极致条件下肯定会出现问题,那么不拆就是最好的选择

## MySQL 中 InnoDB 引擎的行锁是通过加在什么上完成(或称实现)的?为什么是这样子的?

**InnoDB 是基于索引来完成行锁**

例: `select * from tab_with_index where id = 1 for update; for update` 可以根据条件来完成行锁锁定,并且 id 是有索引键的列,如果 id 不是索引键那么 InnoDB 将完成表锁,,并发将无从谈起

innodb 是将 primary key index 和相关的行数据共同放在 B+树的叶节点；innodb 一定会有一个 primary key，secondary index 查找的时候，也是通过找到对应的 primary，再找对应的数据行；

## 开放性问题:据说是腾讯的

一个 6 亿的表 a，一个 3 亿的表 b，通过外间 tid 关联，你如何最快的查询出满足条件的第 50000 到第 50200 中的这 200 条数据记录。

- 如果 A 表 TID 是自增长,并且是连续的,B 表的 ID 为索引

```sql
select * from a,b where a.tid = b.id and a.tid>50000 limit 200;
```

- 如果 A 表的 TID 不是连续的,那么就需要使用覆盖索引。TID 要么是主键,要么是辅助索引,B 表 ID 也需要有索引。

```sql
select * from b , (select tid from a limit 50000,200) where b.id = a .tid;
```

## 索引是什么?有什么作用以及优缺点?

1. 索引是对数据库表中一列或多列的值进行排序的结构，是帮助 MySQL 高效获取数据的数据结构
2. 索引是加快检索表中数据的方法。数据库索引类似于书籍的索引目录，在书籍中，索引允许用户不必翻阅完整个书就能迅速地找到所需要的信息。在数据库中，索引也允许数据库程序迅速地找到表中的数据，而不必扫描整个数据库。

**MySQL 基本索引类型**：普通索引，唯一索引，主键索引，全文索引

**优缺点**：

1. 索引加快数据库的检索速度，但降低插入、删除、修改等维护任务的速度
2. 唯一索引可以确保每一行数据的唯一性
3. 通过使用索引，可以在查询的过程中使用优化隐藏器，提高系统性能
4. 索引需要占用物理和数据空间

## 什么是事务?

事务(Transaction)是并发控制的基本单位。所谓事务，它是一个操作序列，这些操作要么都执行，要么都不执行，它是一个不可分割的工作单位。事务是数据库维护数据一致性的单位，在每个事务结束时，都能保持数据的一致性。

## 使用索引一定能提高查询的性能吗?为什么

通常，通过索引查询数据比全表扫描要快，但是也必须注意它的代价：

1. 索引需要空间来存储，也需要定期维护，每当有记录在表中增减或索引列被修改时，索引本身也会被修改。这意味着每条记录的 INSERT,DELETE,UPDATE 将为此多付出 4，5 次的磁盘 I/O。因为索引需要额外的存储空间和处理，那些不必要的索引反而会使查询反应时间变慢。使用索引查询不一定能提高查询性能，索引范围查询(INDEX RANGE SCAN)适用于两种情况：
   1. 基于一个范围的检索，一般查询返回的结果集小于表中记录数的 30%
   2. 基于非唯一索引的检索

## 简单说一下 drop,delete,truncate 的区别

1. delete 和 truncate 只删除表数据而不删除表结构
2. 速度一般来说：drop>truncate>delete
3. delete 语句是 dml，这个操作会放到 rollback segement 中，事务提交之后才生效
4. 如果有相应的 trigger，执行的时候将被触发。truncate，drop 是 ddl，操作立即生效，原数据不放到 rollback segement 中，不能回滚，操作不触发 trigger。

- **DDL**：Data Definition Language，数据定义语言，这些语句定义了不同的数据段、数据库、表、列、索引等数据库对象的定义。常用的语句关键字主要包括 create、drop、alter 等。
- **DML**：Data Manipulation Language，数据操纵语言，用于添加、删除、更新和查询数据库记录，并检查数据完整性，常用的语句关键字主要包括 insert、delete、udpate 和 select 等。(增添改查）
- **DCL**：Data Control Language，数据控制语言，用于控制不同数据段直接的许可和访问级别的语句。这些语句定义了数据库、表、字段、用户的访问权限和安全级别。主要的语句关键字包括 grant、revoke 等。

## drop,delete,truncate 分别在什么场景下使用?

1. 不需要表时，用 drop
2. 想删除部分数据时，用 delete，并带上 where 子句
3. 保留表而删除所有数据的时候用 truncate

## 超键,候选键,主键,外键分别是什么?

1. **超键**：在关系中能唯一标识元组的属性称为关系模式的超键。一个属性可以为作为一个超键，多个属性组合在一起也可以作为一个超键。超键包含候选键和主键。
2. **候选键**：是最小的超键，即没有冗余元素的超键
3. **主键**：数据库表中对存储数据对象予以唯一和完整标识的数据列或属性的组合，一个数据列只能有一个主键，且主键的取值不能缺失，即不能为空值（Null）。
4. **外键**：在一个表中存在的另一个表的主键称为此表的外键

## 什么是视图?以及视图的使用场景有哪些?

1. 视图是虚拟的表，与包含数据的表不一样，视图只包含使用时动态检索数据的查询；
2. 不包含任何列或数据。使用视图可以简化复杂的 sql 操作，隐藏具体的细节，保护数据；
3. 视图创建后，可以使用与表相同的方式利用它们。
4. 视图不能被索引，也不能有关联的触发器或默认值，如果视图本身内有 order by 则对视图再次 order by 将被覆盖。
5. 创建视图：create view XXX as XXXXXXXXXXXXXX;
6. 对于某些视图比如未使用联结子查询分组聚集函数 Distinct Union 等，是可以对其更新的，对视图的更新将对基表进行更新；但是视图主要用于简化检索，保护数据，并不用于更新，而且大部分视图都不可以更新

## 参考

> - [MySQL innodb 的读写参数优化](http://www.jishuchi.com/read/mysql-interview/2821)
> - [浅析 MySQL 事务中的 redo 与 undo](https://segmentfault.com/a/1190000017888478)
