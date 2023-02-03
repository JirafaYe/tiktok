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

