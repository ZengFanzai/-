# MySQL 数据库存储引擎

## InnoDB

InnoDB 还引入了行级锁定和外键约束，在以下场合下，使用 InnoDB 是最理想的选择：

- 更新密集的表。InnoDB 存储引擎特别适合处理多重并发的更新请求。
- 事务。InnoDB 存储引擎是支持事务的标准 MySQL 存储引擎。
- 自动灾难恢复。与其它存储引擎不同，InnoDB 表能够自动从灾难中恢复。
- 外键约束。MySQL 支持外键的存储引擎只有 InnoDB。
- 支持自动增加列 AUTO_INCREMENT 属性。
- 从 5.7 开始 innodb 存储引擎成为默认的存储引擎。

一般来说，如果需要事务支持，并且有较高的并发读取频率，InnoDB 是不错的选择。

## MyISAM

MyISAM 表是**独立于操作系统**的，这说明可以轻松地将其从 Windows 服务器移植到 Linux 服务器；每当我们建立一个 MyISAM 引擎的表时，就会在本地磁盘上建立三个文件，文件名就是表名。例如，我建立了一个 MyISAM 引擎的 tb_Demo 表，那么就会生成以下三个文件：

- tb_demo.frm，存储表定义。
- tb_demo.MYD，存储数据。
- tb_demo.MYI，存储索引。

MyISAM 表无法处理事务，这就意味着有事务处理需求的表，不能使用 MyISAM 存储引擎。MyISAM 存储引擎特别适合在以下几种情况下使用：

- 选择密集型的表。MyISAM 存储引擎在筛选大量数据时非常迅速，这是它最突出的优点。
- 插入密集型的表。MyISAM 的并发插入特性允许同时选择和插入数据。例如：MyISAM 存储引擎很适合管理邮件或 Web 服务器日志数据。

## MRG_MYISAM

**MRG_MyISAM 存储引擎是一组 MyISAM 表的组合**，老版本叫 MERGE 其实是一回事儿，这些 MyISAM 表结构必须完全相同，尽管其使用不如其它引擎突出，但是在某些情况下非常有用。说白了，**Merge 表就是几个相同 MyISAM 表的聚合器**；**Merge 表中并没有数据，对 Merge 类型的表可以进行查询、更新、删除操作，这些操作实际上是对内部的 MyISAM 表进行操作**。

Merge 存储引擎的使用场景。**对于服务器日志这种信息，一般常用的存储策略是将数据分成很多表，每个名称与特定的时间端相关**。例如：可以用 12 个相同的表来存储服务器日志数据，每个表用对应各个月份的名字来命名。当有必要基于所有 12 个日志表的数据来生成报表，这意味着需要编写并更新多表查询，以反映这些表中的信息。与其编写这些可能出现错误的查询，不如将这些表合并起来使用一条查询，之后再删除 Merge 表，而不影响原来的数据，删除 Merge 表只是删除 Merge 表的定义，对内部的表没有任何影响。

- ENGINE=MERGE，指明使用 MERGE 引擎，其实是跟 MRG_MyISAM 一回事儿，也是对的，在 MySQL 5.7 已经看不到 MERGE 了。
- UNION=(t1, t2)，指明了 MERGE 表中挂接了些哪表，可以通过 alter table 的方式修改 UNION 的值，以实现增删 MERGE 表子表的功能。比如：

  ```sql
  alter table tb_merge engine=merge union(tb_log1) insert_method=last;
  ```

- INSERT_METHOD=LAST，INSERT_METHOD 指明插入方式，取值可以是：0 不允许插入；FIRST 插入到 UNION 中的第一个表； LAST 插入到 UNION 中的最后一个表。
- MERGE 表及构成 MERGE 数据表结构的各成员数据表必须具有完全一样的结构。每一个成员数据表的数据列必须按照同样的顺序定义同样的名字和类型，索引也必须按照同样的顺序和同样的方式定义。

## MEMORY

