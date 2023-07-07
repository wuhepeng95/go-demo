package test

import (
	"fmt"
	"strconv"
	"testing"
)

func TestConvert(t *testing.T) {
	// string to int
	aStr := "100"

	bInt, err := strconv.Atoi(aStr)

	if err == nil {
		fmt.Printf("aStr：%T %s，bInt：%T %d", aStr, aStr, bInt, bInt)
	} else {
		fmt.Printf("err：%s", err)
	}

	println()

	// int to string
	cInt := 200
	dStr := strconv.Itoa(cInt)

	fmt.Printf("cInt：%T %d，dStr：%T %s", cInt, cInt, dStr, dStr)

	println()
}
