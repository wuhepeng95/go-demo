package test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestHttp(t *testing.T) {
	// 发送 GET 请求到百度首页
	resp, err := http.Get("https://www.baidu.com")
	if err != nil {
		fmt.Println("发送请求失败:", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应内容
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应失败:", err)
		return
	}

	// 打印响应内容
	fmt.Println(string(body))
}
