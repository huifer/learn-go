# GO 时间
>> 作者: [HuiFer](https://github.com/huifer)
>
> 仓库地址: https://github.com/huifer/learn-go
## time库
- time 提供了时间相关的操作

### 时间获取
- `time.Now()`
```go
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
```

### 时间戳 
> 时间戳: 时间戳是自1970年1月1日（08:00:00GMT）至当前时间的总毫秒数。它也被称为Unix时间戳（UnixTimestamp）。
- 获取时间戳,`time.Unix()`
```go
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
```
- 时间戳转换成时间
```go
/**
时间戳转换时间
*/
func TimeStamp2Time() {
	unix := time.Now().Unix()
	t := time.Unix(unix, 0)
	fmt.Println("时间: ", t)
}
```

### 时间操作
- 加法\减法\相等判断\时间前后判断
- 加法Add
    - 时间加法,两个时间相加
- 减法Sub
    - 时间减法,求出两个时间的插值
- 相等判断Equal
    - 时间相等判断,判断是否相等,相等返回True
- 之前Before
```go
// t在u之前,返回true，否则返回false
func (t Time) Before(u Time) bool {
	if t.wall&u.wall&hasMonotonic != 0 {
		return t.ext < u.ext
	}
	return t.sec() < u.sec() || t.sec() == u.sec() && t.nsec() < u.nsec()
}
```
- 之后After
```go
// t在u之后,返回true,否则返回false
func (t Time) After(u Time) bool {
	if t.wall&u.wall&hasMonotonic != 0 {
		return t.ext > u.ext
	}
	ts := t.sec()
	us := u.sec()
	return ts > us || ts == us && t.nsec() > u.nsec()
}
```

```go
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

```

### 时间格式化
- 注意: 格式化方式不是常见的`Y-m-d H:M:S`需要使用`"2006-01-02 15:04:05.000 Mon Jan"`
```go
/**
时间格式化
 */
func TimeFormat(){
	now := time.Now()
	// 24小时制
	fmt.Println(now.Format("2006-01-02 15:04:05.000 Mon Jan"))
	// 12小时制
	fmt.Println(now.Format("2006-01-02 03:04:05.000 PM Mon Jan"))
	fmt.Println(now.Format("2006/01/02 15:04"))
	fmt.Println(now.Format("15:04 2006/01/02"))
	fmt.Println(now.Format("2006/01/02"))
}
```