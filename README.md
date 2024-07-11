# Gin Project

## project structure

From: https://www.kancloud.cn/lhj0702/sockstack_gin/1805358

```shell
├── conf                    #项目配置文件目录
│   └── config.toml         #大家可以选择自己熟悉的配置文件管理工具包例如：toml、xml、ini等等
├── requests                #定义入参即入参校验规则
│   └── user_request.go
│   └── food_request.go
├── responses                #定义响应的数据
│   └── user_response.go
│   └── food_response.go
├── services                #服务定义目录
|	└── v1					#服务v1版本
│   |	└── user_service.go
│   |	└── food_service.go
|	└── v2					#服务v2版本
│   |	└── user_service.go
│   |	└── food_service.go
├── api             		#api目录，按模块存放控制器（或者叫控制器函数），必要的时候可以继续划分子目录。
│   └── v1					#apiv1版本
│   |	└── user.go
│   |	└── food.go
│   └── v2					#apiv2版本
│   |	└── user.go
│   |	└── food.go
├── router					#路由目录
│   └── v1					#路由v1版本
│   |	└── user.go
│   |	└── food.go
│   └── v2					#路由v2版本
│   |	└── user.go
│   |	└── food.go
├── init.go					#路由初始化
├── pkg						#自定义的工具类等
│   └── e					#项目统一的响应定义，如错误码，通用的错误信息，响应的结构体
│   └── util				#工具类目录
├── models                  #模型目录，负责项目的数据存储部分，例如各个模块的Mysql表的读写模型。
│   ├── food.go
│   ├── user.go
│	└── init.go				#模型初始化
├── repositories            #数据操作层，定义各种数据操作。
│   └── user_repository.go
│   └── food_repository.go
├── logs                    #日志文件目录，主要保存项目运行过程中产生的日志。
├── main.go                 #项目入口，这里负责Gin框架的初始化，注册路由信息，关联控制器函数等。
```