**使用 MySQL Memory 存储引擎的出发点是速度**。**为得到最快的响应时间，采用的逻辑存储介质是系统内存**。虽然在内存中存储表数据确实会提供很高的性能，但当 mysqld 守护进程崩溃时，所有的 Memory **数据都会丢失**。获得速度的同时也带来了一些缺陷。它要求存储在 Memory 数据表里的数据**使用的是长度不变的格式**，这意味着不能使用 BLOB 和 TEXT 这样的长度可变的数据类型，VARCHAR 是一种长度可变的类型，但因为它在 MySQL 内部当做长度固定不变的 CHAR 类型，所以可以使用。

一般在以下几种情况下使用 Memory 存储引擎：

- **目标数据较小，而且被非常频繁地访问**。在内存中存放数据，所以会造成内存的使用，可以通过参数 max_heap_table_size 控制 Memory 表的大小，设置此参数，就可以限制 Memory 表的最大大小。
- 如果数据是临时的，而且要求必须立即可用，那么就可以存放在内存表中。
- 存储在 Memory 表中的数据如果突然丢失，不会对应用服务产生实质的负面影响。
- Memory 同时支持 hash 索引和 B 树索引。B 树索引的优于 hash 索引的是，可以使用部分查询和通配查询，也可以使用<、>和>=等操作符方便数据挖掘。hash 索引进行“相等比较”非常快，但是对“范围比较”的速度就慢多了，因此 hash 索引值适合使用在=和<>的操作符中，不适合在<或>操作符中，也同样不适合用在 order by 子句中。

## CSV

CSV 存储引擎是基于 CSV 格式文件存储数据。

- CSV 存储引擎因为自身文件格式的原因，所有列必须强制指定 NOT NULL 。
- CSV 引擎也不支持索引，不支持分区。
- CSV 存储引擎也会包含一个存储表结构的 .frm 文件，还会创建一个 .csv 存储数据的文件，还会创建一个同名的元信息文件，该文件的扩展名为 .CSM ，用来保存表的状态及表中保存的数据量。
- 每个数据行占用一个文本行。

因为 csv 文件本身就可以被 Office 等软件直接编辑，保不齐就有不按规则出牌的情况，如果出现 csv 文件中的内容损坏了的情况，也可以使用 CHECK TABLE 或者 REPAIR TABLE 命令检查和修复

## ARCHIVE(归档)

Archive 是归档的意思，在归档之后很多的高级功能就不再支持了，仅仅支持最基本的插入和查询两种功能。在 MySQL 5.5 版以前，Archive 是不支持索引，但是在 MySQL 5.5 以后的版本中就开始支持索引了。Archive 拥有很好的压缩机制，它使用 zlib 压缩库，在记录被请求时会实时压缩，所以它经常被用来**当做仓库使用**。

## BLACKHOLE

黑洞存储引擎，所有插入的数据并不会保存，BLACKHOLE 引擎表永远保持为空，写入的任何数据都会消失，用于记录 binlog 做复制的中继存储！

## PERFORMANCE_SCHEMA

