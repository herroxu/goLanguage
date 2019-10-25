package main

import (
	"strings"
	"fmt"
)

// 字符串处理函数，将字符串切片和处理链放入
func ProcessString(list []string, chain []func(string)string) {
	for index, char := range list {
		newChar := char
		for _,pro := range chain {
			newChar = pro(newChar)
		}
		list[index] = newChar
	}
}

func remow(c string) string {
	return strings.TrimPrefix(c, "go")
}

func main()  {
	list := []string{
		"go language",
		"go zhiming",
		"go world",
	}
	chain := []func(string)string{
		remow,
		strings.ToUpper,
		strings.TrimSpace,
	}
	ProcessString(list, chain)

	for _, str := range list {
		fmt.Println(str)
	}

}

