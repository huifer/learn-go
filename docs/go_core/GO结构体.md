# GO 集合
> 作者: [HuiFer](https://github.com/huifer)
>
> 仓库地址: https://github.com/huifer/learn-go


## 语法

```go
type name struct {
}
```

- demo

```go
package structure

type User struct {
	Name string
	Age  int
}

func GetUser() *User {
	u := User{Age: 1, Name: "huifer"}
	return &u
}
func main (){

	user := GetUser()
	print(user.Age)
	print(user.Name)
}
```

## 注意
- 上文推荐使用 `字段名称: 属性值` 的方式进行使用, 还有一种不添加字段名称的方式

```go
User{"huifer",1}
```
这种方式进行初始化, 一旦 `User` 结构增加字段这个方法就会报错.s