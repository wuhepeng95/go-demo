package test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/robfig/cron"
)

func TestCorn(t *testing.T) {
	c := cron.New()

	// 添加定时任务
	err := c.AddFunc("0/2 * * * * ?", func() {
		fmt.Println("定时任务执行于", time.Now())

		// // 在Go语言中，无法直接获取当前线程的详细信息。
		// // Go运行时（runtime）会在 Goroutine 和线程之间进行调度，并使用固定数量的线程池来执行 Goroutine。
		// // 然而，您可以使用 runtime 包来获取一些与 Goroutine 和并发相关的信息
		// num := runtime.NumGoroutine()
		// fmt.Println("当前 Goroutine 数量:", num)

		url := "https://gateway.test.isrm.going-link.com/sprm/v1/3/purchase-requests/__-702ov1_PyvVG1jHnhvphZQ-__/actions/all"
		method := "GET"

		client := &http.Client{}
		req, err := http.NewRequest(method, url, nil)

		if err != nil {
			fmt.Println(err)
			return
		}
		req.Header.Add("Authorization", "bearer 8d62595a-feec-4800-87bd-d3dfd4e6d016")
		req.Header.Add("H-Menu-Id", "__-majbaCHkaJetQLFSizyQFg-__")
		req.Header.Add("H-Request-Id", "874436c46c4a66e36434fd3b825e576ce51ec72458861")
		req.Header.Add("Pragma", "no-cache")
		req.Header.Add("h-request-trigger", "%E6%93%8D%E4%BD%9C%E8%AE%B0%E5%BD%95")
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
