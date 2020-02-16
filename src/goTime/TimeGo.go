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

/*
时间操作
*/
func TimeOperation() {
	unix := time.Unix(1581837895415/1000, 0)
	// 加法
	add := unix.Add(1)
	fmt.Println("时间加法结果", add)
	// 减法
	sub := time.Now().Sub(add)
	fmt.Println("时间减法结果", sub)
	// 相等判断
	equal := add.Equal(add)
	fmt.Println("相等吗", equal)
	// 判断之前
	before := add.Before(add)
	fmt.Println("之前吗", before)
	// 判断之后
	after := add.After(add)
	fmt.Println("之后吗", after)
}

/**
时间格式化
*/
func TimeFormat() {
	now := time.Now()
	// 24小时制
	fmt.Println(now.Format("2006-01-02 15:04:05.000 Mon Jan"))
	// 12小时制
	fmt.Println(now.Format("2006-01-02 03:04:05.000 PM Mon Jan"))
	fmt.Println(now.Format("2006/01/02 15:04"))
	fmt.Println(now.Format("15:04 2006/01/02"))
	fmt.Println(now.Format("2006/01/02"))
}
