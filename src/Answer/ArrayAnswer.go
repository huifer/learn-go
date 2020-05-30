package Answer

import "fmt"

/**
设置一个int类型的数组长度为10
修改最后一个元素的值为100
*/
func ArrayAnswer() {
	a := [10]int{}
	for i := 0; i < 10; i++ {
		a[i] = i + 1
	}
	fmt.Println(a)
	a[len(a)-1] = 100
	fmt.Println(a)

}
