# GO 切片 slice
> 作者: [HuiFer](https://github.com/huifer)
>
> 仓库地址: https://github.com/huifer/learn-go

- GO slice 
- 变长数组

## 切片的创建
### 通过数组创建
- `array[begin_index:end_index]`
    - array : 数组名
    - begin_index : 开始的索引位置,可以不填写,不填写默认从0开始
    - end_index : 结束的索引位置,可以不填写,不填写默认数组的长度
    
```go
	// 通过数组创建
	a := [3]int{1, 2, 3}
	start := a[1:]
	end1 := a[:len(a)]
	end2 := a[1:len(a)]

	fmt.Println(start)
	fmt.Println(end1)
	fmt.Println(end2)
```

###  通过 make 创建
```go
	b1 := make([]int , 10 )
	b2:= make([]int,10, 20)
	fmt.Println(b1)
	fmt.Println(b2)
```

### 内置函数
1. len(slice): 返回切片长度
1. cap(slice): 返回切片底层数组容量
1. append(slice,element) : 添加元素, 返回值是添加后的结果,slice不会修改
1. copy(slice1,slice2): 拷贝 slice1 的值 复制给 slice2

```go
    // 内置函数
	// 返回切片长度
	fmt.Println(len(b1))
	// 返回切片底层数组容量
	fmt.Println(cap(b1))
	// 追加元素,原slice 不会修改
	ints := append(b1, 1)
	fmt.Println("追加返回值" , ints)
	fmt.Println("原始值",b1)
	// 拷贝
	b3 := make([]int , 10 , 10)
	copy(b1, b3 )
	fmt.Println(b3)
```

## 数据结构 

```go
type slice struct {
	array unsafe.Pointer // 指针
	len   int   // 切片的元素数量
	cap   int   // 底层数组的容量
}
```


## 练习
1. 创建数组 []int{1,2,3}
2. 获取这个数组的前半部分和后半补分
- [答案](/Answer/SliceAnswer.go)