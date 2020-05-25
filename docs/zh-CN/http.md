## 简介
基础的`Http`服务和`Router`以及`Context`处理

### 路由定义
`firmeve`路由是基于`httprouter`进行扩展，更多`httprouter`用法参见其文档

```go
//基础示例
router := http.New(firmeve.New())
router.GET("/ping", func(ctx contract.Context) {
    ctx.RenderWith(200,render2.Plain,"pong")
    ctx.Next()
})
```

### 路由中间件
```go
router.GET("/ping", func(ctx contract.Context) {
    ctx.RenderWith(200,render2.Plain,"pong")
    ctx.Next()
}).Use(func(ctx contract.Context) {
    ctx.RenderWith(200,render2.Plain,"Before")
  ctx.Next()
}).Use(func(ctx contract.Context) {
    ctx.RenderWith(200,render2.Plain,"After")
   ctx.Next()
})
```

### 路由分组
```go
v1 := router.Group("/api/v1").Use(func(ctx *http.Context) {
                                ctx.RenderWith(200,render2.Plain,"Before")
                               ctx.Next()
                             },func(ctx *http.Context) {
                               ctx.RenderWith(200,render2.Plain,"After")
                               ctx.Next()
                             })
{
	v1.Get("/ping", func(ctx *http.Context) {
       ctx.RenderWith(200,render.JSON,map[string]string{
       	    "message": "something"
       })
       ctx.Next()
   })
}
```


### 启动Http服务

#### Http服务
```bash
go run main.go http:serve --host=0.0.0.0:22182
```

#### Https服务
```bash
go run main.go http:serve --host=0.0.0.0:22182 --key-file=server.key --cert-file=server.crt
```

#### Http2服务
```bash
go run main.go http:serve --host=0.0.0.0:22182 --key-file=server.key --cert-file=server.crt --http2
```

> 假设`server.key`,`server.crt`和`main`在同一目录


### Context

#### 可用方法
```go
Firmeve() Application

// 获取当前协议
Protocol() Protocol

Next()

Handlers() []ContextHandler

AddEntity(key string, value interface{})

Entity(key string) *ContextEntity

Abort()

Error(status int, err error)

Bind(v interface{}) error

BindWith(b Binding, v interface{}) error

Get(key string) interface{}

Render(status int, v interface{}) error

RenderWith(status int, r Render, v interface{}) error

Clone() Context

Resolve(abstract interface{}, params ...interface{}) interface{}
```
