package Answer

import "fmt"

/**
乘法表
*/
func multiplication() {
	for col := 1; col <= 9; col++ {
		for row := 1; row <= col; row++ {
			fmt.Printf("%d*%d=%d\t", col, row, col*row)
		}
		fmt.Print("\r\n")
	}
}

/**
阶乘
*/
func summation(n int) int {
	var result int
	for i := 1; i <= n; i++ {
		tem := 1
		for j := 1; j <= i; j++ {
			tem *= j
		}
		result += tem

	}
	return result
}

/**
统计每一个类型的数量
*/
func countString(s string) (wordCount, spaceCount, numCount, otherCount int) {
	chars := []rune(s)
	for _, value := range chars {
		switch {
		case value >= 'a' && value <= 'z':
			wordCount += 1
		case value >= 'A' && value <= 'Z':
			wordCount += 1
		case value == ' ':
			spaceCount += 1
		case value >= '0' && value <= '9':
			numCount += 1
		default:
			otherCount += 1

		}
	}
	return

}
