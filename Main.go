package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	startDate, err := ConvertTimeStringToTodayTime("15:04:05", "14:00:00", "20230815")
	if err != nil {
	}
	endDate, err := ConvertTimeStringToTodayTime("15:04:05", "14:59:59", "20230815")
	if err != nil {
	}

	fmt.Println("startDate time:", startDate)
	fmt.Println("endDate time:", endDate)

	fileInfo, err := os.Stat("/dump/20230815/srm-prm-5bca2-57d5c56686-66knz-20230815-143449.txt")
	if err != nil {
		fmt.Println("read file with error")
	}
	fileDateTime := fileInfo.ModTime()
	fmt.Println("fileDateTime time:", fileDateTime)

	fmt.Println("fileDateTime.After(startDate):", fileDateTime.After(startDate))
	fmt.Println("fileDateTime.Before(endDate)", fileDateTime.Before(endDate))

}

func ConvertTimeStringToTodayTime(timeLayout string, timeStr string, dateString string) (time.Time, error) {
	// 获取当前日期
	currentTime, _ := time.Parse("20060102", dateString)
	timeValue, err := time.Parse(timeLayout, timeStr)
	if err != nil {
		return time.Time{}, err
	}
	// 创建当天的时间对象，保持日期不变，只更新时间部分
	todayTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(),
		timeValue.Hour(), timeValue.Minute(), timeValue.Second(), timeValue.Nanosecond(), time.Now().Location())
	return todayTime, nil
}
