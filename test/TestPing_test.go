package test

import (
	"fmt"
	"os/exec"
	"testing"
)

func TestPing(t *testing.T) {
	ip := "www.baidu.com"

	cmd := exec.Command("ping", "-c", "1", ip)

	err := cmd.Run()
	if err != nil {
		fmt.Println("ping 失败:", err)
		fmt.Println(false)
		return
	}

	fmt.Println("pong")
}