**主要用于收集数据库服务器性能参数**。**MySQL 用户是不能创建存储引擎为 PERFORMANCE_SCHEMA 的表**，**一般用于记录 binlog 做复制的中继**。在这里有官方的一些介绍[MySQL Performance Schema](https://dev.mysql.com/doc/refman/5.7/en/performance-schema.html)

## FEDERATED

**主要用于访问其它远程 MySQL 服务器一个代理**，它通过创建一个到远程 MySQL 服务器的客户端连接，并将查询传输到远程服务器执行，而后完成数据存取；在 MariaDB 的上实现是 FederatedX

## 其他

这里列举一些其它数据库提供的存储引擎，OQGraph、SphinxSE、TokuDB、Cassandra、CONNECT、SQUENCE。提供的名字仅供参考。

## 常用引擎对比

不同存储引起都有各自的特点，为适应不同的需求，需要选择不同的存储引擎，所以首先考虑这些存储引擎各自的功能和兼容。

| 特性                                                   | InnoDB | MyISAM | MEMORY | ARCHIVE |
| ------------------------------------------------------ | ------ | ------ | ------ | ------- |
| 存储限制(Storage limits)                               | 64TB   | No     | YES    | No      |
| 支持事物(Transactions)                                 | Yes    | No     | No     | No      |
| 锁机制(Locking granularity)                            | 行锁   | 表锁   | 表锁   | 行锁    |
| B 树索引(B-tree indexes)                               | Yes    | Yes    | Yes    | No      |
| T 树索引(T-tree indexes)                               | No     | No     | No     | No      |
| 哈希索引(Hash indexes)                                 | Yes    | No     | Yes    | No      |
| 全文索引(Full-text indexes)                            | Yes    | Yes    | No     | No      |
| 集群索引(Clustered indexes)                            | Yes    | No     | No     | No      |
| 数据缓存(Data caches)                                  | Yes    | No     | N/A    | No      |
| 索引缓存(Index caches)                                 | Yes    | Yes    | N/A    | No      |
| 数据可压缩(Compressed data)                            | Yes    | Yes    | No     | Yes     |
| 加密传输(Encrypted data[1])                            | Yes    | Yes    | Yes    | Yes     |
| 集群数据库支持(Cluster databases support)              | No     | No     | No     | No      |
| 复制支持(Replication support[2])                       | Yes    | No     | No     | Yes     |
| 外键支持(Foreign key support)                          | Yes    | No     | No     | No      |
| 存储空间消耗(Storage Cost)                             | 高     | 低     | N/A    | 非常低  |
| 内存消耗(Memory Cost)                                  | 高     | 低     | N/A    | 低      |
| 数据字典更新(Update statistics for data dictionary)    | Yes    | Yes    | Yes    | Yes     |
| 备份/时间点恢复(backup/point-in-time recovery[3])      | Yes    | Yes    | Yes    | Yes     |
| 多版本并发控制(Multi-Version Concurrency Control/MVCC) | Yes    | No     | No     | No      |
| 批量数据写入效率(Bulk insert speed)                    | 慢     | 快     | 快     | 非常快  |
| 地理信息数据类型(Geospatial datatype support)          | Yes    | Yes    | No     | Yes     |
| 地理信息索引(Geospatial indexing support[4])           | Yes    | Yes    | No     | Yes     |

备注：

1. 在服务器中实现（通过加密功能）。在其他表空间加密数据在 MySQL 5.7 或更高版本兼容。
2. 在服务中实现的，而不是在存储引擎中实现的。
3. 在服务中实现的，而不是在存储引擎中实现的。
4. 地理位置索引，InnoDB 支持可 mysql5.7.5 或更高版本兼容

## 查看存储引擎

使用“SHOW VARIABLES LIKE '%storage_engine%';” 命令在 mysql 系统变量搜索个人设置的存储引擎，输入语句如下：

```sql
mysql> SHOW VARIABLES LIKE '%storage_engine%';
+----------------------------------+---------+
| Variable_name                    | Value   |
|----------------------------------+---------|
| default_storage_engine           | InnoDB  |
| default_tmp_storage_engine       | InnoDB  |
| disabled_storage_engines         |         |
| internal_tmp_disk_storage_engine | TempTable  |
+----------------------------------+---------+
4 rows in set
Time: 0.005s
```

使用“SHOW ENGINES;”命令显示安装以后可用的所有的支持的存储引擎和默认引擎

```sql
mysql> SHOW ENGINES;
+--------------------+---------+--------------------------------------+-------------+--------+-----------+
| Engine             | Support | Comment                              | Transactions| XA     | Savepoints|
|--------------------+---------+--------------------------------------+-------------+--------+-----------|
| InnoDB             | DEFAULT | Supports transactions,               | YES         | YES    | YES       |
|                    |         | row-level locking, and foreign keys  |             |        |           |
| MRG_MYISAM         | YES     | Collection of identical MyISAM tables| NO          | NO     | NO        |
| MEMORY             | YES     | Hash based, stored in memory, useful | NO          | NO     | NO        |
|                    |         | for temporary tables                 |             |        |           |
| BLACKHOLE          | YES     | /dev/null storage engine (anything   | NO          | NO     | NO        |
|                    |         | you write to it disappears)          |             |        |           |
| MyISAM             | YES     | MyISAM storage engine                | NO          | NO     | NO        |
| CSV                | YES     | CSV storage engine                   | NO          | NO     | NO        |
| ARCHIVE            | YES     | Archive storage engine               | NO          | NO     | NO        |
| PERFORMANCE_SCHEMA | YES     | Performance Schema                   | NO          | NO     | NO        |
| FEDERATED          | NO      | Federated MySQL storage engine       | <null>      | <null> | <null>    |
+--------------------+---------+--------------------------------------+-------------+--------+-----------+
```

## 设置存储引擎

对上面数据库存储引擎有所了解之后，你可以在`my.cnf` 配置文件中设置你需要的存储引擎，这个参数放在 [mysqld] 这个字段下面的 default_storage_engine 参数值，例如下面配置的片段

```text
[mysqld]
default_storage_engine=CSV
```

在创建表的时候，对表设置存储引擎，例如：

```sql
CREATE TABLE `user` (
  `id`     int(100) unsigned NOT NULL AUTO_INCREMENT,
  `name`   varchar(32) NOT NULL DEFAULT '' COMMENT '姓名',
  `mobile` varchar(20) NOT NULL DEFAULT '' COMMENT '手机',
  PRIMARY KEY (`id`)
)ENGINE=InnoDB;
```

## 如何选择合适的存储引擎

使用哪种引擎需要根据需求灵活选择，一个数据库中多个表可以使用不同的引擎以满足各种性能和实际需求。使用合适的存储引擎，将会提高整个数据库的性能。

- 是否需要支持事务；
- 是否需要使用热备；
- 崩溃恢复，能否接受崩溃；
- 是否需要外键支持；
- 存储的限制；
- 对索引和缓存的支持；

## Innodb 和 MyISAM 区别

- **存储结构**：

  - **MyISAM**：每个 MyISAM 在磁盘上存储成三个文件。分别为：**表定义文件、数据文件、索引文件**。第一个文件的名字以表的名字开始，扩展名指出文件类型。.frm 文件存储表定义。数据文件的扩展名为.MYD (MYData)。索引文件的扩展名是.MYI (MYIndex)。
  - **Innodb** 所有的表都保存在同一个数据文件中（也可能是多个文件，或者是独立的表空间文件），InnoDB 表的大小只受限于操作系统文件的大小，一般为 2GB。

- **存储空间**：

  - **MyISAM**： MyISAM 支持支持三种不同的存储格式：**静态表**(默认，但是注意数据末尾不能有空格，会被去掉)、**动态表**、**压缩表**。当表在创建之后并导入数据之后，不会再进行修改操作，可以使用压缩表，极大的减少磁盘的空间占用。
  - **Innodb**： 需要更多的内存和存储，它会在主内存中建立其专用的缓冲池用于高速缓冲数据和索引。

- **可移植、备份及恢复**：

  - **MyISAM**：数据是以文件的形式存储，所以在跨平台的数据转移中会很方便。`在备份和恢复时可单独针对某个表进行操作`。
  - **InnoDB**：免费的方案可以是拷贝数据文件、备份 binlog，或者用 mysqldump，在数据量达到几十 G 的时候就相对痛苦了。

- **AUTO_INCREMENT**：

  - **MyISAM**：可以和其他字段一起建立联合索引。引擎的自动增长列必须是索引，如果是组合索引，自动增长可以不是第一列，他可以根据前面几列进行排序后递增。
  - **InnoDB**：InnoDB 中必须包含只有该字段的索引。引擎的自动增长列必须是索引，如果是组合索引也必须是组合索引的第一列。

- **事务**：Innodb 支持事务，而 MyISAM 不支持事务
- **锁**：Innodb 支持行级锁，MyISAM 支持表级锁
- **外键**：Innodb 支持外键，MyISAM 不支持
- **全文索引**：Innodb 不支持全文索引(5.6.4 以后版本开始支持 FULLTEXT 类型的索引)，而 Myisam 支持
- **表**：Innodb 是索引组织表， Myisam 是堆表
- **表主键**：

  - **MyISAM**：允许没有任何索引和主键的表存在，索引都是保存行的地址。
  - **InnoDB**：如果没有设定主键或者非空唯一索引，就会自动生成一个 6 字节的主键(用户不可见)，数据是主索引的一部分，附加索引保存的是主索引的值。

- **表的具体行数**：

  - **MyISAM**： 保存有表的总行数，如果 select count() from table;会直接取出出该值。- **InnoDB**： 没有保存表的总行数，如果使用 select count(\*) from table；就会遍历整个表，消耗相当大，但是在加了 wehre 条件后，myisam 和 innodb 处理的方式都一样。

- **CRUD 操作**：

  - **MyISAM**：如果执行大量的 SELECT，MyISAM 是更好的选择。
  - **InnoDB**：如果你的数据执行大量的 INSERT 或 UPDATE，出于性能方面的考虑，应该使用 InnoDB 表。

**简记为：**

- InnoDB 支持事务，MyISAM 不支持事务
- InnoDB 支持行级锁，MyISAM 支持表级锁
- InnoDB 支持 MVCC(多版本并发控制)，而 MyISAM 不支持
- InnoDB 支持外键，而 MyISAM 不支持
- InnoDB 即使没设计主键也会新建一个对用户不可见的主键，而 MyISAM 支持不存在主键和索引的表
- InnoDB 没有保存表的总行数，MyISAM 保存有表的总行数
- InnodB 是索引组织表，MyISAM 是堆表
- InnoDB 中必须包含只有该字段的索引，引擎的自动增长列必须是索引，如果是组合索引也必须是组合索引的第一列，MyISAM 可以和其他字段一起建立联合索引。引擎的自动增长列必须是索引，如果是组合索引，自动增长可以不是第一列，他可以根据前面几列进行排序后递增。
- InnoDB 适合多 update 和 insert 的表，MyISAM 适合多 select 的表

## InnoDB 特性(前四个重要特性)

- **插入缓冲（insert buffer)**：用于**非聚集索引页**的插入和更新操作，提高插入性能。
- **二次写(double write)**：为了确保在宕机时刷新脏页不会出现问题，先将原始页做一份备份到磁盘，再写脏页，出现问题可以通过备份进行恢复。
- **自适应哈希索引(ahi)**：为了提高辅助索引的查询效率，对于频繁查询的辅助索引会自动建立 hash 索引。
- **预读(read ahead)**：预读机制就是发起一个 i/o 请求，异步地在缓冲池中预先回迁若干页面，预计将会用到回迁的页面，这些请求在一个范围内引入所有页面。
- 脏页：缓存中的页面被修改后的页面，我们称之为脏页。
- 重做日志：用来记录用户对数据的修改过程，定时刷新到磁盘，用于断电重启时对缓存中的脏页进行恢复。
- checkPoint：将脏页数据刷新到磁盘的技术。
- LSN(Log Sequence Number): 每个缓存页/重做日志/checkPoint 都有 LSN，用来记录顺序节点位置。
- undo 页：记录未被提交的事务操作过程。
- pull merge：主要用来删除无用的 undo 页。
- 异步 IO：为了提高磁盘操作的性能， InnoDB 存储引擎采用异步 IO（AIO）的方式来处理磁盘操作。
- 刷新邻接页：在刷新脏页时，检查该页所在区的所有页是否存在脏页，存在则一起刷新。

