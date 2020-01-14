package main

import "fmt"

// 寻找最长不含有重复字符的字串

func maxLength(str string) int {

	for i, ch := range []byte(str) {
		fmt.Println(i, ch, []byte(str))
	}
	return 2
}

func main() {
	maxLength("abchedf")

	// 每个中文是三字节 UTF-8
}


