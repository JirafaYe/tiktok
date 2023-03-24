# tiktok

基于 grpc RPC微服务 + gin HTTP服务完成的第5届字节跳动青训营-极简抖音后端项目

## 项目结构

```FILE
grpc模块
.
├── center ------------> 注册中心相关
├── cmd ---------------> grpc服务启动main文件
├── config ------------> 读取配置相关
├── internal ----------> 具体业务逻辑存放目录
│   ├── server --------> grpc服务的具体实现
│   └── service -------> 存放grpc生成的文件以及proto文件
│       ├──proto ------> 存放proto文件
│   └── store ---------> 数据库以及缓存相关
│       ├── local -----> 存放model相关的操作
│       └── cache -----> 存放缓存相关的操作
├── pkg ---------------> 项目所需工具类
├── .gitignore
├── Dockerfile
├── Makefile

gateway
.
├── api ---------------> gin相关的接口实现
├── center ------------> 注册中心相关
├── config ------------> 读取配置相关
├── service -----------> 存放grpc生成的文件以及proto文件
│   ├──proto ----------> 存放proto文件
├── main.go -----------> gateway模块启动main文件
├── .gitignore
├── Dockerfile
├── Makefile
```

- 采用RPC框架grpc，基于 **RPC服务** + **Gin 提供 HTTP服务**
- 基于《[接口文档在线分享- Apifox](https://www.apifox.cn/apidoc/shared-8cc50618-0da6-4d5e-a398-76f3b8f766c5/)》提供的接口进行开发，使用《[极简抖音App使用说明 - 青训营版](https://bytedance.feishu.cn/docs/doccnM9KkBAdyDhg8qaeGlIz7S7) 》提供的APK进行Demo测试， **功能完整实现** ，前端接口匹配良好
- 使用 **JWT** 进行用户token的校验
- 使用 **Consul** 进行服务发现和服务注册，以及配置文件的KV读取
- 使用 **Minio** 实现视频文件和图片的对象存储
- 使用 **Gorm** 对 MySQL 进行 ORM 操作
- 使用**redis**作为Nosql缓存
- [ ]  使用OpenTelemetry实现链路跟踪
- [ ]  数据库表建立了索引和外检约束，对于具有关联性的操作一旦出错立刻回滚，保证数据一致性和安全性

![image](https://user-images.githubusercontent.com/65102150/227455443-f986e5a0-c7cc-4a33-bdc0-c022c0dfa8ac.png)


模块解析：

> consul
> 

consul是一种分布式服务工具，提供：

- 服务发现
- 健康检查
- K-V存储
- 多数据中心

特点：

- consul常常以集群的方式部署，采用主从模式，有server和client两种节点。server节点保存数据，client节点负责健康检查和转发请求
- 使用http接口进行服务发现，可以通过http请求来注册服务和查询服务，也可以通过http api来获取服务的健康状态和元数据
- 使用raft算法保证服务的一致性，Server节点中有一个Leader节点和多个Follower及节点，Leader节点会将数据同步到Follower节点，当Leader节点挂掉时会启动选举机制产生一个新的Leader
- 支持多数据中心，不同数据中心之间通过WAN GOSSIP进行报文交互，可以实现跨区域的服务发现和配置共享
- 提供HTTP接口、DNS接口和SDK接口，可以方便地注册服务、发现服务、获取服务的健康状态和元数据

## 模块分析

gateway模块：

- api：存放gin的接口实现
    - Route函数注册路由
    - 对微服务进行负载均衡
    - 构建request并发送给微服务，等待微服务的response
    - 构建JSON（status code），返回给客户端
- center：存放注册中心相关
    - init函数：初始化一个consul客户端，连接到指定的地址
    - GetValue函数：从Consul的KV存储中读取指定的键值对
    - Register函数：向Consul注册一个服务，包括服务的名称、ID、地址、端口等信息
    - Resolver函数：从Consul中解析一个服务的地址，并使用gRPC进行通信。这个函数使用了一个第三方库grpc-consul-resolver来实现gRPC与Consul的集成。这个函数还指定了一些参数，如等待时间、筛选条件和负载均衡策略
- config：存放读取配置相关
- service：存放grpc生成的文件和protp文件

功能模块（publish模块）：

- center：包含了一个用于连接服务中心的客户端
- cmd：包含了一个用于启动服务的命令行工具
- config：包含了一些用于管理服务配置的类型和函数
- internal：包含了一些内部使用的包
- pkg：包含了一些公共使用的包

## Consul与ETCD的差异

- consul使用gossip协议来管理节点成员关系、失败检测、事件广播等
- consul提供了更多功能，比如健康检查、服务注册、DNS接口，ACL等。而ETCD更专注于核心的键值存储功能

## 链路跟踪

链路跟踪是一种用于分析和监控应用程序的方法，尤其是使用微服务架构构建的应用程序。它可以通过记录和关联每个请求在不同服务之间的路径，来帮助定位性能问题和故障

OpenTelemetry是一个开源项目，它统一了追踪、指标和日志的规范，定义了它们之间的联系。它可以让开发者无需改动或者很小的改动就可以接入不同的链路跟踪系统

