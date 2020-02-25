# Redis Sentinel

- [Redis Sentinel](#redis-sentinel)
  - [介绍](#%e4%bb%8b%e7%bb%8d)
    - [Sentinel 功能的完整列表](#sentinel-%e5%8a%9f%e8%83%bd%e7%9a%84%e5%ae%8c%e6%95%b4%e5%88%97%e8%a1%a8)
    - [Sentinel 的分布式性质](#sentinel-%e7%9a%84%e5%88%86%e5%b8%83%e5%bc%8f%e6%80%a7%e8%b4%a8)
  - [快速开始](#%e5%bf%ab%e9%80%9f%e5%bc%80%e5%a7%8b)
    - [获取 Sentinel](#%e8%8e%b7%e5%8f%96-sentinel)
    - [运行 Sentinel](#%e8%bf%90%e8%a1%8c-sentinel)
    - [部署前有关 Sentinel 的基本知识](#%e9%83%a8%e7%bd%b2%e5%89%8d%e6%9c%89%e5%85%b3-sentinel-%e7%9a%84%e5%9f%ba%e6%9c%ac%e7%9f%a5%e8%af%86)
    - [配置 Sentinel](#%e9%85%8d%e7%bd%ae-sentinel)
    - [其他 Sentinel 选项](#%e5%85%b6%e4%bb%96-sentinel-%e9%80%89%e9%a1%b9)

## 介绍

Redis Sentinel 为 Redis 提供高可用性。使用 Sentinel 可以创建 Redis 部署，能够在没有人工干预的情况下抵抗某些类型的故障。

Redis Sentinel 还提供其他附带任务，例如监控，通知，和充当客户端的配置提供程序。

### Sentinel 功能的完整列表

- **Monitoring**：Sentinel 会不断检查主节点和副本节点是否按预期工作。
- **Notification**：Sentinel 可以通过 API 通知系统管理员或其他计算机程序某个 Redis 出了问题。
- **Automatic failover**：如果一个主服务器未按预期工作，则 Sentinel 可以启动一个故障转移进程，该进程可以将副本节点提升为主节点，其他副本节点被重新配置为使用新的主节点，并通知使用 Redis 服务器的应用程序连接时要使用的新地址。
- **Configuration provider**：Sentinel 充当客户端服务发现的授权来源：客户端连接到 Sentinels，以询问负责给定服务的当前 Redis 主服务器的地址。如果发生故障转移，Sentinels 将报告新地址。

### Sentinel 的分布式性质

Redis Sentinel 是一个分布式系统：  
Sentinel 本身设计为在有多个 Sentinel 进程协同合作的配置中运行。具有多个 Sentinel 进程进行协作的优点如下：

1. 当多个哨兵就给定的主节点不再可用这一事实达成共识时，将执行故障检测。这降低了误报的可能性。
2. 即使不是所有的 Sentinel 进程都在工作，Sentinel 仍能正常工作，从而使系统能够应对故障。毕竟，拥有故障转移系统本身就是一个单点故障。

Sentinels，Redis 实例(主节点和副本节点)以及连接到 Sentinel 和 Redis 的客户端的总和一个具有特定属性的大型分布式系统。

## 快速开始

### 获取 Sentinel

当前 Sentinel 的版本为 **Sentinel 2**。它是使用更强大且更容易预测的算法（在本文档中进行说明）对 Sentinel 初始实现进行的重写。  
自 Redis 2.8 起已发布稳定版本的 Redis Sentinel。  
在*unstable branch*中进行了新的开发，并且有时新功能一旦被认为是稳定的，便会立即移植回最新的稳定分支。  
Redis 2.6 附带的 Redis Sentinel 版本 1 已被弃用，不应再使用。

### 运行 Sentinel

如果您正在使用 redis-sentinel 可执行文件（或者您具有指向 redis-server 可执行文件的符号链接），则可以使用以下命令行运行 Sentinel：

    redis-sentinel /path/to/sentinel.conf

否则，您可以直接使用 redis-server 可执行文件以 Sentinel 模式启动它：

    redis-server /path/to/sentinel.conf --sentinel

两种方式都是相同的。

但是，在运行 Sentinel 时 **`必须使用配置文件`**，因为系统将使用此文件来保存当前状态，以便在重新启动时重新加载该状态。如果为提供配置文件或配置文件路径不可写，Sentinel 将拒绝启动。  
默认情况下，Sentinel 监听 TCP 的 26379 端口连接，因此，保证 Sentinels 正常工作，**`需要打开 26379 端口`**，以接收来自其他 Sentinel 实例的 IP 地址连接。否则，Sentinels 间无法正常通信来达成相应的共识，将永远无法执行故障转移。

### 部署前有关 Sentinel 的基本知识

1. 一个健壮的部署至少需要三个 Sentinel 实例。
2. 应将这三个 Sentinel 实例放入独立的计算机或虚拟机中，它们的实例出现故障的可能相互独立。例如在不同的可用区域上运行的不同的物理服务器或虚拟机。
3. Sentinel+Redis 分布式系统不保证在出现故障时保留已确认的写入，因为 Redis 使用异步复制。但是，有一些部署 Sentinel 的方法是丢失写入的时间段仅限于某些时刻，同时也有其他一些不太安全的方法来部署它。
4. 客户端需要 Sentinel 支持。流行的客户端库具有 Sentinel 支持，但不是全部。
5. 如果你不在开发环境中不断进行测试，则没有安全的 HA(High Availability)设置，如果可以，甚至可以在生产环境中进行更好的测试。
6. Sentinel，Docker 或其他形式的 NAT 或端口映射应谨慎组合使用：Docker 执行端口重新映射时，会破坏主节点对其他 Sentinel 的自动发现(auto discovery)以及主节点的副本节点列表。

### 配置 Sentinel

Redis 源分发包里包含一个名为 sendinel.conf 的文件，该文件是一个自说明的示例配置文件，可用于配置 Sentinel，但是典型的最小配置文件如下所示：

    sentinel monitor mymaster 127.0.0.1 6379 2
    sentinel down-after-milliseconds mymaster 60000
    sentinel failover-timeout mymaster 180000
    sentinel parallel-syncs mymaster 1

    sentinel monitor resque 192.168.1.3 6380 4
    sentinel down-after-milliseconds resque 10000
    sentinel failover-timeout resque 180000
    sentinel parallel-syncs resque 5

只需要指定要监视的 master，即可为每个单独的 master（可以有多个副本）指定一个不同的名称。无需指定自动发现的副本节点。Sentinel 将使用有关副本节点的其他信息自动更新配置(以便在重启时保留该信息)。每当在故障转移期间将副本提升为主节点时，以及每次发现新的 Sentinel 时，都会重写配置。

上边示例配置监视两组 Redis 实例，每组实例由一个主节点和未定义数量的副本节点组成。一组实例为`mymaster`，另一组为`resque`。

sentinel monitor 配置命令的参数含义如下：

    sentinel monitor <master-group-name> <ip> <port> <quorum>
    # 监视一个名为mymaster的主服务器，ip：127.0.0.1，port：6379，仲裁集为2
    sentinel monitor mymaster 127.0.0.1 6379 2

**仲裁集(quorum)**：

- 仲裁集指需要就主节点不可访问这一事实达成共识的哨兵数，用于将主节点真正标记为发生故障，并在可能的情况下启动故障转移。
- 但是，**仲裁集仅用于检测故障**。为了能执行故障转移，需要将其中一个哨兵选为故障转移的 leader，并有权继续进行操作。这仅在**大多数 Sentinel 进程投票通过**后发生。

因此，如果有 5 个 Sentinel 进程，并且给定 master 的仲裁数为 2，则：

- 如果两个哨兵同时就 master 不可访问达成一致，则两者之一将尝试启动故障转移。
- 如果总共至少有三个 Sentinel 可以访问，则故障转移将被授权并实际上开始。

实际上，这意味着在故障期间，**如果大多数 Sentinel 进程无法进行通信（在少数分区中也没有故障转移），则 Sentinel 不会启动故障转移。**

### 其他 Sentinel 选项

其他选项几乎总是采用以下形式：

    sentinel <option_name> <master_name> <option_value>

- **down-after-milliseconds**指 Sentinel 开始认为实例已崩溃的情况下，该实例不可访问（无论是 PING 无响应还是出现错误）的时间（以毫秒为单位）。
- **parallel-syncs**设置可在故障转移后同时重新配置为使用新的主节点的副本数。数字越小，完成故障转移过程所花费的时间就越多，但是，如果将副本配置为允许使用过期数据集，则可能不希望所有副本同时与主数据库重新同步。尽管副本节点大多数情况下不会阻塞复制进程，但从服务器在载入主服务器发来的 RDB 文件时， 仍然会造成从服务器在一段时间内不能处理命令请求。你可以通过将这个值设为 1 来保证每次只有一个从服务器处于不能处理命令请求的状态。

所有配置参数都可以在运行时使用 SENTINEL SET 命令进行修改。
