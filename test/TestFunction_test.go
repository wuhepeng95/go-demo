package test

import (
	"fmt"
	"testing"
)

func TestFunction(t *testing.T){
	// 匿名函数
	max := func(a int , b int) int {
		if(a > b) {
			return a
		} else {
			return b
		}
	}

	fmt.Println(max(1,1))
}