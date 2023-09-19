package test

import (
	"fmt"
	"testing"
	"time"
)

//var DataFixChannel = make(chan string)

var DataFixChannel = make(chan string, 1) // 缓冲区大小为2

func TestChain(t *testing.T) {
	// 初始化
	Init()

	strList := []string{"A", "B", "C", "D"}
	// 放消息
	for _, str := range strList {
		//fmt.Println("循环遍历：" + str)
		str := str
		go func() {
			DataFixChannel <- str
			fmt.Println("已加入通道：" + str)
		}()
	}
	// 主线程等待所有协程结束
	time.Sleep(1 * time.Minute)
}

func Init() {
	SubscribeDataFixInit()
}

func SubscribeDataFixInit() {
	go func() {
		for {
			str, ok := <-DataFixChannel
			fmt.Println("已取出通道：" + str)
			if !ok {
				return
			}
			processTask(str)
		}
	}()
}

func processTask(str string) {
	fmt.Println("----------正在处理：" + str)
	time.Sleep(2 * time.Second)
	fmt.Println("----------处理结束：" + str)
	time.Sleep(1 * time.Second)
}
