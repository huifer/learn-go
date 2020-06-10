package array

import "fmt"

func Array() {
	a := [3]int{1, 2, 3}
	fmt.Println(a)
	// 输出数组的长度
	i := len(a)
	fmt.Println("数字长度= ", i)
	// 数组的访问
	i2 := a[0]
	fmt.Println("数组访问,index=0，value=", i2)

}