### 插入缓存(Insert Buffer)

判断当前的非聚集索引是否在缓冲池，如果在则直接插入，否则插入到`Insert Buffer`对象里，然后以一定频率进行`Insert Buffer`和辅助索引叶子节点的合并。简单来说，就是合并多个插入操作为一个操作，从而提高插入性能。相关参数`innodb_change_buffer_max_size`

使用 Insert Buffer 要满足以下两个条件：

- 索引是二级(/辅助)索引(唯一索引、普通索引、前缀索引)
- 索引不唯一，如果索引唯一还要去查找索引页进行检查唯一性，就失去了 Insert Buffer 离散插入的性能

Insert Buffer 中的记录何时被插入到真正的辅助索引页（Merge Insert Buffer）中：

- 当辅助索引页被读取到缓存池时，检查 Insert Buffer Bitmap 页是否有该页对应的 Insert Buffer，有则合并。
- 当每个辅助索引页的可用空间小于 1/32 时，强制进行一次合并操作。
- Master Thread 线程会每秒或者每 10 秒进行一次 Merge Insert Buffer。（定时合并）

Insert Buffer 内部实现：

- Insert Buffer 存放在一颗 B+ 树中，非叶子节点存放的是 search key（表空间，页的偏移量等信息），叶子节点存放的真正的记录
- Insert Buffer Bitmap 用来标记每个辅助索引页的可用空间以及该页是否已经在缓存中等信息，这样可以保证保证每次 Merge Insert Buffer 成功

