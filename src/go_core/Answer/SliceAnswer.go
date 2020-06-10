package Answer

import "fmt"

/**
GO 语言切片练习
1. 创建数组 []int{1,2,3}
2. 获取这个数组的前半部分和后半补分
*/
func SliceAnswer() {
	a := []int{1, 2, 3}
	front := a[0 : len(a)/2]
	behind := a[len(a)/2:]
	fmt.Println(front)
	fmt.Println(behind)
}
