# GO 字典 map
> 作者: [HuiFer](https://github.com/huifer)
>
> 仓库地址: https://github.com/huifer/learn-go

## map 定义
- map[k]T 
    - K 是任意可以比较的类型
    - T 是值类型

## map 创建
- 创建方式
    1. map 关键字
    2. make 关键字

```go
	// 字面量创建
	a_map := map[string]int{"a": 1, "b": 2}
	fmt.Println(a_map)
	// make 函数创建
	map1 := make(map[string]int)
	fmt.Println(map1)
```

## map 操作


```go
	// 1. 访问
	fmt.Println(a_map["a"])
	// 2. 删除
	delete(a_map, "a")
	// 3. 修改
	a_map["b"]=22
	// 4. 循环
	for s := range a_map {
		fmt.Println("key = " , s , " value = " , a_map[s])
	}

```

## 注意
1. map 是线程不安全的, 线程安全的 map 使用 sync/map
1. 不要直接修改 map valu 内某个元素的值，如果想修改 map 的某个键值，则必须整体赋值