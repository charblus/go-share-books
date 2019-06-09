```
➜ go run main.go 
➜  src curl 127.0.0.1:9901
404 page not found% 
```     
> 启动服务后报404 因为没有路由


2. `web/route.go:15:9: cannot use origin.GetOrigin (type func(http.ResponseWriter, *http.Request)) as type gin.HandlerFunc in argument to r.GET`

筛选信息 ` gin.HandlerFunc in argument`
解决方案 ： 
[golang中vendor引起的相同类型，但是确提示类型不一样问题](https://blog.csdn.net/woodcutter_man/article/details/80386695)

[golang vendor 使用](https://blog.csdn.net/nimei31/article/details/80280297)
