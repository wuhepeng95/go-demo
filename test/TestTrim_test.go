package test

import (
	"fmt"
	"strings"
	"testing"
)

func TestName(t *testing.T) {
	var AndLogicSlice = []rune{'；', ';', ',', '，', '+', '/', '、', '和', '与', '加'}
	for _, num := range Splits(strings.Trim("cdp-37441,    \ncdp-38151\n", "\n"), AndLogicSlice) {
		fmt.Println(strings.Trim(num, " \n"))
	}
}

// Splits 多个单字符分割
func Splits(s string, seps []rune) []string {
	Split := func(r rune) bool {
		for _, v := range seps {
			if v == r {
				return true
			}
		}
		return false
	}
	a := strings.FieldsFunc(s, Split)
	return a
}
