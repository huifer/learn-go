# gin 参数绑定
> 作者: [HuiFer](https://github.com/huifer)
>
> 仓库地址: https://github.com/huifer/learn-go

## 参数绑定的方式
### 在url中
#### query
- 例如: `/user/find?username=张三&age=18`

```go

type User struct {
	Username string `json:"username" form:"username"`
	Age int `json:"age" form:"age"`
}
//Query `/user/find?username=张三&age=18`
func Query( c *gin.Context) {
	username := c.DefaultQuery("username", "默认值")
	age := c.Query("age")
	m:= make(map[string]string)
	m["username"]=username
	m["age"] = age
	user := User{}
	c.ShouldBind(&user)
	fmt.Println("参数绑定到对象",user)
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"data":    m,
	})
}
```

- 路由绑定 

```go
	r.GET("/user/find",controller.Query)
```

- 请求 `GET http://localhost:9999/user/find?username=张三&age=18`


```json
{
  "data": {
    "age": "18",
    "username": "张三"
  },
  "message": "ok"
}
```

- 请求 `GET http://localhost:9999/user/find?age=18`

```json
{
  "data": {
    "age": "18",
    "username": "默认值"
  },
  "message": "ok"
}
```

- 请求 `GET localhost:9999/user/find?username=&age=18`

```json
{
  "data": {
    "age": "18",
    "username": ""
  },
  "message": "ok"
}
```

- 获取请求参数的方式

```go
	username := c.DefaultQuery("username", "默认值")
	age := c.Query("age")
```

1. `c.DefaultQuery` 这个方法参数说明 
    1. 需要获取的参数名称
    1. 默认值
    - 该方法不将参数传递过来会有一个默认值
        - 如果带有`GET localhost:9999/user/find?username=&age=18`那么这个默认值将无效
        - **在使用这个方法的时候一定要看url中是否带有参数名称**




##### 绑定到结构体
- 延续上一个例子我们创建出如下结构体

```go
type User struct {
	Username string `json:"username" form:"username"`
	Age int `json:"age" form:"age"`
}
```

- 使用 `	c.ShouldBind(&user)` 可以将参数绑定给 user 这个对象



#### path
- 例如: `/user/find/张三`

- 先编写一个处理器,用来获取请求参数

```go
//Path Url获取姓名和年龄 /user/find/张三
func Path(c *gin.Context) {
	name := c.Param("name")
	byName := c.Params.ByName("name")
	get, b := c.Params.Get("name")
	if b {
		fmt.Println("获取到 name ", get)
	}
	m := make(map[string]string)
	m["c.Param"] = name
	m["c.Params.ByName"] = byName
	m["c.Params.Get"] = get

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"data":    m,
	})

}

```

- 路由绑定

```go
r.GET("/user/find/:name",controller.Path)
```

- 访问结果

```json
{
  "data": {
    "c.Param": "张三",
    "c.Params.ByName": "张三",
    "c.Params.Get": "张三"
  },
  "message": "ok"
}
```

- 从 url 中获取请求参数的方式有 3 种
    1. `c.Param`
    1. `c.Params.ByName`
    1. `c.Params.Get`
    - 三个方法的参数均为你想要获取的参数名称. 其中,`c.Params.Get`返回2个参数 1.参数值 2. 是否存在这个参数名称, 推荐使用
    
    
### 在请求体中
#### form 表单
- 逻辑处理

```go
func BodyForm(c *gin.Context) {
	user := User{}
	err := c.ShouldBind(&user)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
			"data":    user,
		})
	}
}

```

- 路由绑定 `	r.POST("/user/form", controller.BodyForm)`



- 请求 
![image-20200618085917439](images/image-20200618085917439.png)

#### json 
- 处理逻辑

```go
func BodyJson(c *gin.Context) {
	user := User{}
	err := c.ShouldBind(&user)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
			"data":    user,
		})
	}
}

```

- 路由绑定 `	r.POST("/user/json", controller.BodyJson)`

- 请求 

```http request
POST http://localhost:9999/user/json
Content-Type: application/json

{

"username": "张三",
"age": 18
}
```