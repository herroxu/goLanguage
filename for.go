package main

import (
	"fmt"
	"strings"
)

// Go语言的For循环有3钟形式， 只有其中的一种使用分号

func main() {
	a := "jiu shi yi ge shi yan "
	fmt.Println(len(strings.TrimSpace(a)))
	fmt.Println(len(a))
}
