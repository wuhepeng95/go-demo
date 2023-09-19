package test

import (
	"fmt"
	"testing"
	"time"
)

func TestDate(t *testing.T) {
	// 输入的时间字符串
	timeStr := "2023-08-10 21:12:00"

	// 解析时间字符串为 time.Time 类型
	layout := "2006-01-02 15:04:05"
	targetTime, err := time.Parse(layout, timeStr)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return
	}
	// 默认转换出来的时间是UTC时间
	fmt.Println("targetTime time:", targetTime)

	// 将时间转换为系统时区的时间
	sysTimezone, _ := time.Now().Zone()
	sysTime := targetTime.In(time.FixedZone(sysTimezone, 0))
	fmt.Println("System timezone time:", sysTime)

	fmt.Println("current time format:", time.Now().Format("20060102150405"))
}