### 二次写

数据库宕机可通过重做日志对该页进行恢复，但针对该页已经损坏的情况，重做恢复没什么意义，所以引入`二次写`来提高数据页的稳定性。

**二次写流程：**

- 对缓存池的脏页进行刷新时，不直接写入磁盘，而是通过 memcpy 函数将脏页复制到内存中的 doublewrite buffer。
- 将内存中的 doublewrite buffer 写入共享表空间的物理磁盘上（备份）。
- 将 doublewrite buffer 中的数据真正的刷新到表磁盘中。
- 如果写 doublewrite buffer 失败，那么这些数据不会写到磁盘，innodb 会载入磁盘原始数据和 redo 日志比较，并重新刷到 doublewrite buffer。
- 如果写 doublewrite buffer 成功，但是刷新到磁盘失败，那么 innodb 就不会通过 redo 日志来恢复了，而是直接刷新 double write buffer 中的数据到磁盘。
- `skip_innodb_double_written` 参数可以用来禁止 doublewrite(二次写) 功能。

### 自适应哈希索引(Adaptive Hash Index，AHI)

InnoDB 会监控对表上各个索引页的查询，如果观察到通过哈希索引可以带来性能提升，则自动建立哈希索引。AHI 通过缓存池的 B+ 树页构造而来，因此建立速度很快。

