package mapDemo

import "fmt"

func MapTest() {
	// 字面量创建
	a_map := map[string]int{"a": 1, "b": 2}
	fmt.Println(a_map)
	// make 函数创建
	map1 := make(map[string]int)
	fmt.Println(map1)
	// map 操作

	// 1. 访问
	fmt.Println(a_map["a"])
	// 2. 删除
	delete(a_map, "a")
	// 3. 修改
	a_map["b"] = 22
	// 4. 循环
	for s := range a_map {
		fmt.Println("key = ", s, " value = ", a_map[s])
	}
	// 5. map key 长度
	fmt.Println(len(a_map))

}
