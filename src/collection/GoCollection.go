package collection

import "fmt"

func MapFunc() {
	var cMap map[string]string
	cMap = make(map[string]string)

	cMap["a"] = "1"
	cMap["b"] = "2"

	for s := range cMap {
		fmt.Println(s, cMap[s])
	}

	value, ok := cMap["a"]
	if ok {
		fmt.Println(value)

	} else {
		fmt.Println("not has key")
	}
}
