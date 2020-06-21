package slice

import "fmt"

/**
go 语言slice使用
*/
func SliceDemo() {
	// 通过数组创建
	a := [3]int{1, 2, 3}
	start := a[1:]
	end1 := a[:len(a)]
	end2 := a[1:len(a)]

	fmt.Println(start)
	fmt.Println(end1)
	fmt.Println(end2)
	// 通过 make 创建
	b1 := make([]int, 10)
	b2 := make([]int, 10, 20)
	fmt.Println(b1)
	fmt.Println(b2)
	// 内置函数
	// 返回切片长度
	fmt.Println(len(b1))
	// 返回切片底层数组容量
	fmt.Println(cap(b1))
	// 追加元素,原slice 不会修改
	ints := append(b1, 1)
	fmt.Println("追加返回值", ints)
	fmt.Println("原始值", b1)
	// 拷贝
	b3 := make([]int, 10, 10)
	copy(b1, b3)
	fmt.Println(b3)

}
