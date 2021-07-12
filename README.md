# comment
a comment project

## 设计
这是一个使用go语言搭建的微服务演示项目。
项目模拟了一个评论系统

## 特性
* 统一项目结构
* 实践多个服务间通信
* 实例与基础设施服务gkit结合
* 这个项目只是模拟微服务的方案，请大胆发挥想象力。

## 组件
这一章节描述项目的各个组件。

### `api`
所有的 API `.proto` 文件和生成的 `.go` 文件都在此目录。
此目录结构和目录 `/app/` 类似。

### `app`
所有服务的源码都放在这个目录。

#### Service: `/app/comment/service`
##### 作用/功能
此服务提供前端评论相关api
##### 特性

* 与其他服务交互

#### Job:`/app/comment/job`
##### 作用/功能
评论系统后台job。
##### 
#### Admin:`/app/comment/admin`

##### 作用/功能

订单系统管理后台。

### `/pkg/`
公共的包，各个服务都可以引用。

### `/deploy/`
dockerfile和部署脚本

### `/web/`
前端项目

## 架构
整个项目的架构蓝图[TBD]

 ```
  .
 ├── Dockerfile  
 ├── LICENSE
 ├── Makefile  
 ├── README.md
 ├── api // 下面维护了微服务使用的proto文件以及根据它们所生成的go文件
 │   └── comment
 │       └── service
 │           └── v1
 │               ├── comment.pb.go
 │               ├── comment.proto
 │               ├── greeter.swagger.json
 │               └── comment_grpc.pb.go
 ├── cmd  // 整个项目启动的入口文件
 │   └── server
 │       └── main.go
 │   └── admin
 │       └── main.go
 │   └── job
 │       └── main.go
 ├── configs  // 这里通常维护一些本地调试用的样例配置文件
 │   └── config.yaml
 ├── go.mod
 ├── go.sum
 ├── internal  // 该服务所有不对外暴露的代码，通常的业务逻辑都在这下面，使用internal避免错误引用
 │   ├── biz   // 业务逻辑的组装层，类似 DDD 的 domain 层，data 类似 DDD 的 repo
 │   │   ├── README.md
 │   │   ├── biz.go
 │   │   └── comment.go
 │   ├── conf  // 内部使用的config的结构定义，使用proto格式生成
 │   │   ├── conf.pb.go
 │   │   └── conf.proto
     ├── httpserver  // http handler层
 │   │   ├── comment.go
 │   │   └── response.go
 │   ├── store  // 业务数据访问，包含 cache、db 等封装
 │   │   ├── store.go
 │   │   └── comment.go
 │   ├── server  // http和grpc实例的创建和配置
 │   │   ├── grpc.go
 │   │   ├── http.go
 │   │   └── server.go
 │   └── service  // 实现了 api 定义的服务层，类似 DDD 的 application 层，处理 DTO 到 biz 领域实体的转换(DTO -> DO)，同时协同各类 biz 交互，但是不应处理复杂逻辑
 │       ├── comment.go
 │       └── service.go
 └── third_party  // api 依赖的第三方proto
     ├── README.md
     ├── google
     │   └── api
     │       ├── annotations.proto
     │       ├── http.proto
     │       └── httpbody.proto
 ```