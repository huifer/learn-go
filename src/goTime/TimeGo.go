package goTime

import (
	"fmt"
	"time"
)

/**
时间获取
*/
func TimeDemo() {
	// 当前时间
	now := time.Now()
	// 年月日时分秒
	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()

	fmt.Println(now)
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)

}

/**
时间戳
*/
func TimestampDemo() {
	now := time.Now()
	// 毫秒级别时间戳
	unix := now.Unix()
	// 纳秒级别时间戳
	nano := now.UnixNano()
	fmt.Println("毫秒", unix)
	fmt.Println("纳秒", nano)

}

/**
时间戳转换时间
*/
func TimeStamp2Time() {
	unix := time.Now().Unix()
	t := time.Unix(unix, 0)
	fmt.Println("时间: ", t)
}
