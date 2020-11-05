## immigrate from flask to Gin

> ![head](https://miro.medium.com/max/700/1*zXGly_BHR_FQN3ngX-U7sQ.png)

> Recently one of my project need to be immigrated from python flask to golang Gin. Consider it will be a long-last process, hence it should be immigrate part by part. So my Gin-based proj need to share session with that flask created with sso.

> 近期有一个接入了SSO的项目需要从python flask换成go-gin，因为是个长期的过程，需要一部分一部分进行迁移，同时以前的python后端也是在线的。比较麻烦的地方在SSO相关的功能是原python后端提供的，而flask写到session里的信息是使用`app.secret_key`加密过的，所以golang 的project需要能支持解析python的后端在session中加密过的数据。

### project structure
```
.
├── Makefile
├── README.md
├── config
│   ├── config.go
│   ├── development.yaml
│   └── production.yaml
├── controllers
│   └── health.go
├── db
│   └── db.go
├── go.mod
├── go.sum
├── main.go
├── middlewares
│   ├── cas.go
│   ├── cas_test.go
│   ├── log.go
│   ├── redis.go
│   └── schedulerHandler.go
└── server
    ├── router.go
    └── server.go
```
