package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/robfig/cron"
)

func main() {
	c := cron.New()

	// 添加定时任务
	err := c.AddFunc("0/1 * * * * ?", func() {
		fmt.Println("定时任务执行于", time.Now())

		// // 在Go语言中，无法直接获取当前线程的详细信息。
		// // Go运行时（runtime）会在 Goroutine 和线程之间进行调度，并使用固定数量的线程池来执行 Goroutine。
		// // 然而，您可以使用 runtime 包来获取一些与 Goroutine 和并发相关的信息
		// num := runtime.NumGoroutine()
		// fmt.Println("当前 Goroutine 数量:", num)

		url := "https://gateway.dev.isrm.going-link.com/srpm/v1/30/request-plan/todo-list?page=0&size=20&asyncCountFlag=DEFAULT&containerId=__-lE11y_COtdxoL7Uzm2qgow-__&advancedData=%257B%2522customizeOrderField%2522:%2522rpNum:asc%2522%257D&customizeOrderField=rpNum:asc&customizeUnitCode=SRPM.RP_EXECUTE_PLATFORM.TODO.LIST,SRPM.RP_EXECUTE_PLATFORM.TODO.LIST.SEARCH_BAR"
		method := "GET"

		client := &http.Client{}
		req, err := http.NewRequest(method, url, nil)

		if err != nil {
			fmt.Println(err)
			return
		}
		req.Header.Add("authority", "gateway.dev.isrm.going-link.com")
		req.Header.Add("authorization", "bearer ad54aca8-9dc4-4e0b-92b9-e17bcf9d6756")
		req.Header.Add("h-menu-id", "__-Pu7EAnfSWDIknTHfkS5pGg-__")
		req.Header.Add("h-request-id", "8388f3d89df24e8944445ab8b97825154870d45658861")
		req.Header.Add("pragma", "no-cache")
		req.Header.Add("User-Agent", "Apifox/1.0.0 (https://www.apifox.cn)")

		res, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(body))
	})
	if err != nil {
		fmt.Println("添加定时任务失败:", err)
		return
	}

	c.Start()

	// 程序运行一段时间后停止定时任务
	time.Sleep(1000000 * time.Second)
	c.Stop()

	fmt.Println("定时任务已停止")
}
