# GO 分支语句
> 作者: [HuiFer](https://github.com/huifer)
>
> 仓库地址: https://github.com/huifer/learn-go


## if
- if 后面的条件语句不需要使用`()`
- `{` 必须放在行尾, if 或者 else 的同一行内
- **没有三目运算**
- if 语句遇到 return 直接返回, 遇到 break 跳过 if 语句块

```go

	var a = 1
	if a > 0 {
		fmt.Print(">0")
	} else if a > 1 {
		fmt.Print(">1")
	} else {
		fmt.Print("其他")
	}

```


## switch

```go
var sw = "a"
switch sw {
case "a", "c":
    fmt.Print(sw)
case lang:
    fmt.Print(lang)
default:
    fmt.Print("default")
}
```

- 通过 fallthough 强制执行下一个 case 不管是否满足 case 后面的条件
