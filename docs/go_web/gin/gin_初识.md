# gin 初识
> 作者: [HuiFer](https://github.com/huifer)
>
> 仓库地址: https://github.com/huifer/learn-go

## 简介
> Gin is a web framework written in Go (Golang). It features a martini-like API with performance that is up to 40 times faster thanks to httprouter. If you need performance and good productivity, you will love Gin.
>

Gin是一个用Go（Golang）编写的 web 框架。它采用了类似 martini 的API，由于 httprouter，性能提高了40倍。如果你需要表现和良好的生产力，你会喜欢gin。
- github: https://github.com/gin-gonic/gin
## 安装
- `go get -u github.com/gin-gonic/gin`

## 使用

```go
import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
```