# echoatom
基于 Echo Gorm Casbin 实现的RBAC权限管理脚手架，目的是提供一套轻量的中后台开发框架，方便、快速的完成业务需求的开发。

### 获取代码

```
go get github.com/bodhi369/echoatom
```

### 初始化数据

> 1、修改 cmd/migration/main.go 文件中的 数据连接
> 2、go run ./cmd/migration/main.go

#### 运行服务

> 1、修改 cmd/api/conf.local.yaml 数据连接
> 2、运行服务 go run ./cmd/api/main.go

## 前端实现

- 待完成，将会基于 Ant Design Vue 实现

## Swagger 文档的使用

> 127.0.0.1:8080/swagger/

### 安装工具并生成文档

```
go get -u -v github.com/teambition/swaggo
swagger generate spec -b ./cmd/api -o ./assets/swaggerui/swagger.json --scan-models
```

## 项目结构概览

```
├── cmd
│    └── api
│    │    ├── main 主服务
│    │    └── conf.local.yaml  配置文件
│    └── migration
│        └──main 初始化数据
├── pkg
│    ├── api
│    │    └── api 路由
│    │        └── xxx 单一服务
│    │        ├── service 服务接口
│    │        ├── xxx udb实现
│    │        ├── logging 日志
│    │        ├── platform 具体数据操作
│    │        └── transport
│    │            ├── http 路由
│    │            └── swagger swagger文档 
│    └── utl 公共模块
│        ├── casbinplug 加载角色 用户 权限
│        ├── config 读取配置文件
│        ├── gormplug 数据连接
│        ├── middleware 中间件
│        ├── schemago 表结构
│        ├── secure 认证密码
│        ├── server 启动服务
│        ├── structs 合拼结构体
│        └── zlog 日志
└──  assets
    └── swaggerui 静态目录
```

## 感谢以下框架的开源支持
- [Echo] - [https://echo.labstack.com/](https://echo.labstack.com/)
- [GORM] - [http://gorm.io/](http://gorm.io/)
- [Casbin] - [https://casbin.org/](https://casbin.org/)

## 采用以下想法
- [ribice] - [https://www.ribice.ba/refactoring-gorsk/](https://www.ribice.ba/refactoring-gorsk/) 单文件夹服务
- [LyricTian] - [https://github.com/LyricTian/gin-admin/](https://github.com/LyricTian/gin-admin/) 数据结构



