# GO 数组
> 作者: [HuiFer](https://github.com/huifer)
>
> 仓库地址: https://github.com/huifer/learn-go

## 初始化
- `[len]DataType`
    - len 表示数组长度
    - DataType 数组元素类型

- `a := [3]int{1,1,1}` 
    - 指定长度为3的int数组,字面量值`{1,1,1}`
- `a := [...]int{1}`
    - 不指定长度,通过`{}`中的元素数量来确定长度
- `a :=[3]int{1:1,2:1}`
    - 指定长度为3的int数组,通过索引进行初始化,没有设置使用默认值


## 数组特点
1. 数组创建完长度固定,不能追加元素
1. 数组是值类型的, 数组赋值或作为函数参数都是值拷贝
1. 数组昌都市数组类型的组成部分 `[1]int` 和 `[2]int`表示的是不同的类型
1. 可以根据数组创建切片

## 数组操作

```go
package array

import "fmt"

func Array() {
	a := [3]int{1, 2, 3,}
	fmt.Println(a)
	// 输出数组的长度
	i := len(a)
	fmt.Println("数字长度= " ,i )
	// 数组的访问
	i2 := a[0]
	fmt.Println("数组访问,index=0，value=",i2)

}

```

## 练习
1. 设置一个int类型的数组长度为10
1. 每个元素为索引值+1
1. 修改最后一个元素的值为100
- [答案](/Answer/ArrayAnswer.go)