- AHI 可以提高读写性能
- AHI 只用来搜索等值查询
- 页的连续访问模式必须是一样的，才能建立 AHI。
- 自适应哈希索引**只是针对辅助索引**会自适应建立。

对于(a, b)这样的联合索引，访问模式有以下两种：

1. where a = xxx
2. where a = xxx and b = xxx

### 异步 IO

为了提高磁盘操作的性能，InnoDB 存储引擎采用异步 IO（AIO）的方式来处理磁盘操作：

- 可以减少请求查询时间
- 可以对多个 IO 进行合并操作，提供 IOPS 的性能
- 在 InnoDB1.1.x 之前是通过代码模拟实现 AIO 的，在之后的版本中提供了内核级别的 AIO 支持(Native AIO)
- 可以通过 `innodb_use_native_aio` 参数控制是否开启 Native AIO

### 刷新邻接页

- 刷新邻接页功能是指在刷新脏页时，检查该页所在区的所有页是否存在脏页，存在则一起刷新，这样可以通过 AIO 进行脏页合并刷新
- 对于传统的机械硬盘建议开启该功能，对于固态硬盘有着较高的 IOPS，可以不用开启
- 通过 `innodb_flush_neighbors` 参数控制是否开启刷新邻接页功能

### InnoDB 的启动/关闭/恢复

- `innodb_fast_shutdown` 参数设置数据库关闭时的 InnoDB 的行为，可取 0，1，2 三个值
  - 0 表示完成所有的 pull merge 和 merge insert buffer，并将脏页刷新到磁盘操作以后再关闭
  - 1 表示只完成将脏页刷新到磁盘即关闭
  - 2 只是将日志写入文件就关闭数据库，下次启动时可以根据 redo log 进行恢复
- `innobb_force_recovery` 参数控制数据启动时，InnoDB 恢复和还原的行为，可取值是 1-6， 其中默认值为 0，每次启动时都会对发生问题的表进行恢复操作

## 参考

> - [《史上最详细的一线大厂 Mysql 面试题详解》](https://juejin.im/post/5cb6c4ef51882532b70e6ff0#heading-0)
> - [InnoDB 存储引擎的关键特性学习](https://zhuanlan.zhihu.com/p/64180357)
> - [Mysql 中 MyISAM 和 InnoDB 的区别有哪些？](https://www.zhihu.com/question/20596402)
> - [数据库存储引擎](https://github.com/jaywcjlove/mysql-tutorial/blob/master/chapter3/3.5.md#%E6%95%B0%E6%8D%AE%E5%BA%93%E5%AD%98%E5%82%A8%E5%BC%95%E6%93%8E)
