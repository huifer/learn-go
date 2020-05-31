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